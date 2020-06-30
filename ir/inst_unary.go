package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Unary instructions ] --------------------------------------------------

// ~~~ [ fneg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFNeg is an LLVM IR fneg instruction.
type InstFNeg struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operand.
	X value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewFNeg returns a new fneg instruction based on the given operand.
func NewFNeg(x value.Value) *InstFNeg {
	inst := &InstFNeg{X: x}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFNeg) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstFNeg) GetParent() *Block {
	return inst.Parent
}
func (inst *InstFNeg) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstFNeg) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fneg' FastMathFlags=FastMathFlag* X=TypeValue Metadata=(',' MetadataAttachment)+?
func (inst *InstFNeg) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fneg")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " %s", inst.X)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstFNeg) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "fneg %s, %s", inst.Type(), inst.X.Type())
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
