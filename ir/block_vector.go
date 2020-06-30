package ir

import (
	"github.com/llir/llvm/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractElement appends a new extractelement instruction to the basic block
// based on the given vector and element index.
func (block *Block) NewExtractElement(x, index value.Value) *InstExtractElement {
	inst := NewExtractElement(x, index)
	block.Insts = append(block.Insts, inst)
	inst.Parent = block
	return inst
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertElement appends a new insertelement instruction to the basic block
// based on the given vector, element and element index.
func (block *Block) NewInsertElement(x, elem, index value.Value) *InstInsertElement {
	inst := NewInsertElement(x, elem, index)
	block.Insts = append(block.Insts, inst)
	inst.Parent = block
	return inst
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShuffleVector appends a new shufflevector instruction to the basic block
// based on the given vectors and shuffle mask.
func (block *Block) NewShuffleVector(x, y, mask value.Value) *InstShuffleVector {
	inst := NewShuffleVector(x, y, mask)
	block.Insts = append(block.Insts, inst)
	inst.Parent = block
	return inst
}
