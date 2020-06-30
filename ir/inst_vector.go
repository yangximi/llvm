package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractElement is an LLVM IR extractelement instruction.
type InstExtractElement struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vector.
	X value.Value
	// Element index.
	Index value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func NewExtractElement(x, index value.Value) *InstExtractElement {
	inst := &InstExtractElement{X: x, Index: index}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstExtractElement) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstExtractElement) GetParent() *Block {
	return inst.Parent
}

func (inst *InstExtractElement) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstExtractElement) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.X.Type().(*types.VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", inst.X.Type()))
		}
		inst.Typ = t.ElemType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'extractelement' X=TypeValue ',' Index=TypeValue Metadata=(',' MetadataAttachment)+?
func (inst *InstExtractElement) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "extractelement %s, %s", inst.X, inst.Index)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstExtractElement) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "extractelement %s,%s, %s", inst.Type(), inst.X.Type(), inst.Index.Type())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertElement is an LLVM IR insertelement instruction.
type InstInsertElement struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vector.
	X value.Value
	// Element to insert.
	Elem value.Value
	// Element index.
	Index value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ *types.VectorType
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func NewInsertElement(x, elem, index value.Value) *InstInsertElement {
	inst := &InstInsertElement{X: x, Elem: elem, Index: index}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstInsertElement) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstInsertElement) GetParent() *Block {
	return inst.Parent
}
func (inst *InstInsertElement) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstInsertElement) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.X.Type().(*types.VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", inst.X.Type()))
		}
		inst.Typ = t
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'insertelement' X=TypeValue ',' Elem=TypeValue ',' Index=TypeValue
// Metadata=(',' MetadataAttachment)+?
func (inst *InstInsertElement) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "insertelement %s, %s, %s", inst.X, inst.Elem, inst.Index)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstInsertElement) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "insertelement %s, %s, %s, %s", inst.Type(), inst.X.Type(), inst.Elem.Type(), inst.Index.Type())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShuffleVector is an LLVM IR shufflevector instruction.
type InstShuffleVector struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vectors.
	X, Y value.Value
	// Shuffle mask.
	Mask value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ *types.VectorType
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask value.Value) *InstShuffleVector {
	inst := &InstShuffleVector{X: x, Y: y, Mask: mask}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstShuffleVector) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}
func (inst *InstShuffleVector) GetParent() *Block {
	return inst.Parent
}
func (inst *InstShuffleVector) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstShuffleVector) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		xType, ok := inst.X.Type().(*types.VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", inst.X.Type()))
		}
		maskType, ok := inst.Mask.Type().(*types.VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", inst.Mask.Type()))
		}
		inst.Typ = types.NewVector(maskType.Len, xType.ElemType)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'shufflevector' X=TypeValue ',' Y=TypeValue ',' Mask=TypeValue
// Metadata=(',' MetadataAttachment)+?
func (inst *InstShuffleVector) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "shufflevector %s, %s, %s", inst.X, inst.Y, inst.Mask)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstShuffleVector) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "shufflevector %s, %s, %s, %s", inst.Type(), inst.X.Type(), inst.Y.Type(), inst.Mask.Type())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
