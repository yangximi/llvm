package ir

import (
	"github.com/llir/llvm/ir/value"
)

// --- [ Unary instructions ] --------------------------------------------------

// ~~~ [ fneg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFNeg appends a new fneg instruction to the basic block based on the given
// operand.
func (block *Block) NewFNeg(x value.Value) *InstFNeg {
	inst := NewFNeg(x)
	block.Insts = append(block.Insts, inst)
	inst.Parent = block
	return inst
}
