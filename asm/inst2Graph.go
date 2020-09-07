package asm

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
)

type EdgeType uint

const (
	DEFCHAIN EdgeType = 1 //from def to use
	USECHAIN          = 2 //from use to def
	EXCUTION          = 3 //excution order
	CHILD             = 4 //inst
	JUMPTO            = 5
)

type Node struct {
	Nodeindex int
	Nodetype  string
	Tag       operation
	Content   string
	Fields    []NodeField
}

func NewNode(index int, ty, content string, f []NodeField) Node {
	return Node{
		Nodeindex: index,
		Nodetype:  ty,
		Tag:       MOVE,
		Content:   content,
		Fields:    f,
	}
}

func (n *Node) SetTag(op operation) {
	n.Tag = op
}

type NodeField struct {
	Label string
	Type  string
	Value string
}

func (n *Node) PrintField() string {
	buf := &strings.Builder{}
	for _, field := range n.Fields {
		buf.WriteString(fmt.Sprintf("%s:%s;", field.Label, field.Value))
	}
	return buf.String()
}

func GetFieldsOfNode(n interface{}) []NodeField {
	var f []NodeField
	switch t := (n).(type) {
	case *ir.Block:

	case *ir.Instruction:
		inst_type := reflect.TypeOf(*t).Elem()
		inst_value := reflect.ValueOf(*t).Elem()
		for i := 0; i < inst_type.NumField(); i++ {
			fieldType := inst_type.Field(i)
			fieldValue := inst_value.Field(i)
			if fieldType.Name != "Parent" {
				f = append(f, NodeField{fieldType.Name, fieldType.Type.String(), fmt.Sprint(fieldValue)})
			}
		}
	case *ir.Terminator:
		term_type := reflect.TypeOf(*t).Elem()
		term_value := reflect.ValueOf(*t).Elem()
		for i := 0; i < term_type.NumField(); i++ {
			fieldType := term_type.Field(i)
			fieldValue := term_value.Field(i)
			if fieldType.Name != "Parent" {
				f = append(f, NodeField{fieldType.Name, fieldType.Type.String(), fmt.Sprint(fieldValue)})
			}
		}
	}
	return f
}

type Edge struct {
	Src_index int
	Src_type  string
	Dst_index int
	Dst_type  string
	Edge_type EdgeType
}

func NewEdge(src_index, dst_index int, src_ty, dst_ty string, e_ty EdgeType) Edge {
	return Edge{
		Src_index: src_index,
		Src_type:  src_ty,
		Dst_index: dst_index,
		Dst_type:  dst_ty,
		Edge_type: e_ty,
	}
}

type Graph struct {
	Nodes        []Node
	Edges        []Edge
	indexOfInst  map[*ir.Instruction]int
	indexOfTerm  map[*ir.Terminator]int
	indexOfBlock map[*ir.Block]int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:        []Node{},
		Edges:        []Edge{},
		indexOfInst:  make(map[*ir.Instruction]int),
		indexOfTerm:  make(map[*ir.Terminator]int),
		indexOfBlock: make(map[*ir.Block]int),
	}
}

func (g *Graph) ToString() string {
	buf := &strings.Builder{}
	buf.WriteString(fmt.Sprintf("? %d\n", len(g.Nodes)))
	for _, node := range g.Nodes {
		buf.WriteString(fmt.Sprintf("[%s] %d:%s; %s|%s\n", node.Tag.String(), node.Nodeindex, node.Nodetype, node.Content, node.PrintField()))
	}
	for _, edge := range g.Edges {
		buf.WriteString(fmt.Sprintf("%d:%s,%d,%d:%s\n", edge.Src_index, edge.Src_type, edge.Edge_type, edge.Dst_index, edge.Dst_type))
	}
	return buf.String()
}

func (g *Graph) addNode(index int, ty, content string, fields []NodeField) {
	g.Nodes = append(g.Nodes, NewNode(index, ty, content, fields))
}

func (g *Graph) addBlockNode(b *ir.Block) {
	b_ty := reflect.TypeOf(b).Elem()
	b_index := len(g.Nodes)
	g.indexOfBlock[b] = b_index
	g.addNode(b_index, b_ty.String(), b.String(), GetFieldsOfNode(b)) //add BlockNode Fixme: use block's hash() as its content
}

func (g *Graph) addBlock(b *ir.Block, df *DataFlow) {
	g.addBlockNode(b) //add BlockNode

	for k := range b.Insts {
		g.addInstNode(&(b.Insts)[k], df) //add InstNode in Block
	}

	g.addTermNode(&(b.Term), df) //add TermNode in Block
}

func (g *Graph) addInstNode(inst *ir.Instruction, df *DataFlow) {
	inst_ty := reflect.TypeOf(*inst).Elem().String()
	inst_index := len(g.Nodes)
	g.indexOfInst[inst] = inst_index
	content_buf := &strings.Builder{}
	ident := df.IdentTable[inst]

	for _, def_inst := range df.VariableTrace[ident] {
		def_inst_ty := reflect.TypeOf(*def_inst).Elem().String()
		content_buf.WriteString(def_inst_ty + " ")
	}
	content_buf.WriteString(";")
	for _, use_inst := range df.UseChain[inst] {
		use_inst_ty := reflect.TypeOf(*use_inst).Elem().String()
		content_buf.WriteString(use_inst_ty + " ")
	}
	content := strings.Replace(content_buf.String(), "\n", "", -1) //remove “\n”
	// content := strings.Replace((*inst).LLString(), "\n", "", -1) //remove “\n”
	g.addNode(inst_index, inst_ty, content, GetFieldsOfNode(inst))
}

func (g *Graph) addTermNode(term *ir.Terminator, df *DataFlow) {
	term_ty := reflect.TypeOf(*term).Elem().String()
	term_index := len(g.Nodes)
	g.indexOfTerm[term] = term_index
	content_buf := &strings.Builder{}
	for _, use_inst := range df.TermToInst[term] {
		use_inst_ty := reflect.TypeOf(*use_inst).Elem().String()
		content_buf.WriteString(use_inst_ty)
	}
	for _, b := range df.TermToBlock[term] {
		content_buf.WriteString(b.String())
	}
	content := strings.Replace(content_buf.String(), "\n", "", -1) //remove “\n”
	// content := strings.Replace((*term).LLString(), "\n", "", -1) //remove “\n”
	g.addNode(term_index, term_ty, content, GetFieldsOfNode(term))
}

func (g *Graph) addEdge(e Edge) {
	g.Edges = append(g.Edges, e)
}

func (g *Graph) GetInstIndex(inst *ir.Instruction) int {
	return g.indexOfInst[inst]
}

func (g *Graph) GetBlockIndex(b *ir.Block) int {
	return g.indexOfBlock[b]
}

func (g *Graph) GetTermIndex(term *ir.Terminator) int {
	return g.indexOfTerm[term]
}

func (g *Graph) GetNodeType(index int) string {
	return g.Nodes[index].Nodetype
}

func (g *Graph) UpdateInstTag(inst *ir.Instruction, op operation, df *DataFlow) {
	switch op {
	case INSERT:
		g.addInstNode(inst, df)
	}
	index := g.GetInstIndex(inst)
	p_node := &g.Nodes[index]
	p_node.SetTag(op)
}

func (g *Graph) UpdateBlockTag(b *ir.Block, op operation, df *DataFlow) {
	switch op {
	case INSERT:
		g.addBlockNode(b) //insts and term nodes will be added when call their own updatetag function
	}
	index := g.GetBlockIndex(b)
	p_node := &g.Nodes[index]
	p_node.SetTag(op) //Update BlockNode

	for k := range b.Insts {
		inst := &(b.Insts)[k]
		g.UpdateInstTag(inst, op, df) //Update InstNode in Block
	}
	g.UpdateTermTag(&b.Term, op, df) //Update TermNode in Block
}

func (g *Graph) UpdateTermTag(term *ir.Terminator, op operation, df *DataFlow) {
	switch op {
	case INSERT:
		g.addTermNode(term, df)
	}
	index := g.GetTermIndex(term)
	p_node := &g.Nodes[index]
	p_node.SetTag(op)
}

func (g *Graph) PrintToFile(fileName string) {
	data := []byte(g.ToString())
	ioutil.WriteFile(fileName, data, 0644)
	// f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Println("file create failed. err: " + err.Error())
	// } else {
	// 	n, _ := f.Seek(0, os.SEEK_END)
	// 	_, err = f.WriteAt([]byte(g.ToString()), n)
	// 	defer f.Close()
	// }
	// return err
}

func (g *Graph) Func2Graph(f *ir.Func, df *DataFlow) *Graph {
	for _, b := range f.Blocks {
		g.addBlock(b, df)
	}
	//Step3: link node by edge
	g.drawEdges(f, df)
	return g
}

func (g *Graph) Module2Graph(mod *ir.Module) {
	// Step1:get debuginfo
	debugInfo := make(map[string]metadata.Definition)
	for _, md := range mod.MetadataDefs {
		// fmt.Printf("%s %s\n", md.Ident(), md.LLString())
		debugInfo[md.Ident()] = md
	}

	// Step2:flatten inst
	for _, f := range mod.Funcs {
		df := InstDependence(f, debugInfo)
		g.Func2Graph(f, df)
	}

}

func (g *Graph) drawEdges(f *ir.Func, df *DataFlow) {
	var pre_index int            // record last excution index include basic block
	for _, b := range f.Blocks { // link block node with its child inst by bidirectional edge
		b_index := g.GetBlockIndex(b)
		pre_index = b_index
		for k := range b.Insts {
			inst := &(b.Insts)[k]
			inst_index := g.GetInstIndex(inst)

			child_edge := NewEdge(b_index, inst_index, g.GetNodeType(b_index), g.GetNodeType(inst_index), CHILD)
			g.addEdge(child_edge)

			exc_edge := NewEdge(pre_index, inst_index, g.GetNodeType(pre_index), g.GetNodeType(inst_index), EXCUTION)
			g.addEdge(exc_edge)

			pre_index = inst_index
		}
		term := &b.Term
		term_index := g.GetTermIndex(term)

		exc_edge := NewEdge(pre_index, term_index, g.GetNodeType(pre_index), g.GetNodeType(term_index), EXCUTION)
		g.addEdge(exc_edge)

		for _, suc_block := range (*term).Succs() { //link term node to succesor block
			suc_block_index := g.GetBlockIndex(suc_block)

			term_block_j_edge := NewEdge(term_index, suc_block_index, g.GetNodeType(term_index), g.GetNodeType(suc_block_index), JUMPTO)
			g.addEdge(term_block_j_edge)

			term_block_e_edge := NewEdge(term_index, suc_block_index, g.GetNodeType(term_index), g.GetNodeType(suc_block_index), EXCUTION)
			g.addEdge(term_block_e_edge)

		}
	}

	for defInst, useChain := range df.UseChain { //link defInst with its useInst by bidirectional edge
		for _, useInst := range useChain {
			def_index := g.GetInstIndex(defInst)
			use_index := g.GetInstIndex(useInst)

			def_edge := NewEdge(def_index, use_index, g.GetNodeType(def_index), g.GetNodeType(use_index), DEFCHAIN)
			g.addEdge(def_edge)

			use_edge := NewEdge(use_index, def_index, g.GetNodeType(use_index), g.GetNodeType(def_index), USECHAIN)
			g.addEdge(use_edge)
		}
	}

	for term, insts := range df.TermToInst {
		for _, defInst := range insts {
			inst_index := g.GetInstIndex(defInst)
			term_index := g.GetTermIndex(term)

			term_inst_edge := NewEdge(term_index, inst_index, g.GetNodeType(term_index), g.GetNodeType(inst_index), USECHAIN)
			g.addEdge(term_inst_edge)

			inst_term_edge := NewEdge(inst_index, term_index, g.GetNodeType(inst_index), g.GetNodeType(term_index), DEFCHAIN)
			g.addEdge(inst_term_edge)
		}
	}

}
