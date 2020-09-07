package ir

import (
	"fmt"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShl is an LLVM IR shl instruction.
type InstShl struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Overflow flags.
	OverflowFlags []enum.OverflowFlag
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewShl returns a new shl instruction based on the given operands.
func NewShl(x, y value.Value) *InstShl {
	inst := &InstShl{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstShl) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstShl) GetParent() *Block {
	return inst.Parent
}
func (inst *InstShl) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstShl) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'shl' OverflowFlags=OverflowFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstShl) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("shl")
	for _, flag := range inst.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstShl) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "shl %s", inst.Type())

	return buf.String()
}
func (inst *InstShl) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "shl_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "shl_y"})
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLShr is an LLVM IR lshr instruction.
type InstLShr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewLShr returns a new lshr instruction based on the given operands.
func NewLShr(x, y value.Value) *InstLShr {
	inst := &InstLShr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstLShr) GetParent() *Block {
	return inst.Parent
}
func (inst *InstLShr) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstLShr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'lshr' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstLShr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("lshr")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstLShr) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "lshr %s", inst.Type())

	return buf.String()
}
func (inst *InstLShr) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "lshl_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "lshl_y"})
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAShr is an LLVM IR ashr instruction.
type InstAShr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewAShr returns a new ashr instruction based on the given operands.
func NewAShr(x, y value.Value) *InstAShr {
	inst := &InstAShr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAShr) GetParent() *Block {
	return inst.Parent
}
func (inst *InstAShr) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstAShr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'ashr' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstAShr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("ashr")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstAShr) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "ashr %s", inst.Type())

	return buf.String()
}
func (inst *InstAShr) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "ashr_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "ashr_y"})
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAnd is an LLVM IR and instruction.
type InstAnd struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewAnd returns a new and instruction based on the given operands.
func NewAnd(x, y value.Value) *InstAnd {
	inst := &InstAnd{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAnd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstAnd) GetParent() *Block {
	return inst.Parent
}
func (inst *InstAnd) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstAnd) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'and' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstAnd) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "and %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstAnd) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "and %s", inst.Type())

	return buf.String()
}
func (inst *InstAnd) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "and_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "and_y"})
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstOr is an LLVM IR or instruction.
type InstOr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewOr returns a new or instruction based on the given operands.
func NewOr(x, y value.Value) *InstOr {
	inst := &InstOr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstOr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstOr) GetParent() *Block {
	return inst.Parent
}
func (inst *InstOr) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstOr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'or' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstOr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "or %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstOr) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "or %s", inst.Type())

	return buf.String()
}
func (inst *InstOr) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "or_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "or_y"})
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstXor is an LLVM IR xor instruction.
type InstXor struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata
	//parent
	Parent *Block
}

// NewXor returns a new xor instruction based on the given operands.
func NewXor(x, y value.Value) *InstXor {
	inst := &InstXor{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstXor) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

func (inst *InstXor) GetParent() *Block {
	return inst.Parent
}
func (inst *InstXor) SetParent(b *Block) {
	inst.Parent = b
}

//func (inst *) equal(other *i) {return false}

// Type returns the type of the instruction.
func (inst *InstXor) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'xor' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstXor) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "xor %s, %s", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
func (inst *InstXor) Hash() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "xor %s", inst.Type())

	return buf.String()
}
func (inst *InstXor) ToDotGraph(graph *gographviz.Graph, prefix string) {
	cluster_f := Add_quotation_marks(inst.Parent.Parent.Ident(), "cluster_"+prefix)
	x_id := Add_quotation_marks(inst.X.Ident(), prefix)
	y_id := Add_quotation_marks(inst.Y.Ident(), prefix)
	dst_id := Add_quotation_marks(inst.Ident(), prefix)

	graph.AddNode(cluster_f, x_id, nil)
	graph.AddNode(cluster_f, y_id, nil)
	graph.AddNode(cluster_f, dst_id, nil)
	graph.AddEdge(x_id, dst_id, true, map[string]string{"label": "xor_x"})
	graph.AddEdge(y_id, dst_id, true, map[string]string{"label": "xor_y"})
}
