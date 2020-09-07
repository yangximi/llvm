package asm

import "github.com/llir/llvm/ir"

type ControlFlowGraph struct {
	pred      map[*ir.Block][]*ir.Block
	successor map[*ir.Block][]*ir.Block
}

func NewControlFlowGraph() *ControlFlowGraph {
	return &ControlFlowGraph{
		pred:      make(map[*ir.Block][]*ir.Block),
		successor: make(map[*ir.Block][]*ir.Block),
	}
}

func (cfg *ControlFlowGraph) addNode(b *ir.Block) {
	successors := b.Term.Succs()
	cfg.successor[b] = successors //fixme maybe should copy
	for _, suc_b := range successors {
		cfg.pred[suc_b] = append(cfg.pred[suc_b], b)
	}
}

func createCFG(f *ir.Func) *ControlFlowGraph {
	cfg := NewControlFlowGraph()
	for _, block := range f.Blocks {
		//Add node (basic block) to graph
		cfg.addNode(block)
	}
	return cfg
}
