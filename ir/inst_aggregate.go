package ir

import (
	"fmt"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractValue is an LLVM IR extractvalue instruction.
type InstExtractValue struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Aggregate value.
	X value.Value // array or struct
	// Element indices.
	Indices []uint64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewExtractValue returns a new extractvalue instruction based on the given
// aggregate value and indicies.
func NewExtractValue(x value.Value, indices ...uint64) *InstExtractValue {
	inst := &InstExtractValue{X: x, Indices: indices}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstExtractValue) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstExtractValue) GetParent() *Block {
	return inst.Parent
}
func (inst *InstExtractValue) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstExtractValue) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = aggregateElemType(inst.X.Type(), inst.Indices)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'extractvalue' X=TypeValue Indices=(',' UintLit)+ Metadata=(',' MetadataAttachment)+?
func (inst *InstExtractValue) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "extractvalue %s", inst.X)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstExtractValue) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "extractvalue %s", inst.Type())
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}

	return buf.String()
}
func (inst *InstExtractValue) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	indice := &strings.Builder{}
	for _, index := range inst.Indices {
		fmt.Fprintf(indice, ", %d", index)
	}
	dst_id := Add_quotation_marks(inst.Ident(), prefix)
	indice_id := Add_quotation_marks(indice.String(), prefix)
	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, indice_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "extractvalue"})
	graph.AddEdge(indice_id, dst_id, true, map[string]string{"label": "indice"})
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertValue is an LLVM IR insertvalue instruction.
type InstInsertValue struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Aggregate value.
	X value.Value // array or struct
	// Element to insert.
	Elem value.Value
	// Element indices.
	Indices []uint64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewInsertValue returns a new insertvalue instruction based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem value.Value, indices ...uint64) *InstInsertValue {
	elemType := aggregateElemType(x.Type(), indices)
	if !elemType.Equal(elem.Type()) {
		panic(fmt.Errorf("insertvalue elem type mismatch, expected %v, got %v", elemType, elem.Type()))
	}
	inst := &InstInsertValue{X: x, Elem: elem, Indices: indices}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstInsertValue) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstInsertValue) GetParent() *Block {
	return inst.Parent
}
func (inst *InstInsertValue) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstInsertValue) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'insertvalue' X=TypeValue ',' Elem=TypeValue Indices=(',' UintLit)+ Metadata=(',' MetadataAttachment)+?
func (inst *InstInsertValue) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "insertvalue %s, %s", inst.X, inst.Elem)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstInsertValue) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "insertvalue %s, %s", inst.Type(), inst.Elem.Type())
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}

	return buf.String()
}
func (inst *InstInsertValue) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "shl_x"})
}

// ### [ Helper functions ] ####################################################

// aggregateElemType returns the element type at the position in the aggregate
// type specified by the given indices.
func aggregateElemType(t types.Type, indices []uint64) types.Type {
	// Base case.
	if len(indices) == 0 {
		return t
	}
	switch t := t.(type) {
	case *types.ArrayType:
		return aggregateElemType(t.ElemType, indices[1:])
	case *types.StructType:
		return aggregateElemType(t.Fields[indices[0]], indices[1:])
	case *types.PointerType:
		return aggregateElemType(t.ElemType, indices[1:])
	default:
		panic(fmt.Errorf("support for aggregate type %T not yet implemented", t))
	}
}
