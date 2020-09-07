package asm

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
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

func generateFuncDiff(src_f, dst_f *ir.Func, m *ASTMapping, container_mapping map[*ir.Block]*ir.Block, src_g *Graph) ([]operation, string, *gographviz.Graph) {
	var func_script []operation
	var contents bytes.Buffer
	src_df := m.src_dfs[src_f]
	// dst_df := m.dst_dfs[dst_f]
	graph_diff := DiffGraphInit(src_f, dst_f)
	//TODO: update src_g's tags by edit_script
	block_script := blockEditScript(src_f, dst_f, container_mapping)
	fmt.Println(block_script)
	srcIndex, dstIndex := 0, 0
	for _, op := range block_script {
		switch op {
		case INSERT:
			var script []operation
			buf := &strings.Builder{}
			ins_b := dst_f.Blocks[dstIndex]
			// src_g.UpdateBlockTag(ins_b, op)
			for k := range ins_b.Insts {
				dst_inst := &(*ins_b).Insts[k]
				//Step1. EditScript
				script = append(script, INSERT)
				fmt.Fprintln(buf, colors[1]+"+"+(*dst_inst).LLString())

				//Step2. DiffDotGraph
				v := reflect.ValueOf(*dst_inst)
				m := v.MethodByName("Ident")
				if m.IsValid() {
					dst_inst_id := m.Call(nil)[0].String()
					if dst_inst_id != "%0" {
						dst_inst_id = Add_quotation_marks(dst_inst_id, "Dst")
						dst_b_id := Add_quotation_marks(ins_b.Ident(), "Dst")
						// graph_diff.RemoveNode(dst_b_id, dst_inst_id)
						graph_diff.AddNode(dst_b_id, dst_inst_id, map[string]string{"style": "filled", "fillcolor": "yellow"})

					}
				}

			}
			// Insert Term node
			script = append(script, INSERT)
			fmt.Fprintln(buf, colors[1]+"+"+ins_b.Term.LLString())

			contents.WriteString(buf.String())
			dstIndex += 1

		case MOVE:
			//Step1. EditScript
			script, buf := generateBlockDiff(src_f.Blocks[srcIndex], dst_f.Blocks[dstIndex], m, container_mapping, src_g, graph_diff)
			func_script = append(func_script, script...)
			contents.WriteString(buf)

			//Step2. DiffDotGraph
			src_b_id := Add_quotation_marks(src_f.Blocks[srcIndex].Ident(), "Src")
			dst_b_id := Add_quotation_marks(dst_f.Blocks[dstIndex].Ident(), "Dst")
			graph_diff.AddEdge(src_b_id, dst_b_id, true, map[string]string{"color": "blue", "label": "MOV"})

			srcIndex += 1
			dstIndex += 1

		case DELETE:
			//Step1. EditScript
			var script []operation
			buf := &strings.Builder{}
			del_b := src_f.Blocks[srcIndex]

			src_g.UpdateBlockTag(del_b, DELETE, src_df) //set whole basic block to DELETE
			for k := range del_b.Insts {
				src_inst := &(src_f.Blocks[srcIndex].Insts)[k]
				script = append(script, DELETE)
				fmt.Fprintln(buf, colors[2]+"-"+(*src_inst).LLString())

				//Step2. DiffDotGraph
				v := reflect.ValueOf(*src_inst)
				m := v.MethodByName("Ident")
				if m.IsValid() {
					src_inst_id := m.Call(nil)[0].String()
					if src_inst_id != "%0" {
						src_inst_id = Add_quotation_marks(src_inst_id, "Src")
						src_b_id := Add_quotation_marks(del_b.Ident(), "Src")
						// graph_diff.RemoveNode(src_b_id, src_inst_id)
						graph_diff.AddNode(src_b_id, src_inst_id, map[string]string{"style": "filled", "fillcolor": "red"})

					}
				}
			}

			//Delete Term node
			script = append(script, DELETE)
			fmt.Fprintln(buf, colors[2]+"-"+del_b.Term.LLString())

			contents.WriteString(buf.String())

			//Step3.Update
			srcIndex += 1

		case UPDATE:
			fmt.Println("error")
		}
	}
	return func_script, contents.String(), graph_diff
}

func generateBlockDiff(src_block, dst_block *ir.Block, m *ASTMapping, container_mapping map[*ir.Block]*ir.Block, src_g *Graph, graph_diff *gographviz.Graph) ([]operation, string) {
	src_f := src_block.Parent
	dst_f := dst_block.Parent
	src_df := m.src_dfs[src_block.Parent]
	dst_df := m.dst_dfs[dst_block.Parent]

	var src []*ir.Instruction
	for k := range src_block.Insts {
		src = append(src, &(src_block.Insts)[k])
	}

	var dst []*ir.Instruction
	for k := range dst_block.Insts {
		dst = append(dst, &(dst_block.Insts)[k])
	}

	//Step 1 use myers diff algorithm
	script := shortestInstEditScript(src, dst, src_f, dst_f, src_df, dst_df)

	//Step 2.1 Optimal : search for update action
	// src_ambiguous := make(map[*ir.Instruction]int)
	// dst_ambiguous := make(map[*ir.Instruction]int)
	// srcIndex, dstIndex := 0, 0
	// for i, op := range script {
	// 	switch op {
	// 	case INSERT:
	// 		dst_ambiguous[dst[dstIndex]] = i
	// 		dstIndex += 1

	// 	case MOVE:
	// 		srcIndex += 1
	// 		dstIndex += 1

	// 	case DELETE:
	// 		src_ambiguous[src[srcIndex]] = i
	// 		srcIndex += 1
	// 	case UPDATE:
	// 		fmt.Println("error")
	// 	}
	// }
	// Step 2.2 [Optinal] optimal for Update operation

	// for src_inst, i := range src_ambiguous {
	// 	for dst_inst, j := range dst_ambiguous {
	// 		if ambCompare(src_inst, dst_inst, src_df, dst_df) {
	// 			//优化为update
	// 			script[i] = 4
	// 			script[j] = 5 //变成unknown,不输出
	// 			delete(dst_ambiguous, dst_inst)
	// 		}
	// 	}
	// }
	//Step 3 Print edit_scripts
	srcIndex, dstIndex := 0, 0
	buf := &strings.Builder{}
	for _, op := range script {
		switch op {
		case INSERT:
			// src_g.UpdateInstTag(dst[dstIndex], op)
			fmt.Fprintln(buf, colors[op]+"+"+(*dst[dstIndex]).LLString())
			if e_str := getExtends(dst[dstIndex]); e_str != "" { //print struct extends
				fmt.Fprintln(buf, colors[op]+"++", e_str)
			}

			//Step2. DiffDotGraph
			dst_inst := dst[dstIndex]
			dst_reflect_v := reflect.ValueOf(*dst_inst)
			dst_reflect_m := dst_reflect_v.MethodByName("Ident")
			if dst_reflect_m.IsValid() {
				dst_inst_id := dst_reflect_m.Call(nil)[0].String()
				if dst_inst_id != "%0" {
					dst_inst_id = Add_quotation_marks(dst_inst_id, "Dst")
					dst_b_id := Add_quotation_marks(dst_block.Ident(), "Dst")
					// graph_diff.RemoveNode(dst_b_id, dst_inst_id)
					graph_diff.AddNode(dst_b_id, dst_inst_id, map[string]string{"style": "filled", "fillcolor": "yellow"})

				}
			}

			//Step3. Update
			// fmt.Println(colors[op] + "+" + (*dst[dstIndex]).LLString())
			dstIndex += 1

		case MOVE:
			//Step1. EditScript
			fmt.Fprintln(buf, colors[op]+" "+(*src[srcIndex]).LLString())

			//Step2. DiffDotGraph
			src_inst := src[srcIndex]
			dst_inst := dst[dstIndex]
			src_reflect_v := reflect.ValueOf(*src_inst)
			dst_reflect_v := reflect.ValueOf(*dst_inst)
			src_reflect_m := src_reflect_v.MethodByName("Ident")
			dst_reflect_m := dst_reflect_v.MethodByName("Ident")
			if src_reflect_m.IsValid() && dst_reflect_m.IsValid() {
				src_inst_id := src_reflect_m.Call(nil)[0].String()
				dst_inst_id := dst_reflect_m.Call(nil)[0].String()
				if src_inst_id != "%0" && dst_inst_id != "%0" {
					src_inst_id = Add_quotation_marks(src_inst_id, "Src")
					dst_inst_id = Add_quotation_marks(dst_inst_id, "Dst")
					graph_diff.AddEdge(src_inst_id, dst_inst_id, true, map[string]string{"color": "blue", "label": "MOV"})
				}

			}

			//Step3. Update
			// fmt.Println(colors[op] + " " + (*src[srcIndex]).LLString())
			srcIndex += 1
			dstIndex += 1

		case DELETE:
			//Step1. EditScript
			src_g.UpdateInstTag(src[srcIndex], op, src_df)
			fmt.Fprintln(buf, colors[op]+"-"+(*src[srcIndex]).LLString())
			if e_str := getExtends(src[srcIndex]); e_str != "" { //print struct extends
				fmt.Fprintln(buf, colors[op]+"--", e_str)
			}

			//Step2. DiffDotGraph
			src_inst := src[srcIndex]
			src_reflect_v := reflect.ValueOf(*src_inst)
			src_reflect_m := src_reflect_v.MethodByName("Ident")

			if src_reflect_m.IsValid() {

				src_inst_id := src_reflect_m.Call(nil)[0].String()
				if src_inst_id != "%0" {
					src_inst_id = Add_quotation_marks(src_inst_id, "Src")
					src_b_id := Add_quotation_marks(src_block.Ident(), "Src")

					// graph_diff.RemoveNode(src_b_id, src_inst_id)
					graph_diff.AddNode(src_b_id, src_inst_id, map[string]string{"style": "filled", "fillcolor": "red"})
				}
			}

			//Step3. Update
			// fmt.Println(colors[op] + "-" + (*src[srcIndex]).LLString())
			srcIndex += 1

		case UPDATE:
			src_g.UpdateInstTag(src[srcIndex], DELETE, src_df) //Fixme
			src_g.UpdateInstTag(dst[dstIndex], INSERT, dst_df) //Fixme
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
		src_g.UpdateTermTag(&(src_block.Term), DELETE, src_df) //Fixme
		fmt.Fprintln(buf, colors[term_op]+" "+dst_block.Term.LLString())
	}
	script = append(script, term_op)

	return script, buf.String()
}

// Meyers for block
func blockEditScript(src_f, dst_f *ir.Func, container_mapping map[*ir.Block]*ir.Block) []operation {
	//TODO:基于控制流图来生成edit script更符合编程人员逻辑

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

func getExtends(inst *ir.Instruction) string {
	switch inst_v := (*inst).(type) {
	case *ir.InstCall:
		if f := inst_v.GetGlobalFunc(); f != nil {
			return f.LLString()
		}
	case *ir.InstLoad:
		if g := inst_v.GetGlobal(); g != nil {
			return g.LLString()
		}
	case *ir.InstStore:
		if g := inst_v.GetGlobal(); g != nil {
			return g.LLString()
		}
	case *ir.InstAlloca:
		if t := inst_v.GetTypeDefs(); t != nil {
			if types.IsPointer(t) {
				return t.(*types.PointerType).ElemType.LLString()
			} else {
				return fmt.Sprintf("%s = type %s", t, t.LLString())
			}
		}
	}
	return ""
}
