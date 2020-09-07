package asm

import (
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
)

type Mapping struct {
	src    *ir.Block
	dst    *ir.Block
	dice   float64
	n_dice float64
	common int
}

type Mappings []Mapping

//Len()
func (m Mappings) Len() int {
	return len(m)
}

//Less():
func (m Mappings) Less(i, j int) bool { //逆序排列，dice值从高到低
	if m[i].common != m[j].common {
		return m[i].common > m[j].common
	} else if m[i].dice != m[j].dice {
		return m[i].dice > m[j].dice
	} else if m[i].n_dice != m[j].n_dice {
		return m[i].n_dice > m[j].n_dice
	} else { //靠前匹配，先匹配label总和更小的：%12:%12 > %12:%42
		return m[i].src.LocalID+m[i].dst.LocalID < m[j].src.LocalID+m[j].dst.LocalID
	}
}

//Swap()
func (m Mappings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Mappings) String() string {
	buf := &strings.Builder{}
	for _, k := range m {
		s := fmt.Sprintf("[%s:%s]:(%f) ", k.src.String(), k.dst.String(), k.dice)
		buf.WriteString(s)
	}
	return buf.String()
}

func commonNumber(src_block, dst_block *ir.Block, m *ASTMapping) int {
	common := 0
	src_f := src_block.Parent
	dst_f := dst_block.Parent
	dstIgnored := make(map[*ir.Instruction]bool)
	var src_df, dst_df *DataFlow

	if _, ok := m.src_dfs[src_f]; ok { //命中
		src_df = m.src_dfs[src_f]
	} else {
		src_df = InstDependence(src_f, m.debugInfo)
		m.src_dfs[src_f] = src_df
	}
	if _, ok := m.dst_dfs[dst_f]; ok { //命中
		dst_df = m.dst_dfs[dst_f]
	} else {
		dst_df = InstDependence(dst_f, m.debugInfo)
		m.dst_dfs[dst_f] = dst_df
	}

	for i := range src_block.Insts { //Fixme 有重复匹配
		for j := range dst_block.Insts {
			if !dstIgnored[&(dst_block.Insts)[j]] {
				if instCompare(&(src_block.Insts)[i], &(dst_block.Insts)[j], src_f, dst_f, src_df, dst_df) {
					dstIgnored[&(dst_block.Insts)[j]] = true
					common++
				}
			}
		}
	}
	return common
}

func (k *Mapping) diceSimilarity(m *ASTMapping) float64 {
	src := k.src
	dst := k.dst
	common := commonNumber(src, dst, m)
	src_term := src.Term
	dst_term := dst.Term
	totalInst := len(src.Insts) + len(dst.Insts)
	isTermEqual := reflect.TypeOf(src_term) == reflect.TypeOf(dst_term)
	if totalInst == 0 { //src dst 只有term node的情况
		if isTermEqual {
			k.common = 1
			k.dice = 1
		} else {
			k.common = 0
			k.dice = 0
		}
		return k.dice
	} else {
		if isTermEqual {
			common += 1
			totalInst += 2
		}
		k.dice = float64(2*common) / float64(totalInst)
		k.common = common
		return k.dice
	}
}

func (k *Mapping) ndiceSimilarity(src_cfg, dst_cfg *ControlFlowGraph, m *ASTMapping) float64 {
	src := k.src
	dst := k.dst
	src_preds := src_cfg.pred[src]
	dst_preds := dst_cfg.pred[dst]
	pred_dices := float64(0)
	for _, src_pred := range src_preds {
		max := float64(0)
		for _, dst_pred := range dst_preds {
			tmp := Mapping{src_pred, dst_pred, 0, 0, 0}
			dice := tmp.diceSimilarity(m)
			if dice > max {
				max = dice
			}
		}
		pred_dices = pred_dices + max
	}
	src_sucss := src_cfg.successor[src]
	dst_sucss := dst_cfg.successor[dst]
	suc_dices := float64(0)
	for _, src_suc := range src_sucss {
		max := float64(0)
		for _, dst_suc := range dst_sucss {
			tmp := Mapping{src_suc, dst_suc, 0, 0, 0}
			dice := tmp.diceSimilarity(m)
			if dice > max {
				max = dice
			}
		}
		suc_dices = suc_dices + max
	}
	k.n_dice = pred_dices + suc_dices
	return k.n_dice
}

func FuncDiceSimilarity(src, dst *ir.Func, m *ASTMapping) (float64, map[*ir.Block]*ir.Block) {
	srcIgnored := make(map[*ir.Block]bool)
	dstIgnored := make(map[*ir.Block]bool)
	var ambiguousList Mappings
	Container_mapping := make(map[*ir.Block]*ir.Block) // map[src]dst is better
	src_cfg := createCFG(src)
	dst_cfg := createCFG(dst)

	//计算src和dst的指令数
	srcNumInst := 0
	dstNumInst := 0
	common_inst := 0

	for _, src_block := range src.Blocks {
		srcNumInst += len(src_block.Insts)
		//term node
		srcNumInst += 1
	}
	for _, dst_block := range dst.Blocks {
		dstNumInst += len(dst_block.Insts)
		//term nod
		dstNumInst += 1
	}

	// src 和 dst 中的block两两比较相似度
	for _, src_block := range src.Blocks {
		for _, dst_block := range dst.Blocks {
			ambiguousList = append(ambiguousList, Mapping{src_block, dst_block, 0, 0, 0})
		}
	}
	for i := range ambiguousList {
		ambiguousList[i].diceSimilarity(m)
	}
	for i := range ambiguousList {
		ambiguousList[i].ndiceSimilarity(src_cfg, dst_cfg, m)
	}
	sort.Sort(ambiguousList)
	//贪婪选取dice值最高的block组合
	for {
		if len(ambiguousList) == 0 {
			break
		}
		first := ambiguousList[0]
		ambiguousList = ambiguousList[1:]
		// if first.dice < 0.2 { //阈值
		// 	break
		// }
		_, src_ok := srcIgnored[first.src]
		_, dst_ok := dstIgnored[first.dst]
		if !(src_ok || dst_ok) {
			common_inst += first.common
			Container_mapping[first.src] = first.dst
			srcIgnored[first.src] = true
			dstIgnored[first.dst] = true
		}
	}
	return float64(common_inst*2) / float64(srcNumInst+dstNumInst), Container_mapping
}

type ASTMapping struct {
	src_dfs   map[*ir.Func]*DataFlow
	dst_dfs   map[*ir.Func]*DataFlow
	debugInfo map[string]metadata.Definition
	//Inst has multiMapping
	srcToDsts         map[*ir.Instruction]([]*ir.Instruction) //key是无序的，但value是有序的
	dstToSrcs         map[*ir.Instruction]([]*ir.Instruction)
	ambiguous_mapping map[*ir.Instruction]([]*ir.Instruction)

	src_func_table  map[string](*ir.Func)
	src_block_table map[string](*ir.Block)
	src_inst_table  map[*ir.Instruction]bool

	Func_mapping        map[*ir.Func]*ir.Func
	Block_mapping       map[*ir.Block]*ir.Block
	Instruction_mapping map[*ir.Instruction]*ir.Instruction
	//cotainer mapping for block map[src]dst
	Container_mapping map[*ir.Block]*ir.Block
}

func NewASTMapping() *ASTMapping {
	return &ASTMapping{
		src_dfs:           make(map[*ir.Func]*DataFlow),
		dst_dfs:           make(map[*ir.Func]*DataFlow),
		debugInfo:         make(map[string]metadata.Definition),
		srcToDsts:         make(map[*ir.Instruction]([]*ir.Instruction)),
		dstToSrcs:         make(map[*ir.Instruction]([]*ir.Instruction)),
		ambiguous_mapping: make(map[*ir.Instruction]([]*ir.Instruction)),

		src_func_table:  make(map[string]*ir.Func),
		src_block_table: make(map[string]*ir.Block),
		src_inst_table:  make(map[*ir.Instruction]bool),

		Func_mapping:        make(map[*ir.Func]*ir.Func),
		Block_mapping:       make(map[*ir.Block]*ir.Block),
		Instruction_mapping: make(map[*ir.Instruction]*ir.Instruction),
		Container_mapping:   make(map[*ir.Block]*ir.Block),
	}
}

func (m *ASTMapping) isSrcUnique(src *ir.Instruction) bool {
	return len(m.srcToDsts[src]) == 1
}

func (m *ASTMapping) isDstUnique(dst *ir.Instruction) bool {
	return len(m.dstToSrcs[dst]) == 1
}

func (m *ASTMapping) AstCompare(src, dst *ir.Module) []operation {

	marksForSrcFunc := make(map[*ir.Func]bool)
	marksForDstFunc := make(map[*ir.Func]bool)

	for _, src_f := range src.Funcs {
		m.src_func_table[(src_f.Hash())] = src_f
	}
	for _, dst_f := range dst.Funcs {
		if src_f, ok := m.src_func_table[(dst_f.Hash())]; ok {
			//命中
			// fmt.Printf("命中:%s\n", src_f.LLString())
			marksForSrcFunc[src_f] = true
			marksForDstFunc[dst_f] = true
			m.Func_mapping[dst_f] = src_f
		}
	}
	marksForSrcBlock := make(map[*ir.Block]bool)
	marksForDstBlock := make(map[*ir.Block]bool)
	for _, src_f := range src.Funcs {
		if !marksForSrcFunc[src_f] {
			for _, src_block := range src_f.Blocks {
				m.src_block_table[(src_block.Hash())] = src_block
			}
		}
	}
	for _, dst_f := range dst.Funcs {
		if !marksForDstFunc[dst_f] {
			for _, dst_block := range dst_f.Blocks {
				if src_block, ok := m.src_block_table[(dst_block.Hash())]; ok {
					//命中
					// fmt.Printf("命中:%s\n", src_block.LLString())
					marksForSrcBlock[src_block] = true
					marksForDstBlock[dst_block] = true
					m.Block_mapping[dst_block] = src_block
				}
			}
		}
	}
	marksForSrcInst := make(map[*ir.Instruction]bool)
	marksForDstInst := make(map[*ir.Instruction]bool)
	for _, src_f := range src.Funcs {
		if !marksForSrcFunc[src_f] {
			for _, src_block := range src_f.Blocks {
				if !marksForSrcBlock[src_block] {
					for k := range src_block.Insts {
						src_inst := &(src_block.Insts)[k]
						m.src_inst_table[src_inst] = false
					}
				}
			}
		}
	}

	for i, dst_f := range dst.Funcs {
		if !marksForDstFunc[dst_f] {
			src_f := src.Funcs[i] //Fixme：假设src 和dst Funcs是一一对应的
			m.src_dfs[src_f] = InstDependence(src_f, m.debugInfo)
			m.dst_dfs[dst_f] = InstDependence(dst_f, m.debugInfo)
			for _, dst_block := range dst_f.Blocks {
				if !marksForDstBlock[dst_block] {
					for k := range dst_block.Insts {
						dst_inst := &(dst_block.Insts)[k]
						for src_inst := range m.src_inst_table {
							if instCompare(src_inst, dst_inst, src_f, dst_f, m.src_dfs[src_f], m.dst_dfs[dst_f]) {
								//命中 multi_mapping
								marksForSrcInst[src_inst] = true
								marksForDstInst[dst_inst] = true
								m.srcToDsts[src_inst] = append(m.srcToDsts[src_inst], dst_inst)
								m.dstToSrcs[dst_inst] = append(m.dstToSrcs[dst_inst], src_inst)
							}
						}
					}
				}
			}
		}
	}
	//对于inst，可能存在多个候选的mapping，需要filter
	m.filterMapping()
	var scripts []operation
	// for dst, src := range m.Container_mapping {
	// 	script, _ := generateBlockDiff(src, dst, m, m.Container_mapping)
	// 	scripts = append(scripts, script...)
	// }
	return scripts
}

func (m *ASTMapping) filterMapping() {
	var ambiguousList Mappings
	ignored := make(map[*ir.Instruction]bool)
	for src_inst, dst_insts := range m.srcToDsts {
		isMappingUnique := false
		dst_inst := dst_insts[0]
		if m.isSrcUnique(src_inst) {
			if m.isDstUnique(dst_inst) {
				isMappingUnique = true
				m.Instruction_mapping[src_inst] = dst_inst
			}
		}
		_, ok := ignored[src_inst]
		if !(ok || isMappingUnique) {
			//src_insts是切片
			src_insts := m.dstToSrcs[dst_inst]
			for _, _src_inst := range src_insts {
				for _, _dst_inst := range dst_insts {
					ambiguousList = append(ambiguousList, Mapping{(*_src_inst).GetParent(), (*_dst_inst).GetParent(), 0, 0, 0})
				}
			}
			for _, src_inst := range src_insts {
				ignored[src_inst] = true
			}
		}
	}
	//Rank
	for i := range ambiguousList {
		ambiguousList[i].diceSimilarity(m)
	}
	sort.Sort(ambiguousList) //降序排列
	//Recovery block ambiguous mapping
	m.retainBestMapping(ambiguousList)

}

func (m *ASTMapping) retainBestMapping(ambiguousList Mappings) {
	srcIgnored := make(map[*ir.Block]bool)
	dstIgnored := make(map[*ir.Block]bool)
	for {
		if len(ambiguousList) == 0 {
			break
		}
		// fmt.Println(ambiguousList)
		first := ambiguousList[0]
		if first.dice < 0.5 { //阈值
			break
		}
		ambiguousList = ambiguousList[1:]
		_, src_ok := srcIgnored[first.src]
		_, dst_ok := dstIgnored[first.dst]
		if !(src_ok || dst_ok) {
			m.Container_mapping[first.src] = first.dst
			srcIgnored[first.src] = true
			dstIgnored[first.dst] = true
		}
	}

}

func (m *ASTMapping) FuncCompare(mod *ir.Module, CFilePth string) []operation {
	var goodFuncsTable []*ir.Func
	var badFuncsTable []*ir.Func

	//TODO:get debug infor from Module
	for _, md := range mod.MetadataDefs {
		// fmt.Printf("%s %s\n", md.Ident(), md.LLString())
		m.debugInfo[md.Ident()] = md
	}

	//Step 1 找出good 和bad func

	for _, f := range mod.Funcs { //Fixme f可能是值复制
		if f.Linkage != enum.LinkageLinkOnceODR {
			if ok := strings.Contains(f.GlobalIdent.Name(), "good"); ok {
				goodFuncsTable = append(goodFuncsTable, f)
			}
			if ok := strings.Contains(f.GlobalIdent.Name(), "bad"); ok {
				badFuncsTable = append(badFuncsTable, f)
			}
		}
	}

	//Step2 提取good 和 bad 函数

	for i, bad_f := range badFuncsTable {
		for j, good_f := range goodFuncsTable {
			dice, container_mapping := FuncDiceSimilarity(bad_f, good_f, m)
			if dice != float64(1) && dice > 0.7 {
				fmt.Println(dice)
				fmt.Println(bad_f.GlobalIdent.Name())
				fmt.Println(good_f.GlobalIdent.Name())
				fmt.Println(container_mapping)

				// //generate bad graph sample and good graph sample
				bad_g := NewGraph()
				bad_g.Func2Graph(bad_f, m.src_dfs[bad_f]) //need update tag becasue it's bad sample, need to be changed.
				good_g := NewGraph()
				good_g.Func2Graph(good_f, m.dst_dfs[good_f]) //dont need update tag because it's good sample

				script, content, graph_diff := generateFuncDiff(bad_f, good_f, m, container_mapping, bad_g)
				// fmt.Println(graph_diff.String())
				fmt.Println(content)
				fmt.Println(script)
				// fmt.Println(i, j)
				generateSamples(mod, CFilePth, bad_f, good_f, i, j, bad_g, good_g, graph_diff)

			}
		}
	}

	return nil
}

func generateSamples(mod *ir.Module, CFilePth string, bad_f, good_f *ir.Func, i, j int, bad_g, good_g *Graph, graph_diff *gographviz.Graph) {
	//Step 0 setup filepath config
	filenameBase := mod.SourceFilename
	var fileSuffix string
	fileSuffix = path.Ext(filenameBase)
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameBase, fileSuffix)
	// prefix := "data" //root directory
	CFilePth = strings.TrimSuffix(CFilePth, ".ll") + fileSuffix

	//Step3 打印到文件，生成good sample和bad sample

	// badFilename := fmt.Sprintf("%s/%s_%d%d_bad.ll", prefix, filenameOnly, i, j)
	// goodFilename := fmt.Sprintf("%s/%s_%d%d_good.ll", prefix, filenameOnly, i, j)
	// cFilename := fmt.Sprintf("%s/%s_%d%d_cfile%s", prefix, filenameOnly, i, j, fileSuffix)
	badGraphFile := fmt.Sprintf("%s/%s_%d%d_bad.txt", "data/bad", filenameOnly, i, j)
	goodGraphFile := fmt.Sprintf("%s/%s_%d%d_good.txt", "data/good", filenameOnly, i, j)
	diffDotGraphFile := fmt.Sprintf("%s/%s_%d%d_diff.dot", "data/diff", filenameOnly, i, j)
	//copy cFile from srouce
	bad_g.PrintToFile(badGraphFile)
	good_g.PrintToFile(goodGraphFile)
	WriteToFile(diffDotGraphFile, graph_diff.String())
	// srcCFile, err := os.Open(CFilePth)
	// if err != nil {
	// 	panic(err)
	// }
	// defer srcCFile.Close()
	// dstCFile, err := os.OpenFile(cFilename, os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer dstCFile.Close()
	// io.Copy(dstCFile, srcCFile)

	// //remove and keep funcs
	// badFuncs := []*ir.Func{bad_f}
	// goodFuncs := []*ir.Func{good_f}
	// for _, f := range mod.Funcs {
	// 	if !(strings.Contains(f.GlobalIdent.Name(), "CWE") || strings.Contains(f.GlobalIdent.Name(), "good") || strings.Contains(f.GlobalIdent.Name(), "bad") || strings.Contains(f.GlobalIdent.Name(), "main")) {
	// 		badFuncs = append(badFuncs, f)
	// 		goodFuncs = append(goodFuncs, f)
	// 	}
	// }
	// //DO NOT remove mod.metadata !

	// mod.Funcs = badFuncs
	// //generate bad sample
	// badContent := []byte(fmt.Sprintln(mod))
	// bad_err := ioutil.WriteFile(badFilename, badContent, 0644)
	// if bad_err != nil {
	// 	panic(bad_err)
	// }

	// mod.Funcs = goodFuncs
	// //generate good sample
	// goodContent := []byte(fmt.Sprintln(mod))
	// good_err := ioutil.WriteFile(goodFilename, goodContent, 0644)
	// if good_err != nil {
	// 	panic(good_err)
	// }
}
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
