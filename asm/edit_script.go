package asm

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/llir/llvm/ir"
)

type operation uint

const (
	INSERT  operation = 1
	DELETE            = 2
	MOVE              = 3
	UPDATE            = 4
	UNKNOWN           = 5
)

func (op operation) String() string {
	switch op {
	case INSERT:
		return "INS"
	case DELETE:
		return "DEL"
	case MOVE:
		return "MOV"
	case UPDATE:
		return "UPDATE"
	default:
		return "UNKNOWN"
	}
}

var colors = map[operation]string{
	INSERT: "\033[32m",
	DELETE: "\033[31m",
	MOVE:   "\033[39m",
	UPDATE: "\033[32m",
}

func generateFuncDiff(src_f, dst_f *ir.Func, m *ASTMapping, container_mapping map[*ir.Block]*ir.Block) ([]operation, string) {
	var func_script []operation
	var contents bytes.Buffer
	block_script := blockEditScript(src_f, dst_f, container_mapping)
	srcIndex, dstIndex := 0, 0
	for _, op := range block_script {
		switch op {
		case INSERT:
			var script []operation
			buf := &strings.Builder{}
			for _, dst_inst := range dst_f.Blocks[dstIndex].Insts {
				script = append(script, INSERT)
				fmt.Fprintln(buf, colors[1]+"+"+dst_inst.LLString())
			}
			contents.WriteString(buf.String())
			dstIndex += 1

		case MOVE:
			script, buf := generateBlockDiff(src_f.Blocks[srcIndex], dst_f.Blocks[dstIndex], m, container_mapping)
			func_script = append(func_script, script...)
			contents.WriteString(buf)
			srcIndex += 1
			dstIndex += 1

		case DELETE:
			var script []operation
			buf := &strings.Builder{}
			for _, src_inst := range src_f.Blocks[srcIndex].Insts {
				script = append(script, DELETE)
				fmt.Fprintln(buf, colors[2]+"-"+src_inst.LLString())
			}
			contents.WriteString(buf.String())
			srcIndex += 1

		case UPDATE:
			fmt.Println("error")
		}
	}
	return func_script, contents.String()
}

func generateBlockDiff(src_block, dst_block *ir.Block, m *ASTMapping, container_mapping map[*ir.Block]*ir.Block) ([]operation, string) {
	var src []*ir.Instruction
	src_f := src_block.Parent
	dst_f := dst_block.Parent
	for k := range src_block.Insts {
		src = append(src, &(src_block.Insts)[k])
	}

	var dst []*ir.Instruction
	for k := range dst_block.Insts {
		dst = append(dst, &(dst_block.Insts)[k])
	}

	src_df := m.src_dfs[src_block.Parent]
	dst_df := m.dst_dfs[dst_block.Parent]

	//Step 1 use myers diff algorithm
	script := shortestInstEditScript(src, dst, src_f, dst_f, src_df, dst_df)

	//Step 2 Optimal : search for update action
	src_ambiguous := make(map[*ir.Instruction]int)
	dst_ambiguous := make(map[*ir.Instruction]int)
	srcIndex, dstIndex := 0, 0
	for i, op := range script {
		switch op {
		case INSERT:
			dst_ambiguous[dst[dstIndex]] = i
			dstIndex += 1

		case MOVE:
			srcIndex += 1
			dstIndex += 1

		case DELETE:
			src_ambiguous[src[srcIndex]] = i
			srcIndex += 1
		case UPDATE:
			fmt.Println("error")
		}
	}
	for src_inst, i := range src_ambiguous {
		for dst_inst, j := range dst_ambiguous {
			if ambCompare(src_inst, dst_inst, src_df, dst_df) {
				//优化为update
				script[i] = 4
				script[j] = 5 //变成unknown,不输出
				delete(dst_ambiguous, dst_inst)
			}
		}
	}
	//Step 3 Print edit_scripts
	srcIndex, dstIndex = 0, 0
	buf := &strings.Builder{}
	for _, op := range script {
		switch op {
		case INSERT:
			fmt.Fprintln(buf, colors[op]+"+"+(*dst[dstIndex]).LLString())
			// fmt.Println(colors[op] + "+" + (*dst[dstIndex]).LLString())
			dstIndex += 1

		case MOVE:
			fmt.Fprintln(buf, colors[op]+" "+(*src[srcIndex]).LLString())
			// fmt.Println(colors[op] + " " + (*src[srcIndex]).LLString())
			srcIndex += 1
			dstIndex += 1

		case DELETE:
			fmt.Fprintln(buf, colors[op]+"-"+(*src[srcIndex]).LLString())
			// fmt.Println(colors[op] + "-" + (*src[srcIndex]).LLString())
			srcIndex += 1

		case UPDATE:
			fmt.Fprintln(buf, colors[op]+" "+(*dst[dstIndex]).LLString())
			// fmt.Println(colors[op] + " " + (*dst[dstIndex]).LLString())
			dstIndex += 1
			srcIndex += 1

		case UNKNOWN:
			continue
		}
	}

	//TODO:比较两个basic block terminator 节点
	term_op := terminatorCompare(src_block, dst_block, src_f, dst_f, src_df, dst_df, container_mapping)
	switch term_op {
	case MOVE:
		fmt.Fprintln(buf, colors[term_op]+" "+src_block.Term.LLString())
	case UPDATE:
		fmt.Fprintln(buf, colors[term_op]+" "+dst_block.Term.LLString())
	}
	script = append(script, term_op)

	return script, buf.String()
}

// Meyers for block
func blockEditScript(src_f, dst_f *ir.Func, container_mapping map[*ir.Block]*ir.Block) []operation {
	src := src_f.Blocks
	dst := dst_f.Blocks
	n := len(src)
	m := len(dst)
	max := n + m
	var trace []map[int]int
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		// 最多只有 d+1 个 k
		v := make(map[int]int, d+2)
		trace = append(trace, v)

		// 需要注意处理对角线
		if d == 0 {
			t := 0
			for len(src) > t && len(dst) > t && container_mapping[src[t]] == dst[t] {
				t++
			}
			v[0] = t
			if t == len(src) && t == len(dst) {
				break loop
			}
			continue
		}

		lastV := trace[d-1]

		for k := -d; k <= d; k += 2 {
			// 向下
			if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
				x = lastV[k+1]
			} else { // 向右
				x = lastV[k-1] + 1
			}

			y = x - k

			// 处理对角线
			for x < n && y < m && container_mapping[src[x]] == dst[y] {
				x, y = x+1, y+1
			}

			v[k] = x

			if x == n && y == m {
				break loop
			}
		}
	}

	// for debug
	// printTrace(trace)

	// 反向回溯
	var script []operation

	x = n
	y = m
	var k, prevK, prevX, prevY int

	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		lastV := trace[d-1]

		if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX = lastV[prevK]
		prevY = prevX - prevK

		for x > prevX && y > prevY {
			script = append(script, MOVE)
			x -= 1
			y -= 1
		}

		if x == prevX {
			script = append(script, INSERT)
		} else {
			script = append(script, DELETE)
		}

		x, y = prevX, prevY
	}

	if trace[0][0] != 0 {
		for i := 0; i < trace[0][0]; i++ {
			script = append(script, MOVE)
		}
	}

	return reverse(script)
}

// 生成最短的编辑脚本
func shortestInstEditScript(src, dst []*ir.Instruction, src_f, dst_f *ir.Func, src_df, dst_df *DataFlow) []operation {

	n := len(src)
	m := len(dst)
	max := n + m
	var trace []map[int]int
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		// 最多只有 d+1 个 k
		v := make(map[int]int, d+2)
		trace = append(trace, v)

		// 需要注意处理对角线
		if d == 0 {
			t := 0
			for len(src) > t && len(dst) > t && instCompare(src[t], dst[t], src_f, dst_f, src_df, dst_df) {
				t++
			}
			v[0] = t
			if t == len(src) && t == len(dst) {
				break loop
			}
			continue
		}

		lastV := trace[d-1]

		for k := -d; k <= d; k += 2 {
			// 向下
			if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
				x = lastV[k+1]
			} else { // 向右
				x = lastV[k-1] + 1
			}

			y = x - k

			// 处理对角线
			for x < n && y < m && instCompare(src[x], dst[y], src_f, dst_f, src_df, dst_df) {
				x, y = x+1, y+1
			}

			v[k] = x

			if x == n && y == m {
				break loop
			}
		}
	}

	// for debug
	// printTrace(trace)

	// 反向回溯
	var script []operation

	x = n
	y = m
	var k, prevK, prevX, prevY int

	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		lastV := trace[d-1]

		if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX = lastV[prevK]
		prevY = prevX - prevK

		for x > prevX && y > prevY {
			script = append(script, MOVE)
			x -= 1
			y -= 1
		}

		if x == prevX {
			script = append(script, INSERT)
		} else {
			script = append(script, DELETE)
		}

		x, y = prevX, prevY
	}

	if trace[0][0] != 0 {
		for i := 0; i < trace[0][0]; i++ {
			script = append(script, MOVE)
		}
	}

	return reverse(script)
}

func printTrace(trace []map[int]int) {
	for d := 0; d < len(trace); d++ {
		fmt.Printf("d = %d:\n", d)
		v := trace[d]
		for k := -d; k <= d; k += 2 {
			x := v[k]
			y := x - k
			fmt.Printf("  k = %2d: (%d, %d)\n", k, x, y)
		}
	}
}

func reverse(s []operation) []operation {
	result := make([]operation, len(s))

	for i, v := range s {
		result[len(s)-1-i] = v
	}

	return result
}

func getFileLines(p string) ([]string, error) {
	f, err := os.Open(p)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
