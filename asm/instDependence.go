package asm

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
)

type DataFlow struct {
	Index      map[*ir.Instruction]bool
	Define     map[string]*ir.Instruction
	BlockTable map[string]*ir.Block
	//Use indicate which definition instructions it use
	Use           map[*ir.Instruction][]*ir.Instruction
	DefChain      map[*ir.Instruction][]*ir.Instruction
	VariableTable map[string]string
}

func NewDataFlow() *DataFlow {
	return &DataFlow{
		Index:         make(map[*ir.Instruction]bool),
		Define:        make(map[string]*ir.Instruction),
		BlockTable:    make(map[string]*ir.Block),
		Use:           make(map[*ir.Instruction][]*ir.Instruction),
		DefChain:      make(map[*ir.Instruction][]*ir.Instruction),
		VariableTable: make(map[string]string),
	}
}

func InstDependence(f *ir.Func, dbg map[string]metadata.Definition) *DataFlow {
	df := NewDataFlow()
	for _, block := range f.Blocks {
		//Index ir.Block
		df.BlockTable[block.LocalIdent.Ident()] = block //for example %42

		//For instruction
		for k := range block.Insts {
			inst := &(block.Insts)[k]
			switch inst_c := (*inst).(type) {
			// inst aggregate 0-2:
			case *ir.InstExtractValue:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstInsertValue:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst binary 2-12:
			case *ir.InstAdd:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFAdd:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstSub:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFSub:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstMul:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFMul:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstUDiv:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstSDiv:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFDiv:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstURem:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstSRem:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFRem:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst bitwise 14-20
			case *ir.InstShl:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstLShr:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstAShr:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstAnd:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstOr:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstXor:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst conversion 20-33
			case *ir.InstTrunc:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstZExt:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstSExt:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFPTrunc:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFPExt:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFPToUI:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFPToSI:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstUIToFP:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstSIToFP:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstPtrToInt:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstIntToPtr:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstBitCast:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstAddrSpaceCast:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.From.Ident()] != nil {
					def_inst := df.Define[inst_c.From.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst memory 33-40
			case *ir.InstAlloca:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
			case *ir.InstLoad:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.Src.Ident()] != nil {
					def_inst := df.Define[inst_c.Src.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}

			case *ir.InstStore:
				df.Index[inst] = false
				if df.Define[inst_c.Src.Ident()] != nil {
					def_inst := df.Define[inst_c.Src.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Dst.Ident()] != nil {
					def_inst := df.Define[inst_c.Dst.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFence: //Fixme
				df.Index[inst] = false
			case *ir.InstCmpXchg: //Fixme
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.Ptr.Ident()] != nil {
					def_inst := df.Define[inst_c.Ptr.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Cmp.Ident()] != nil {
					def_inst := df.Define[inst_c.Cmp.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.New.Ident()] != nil {
					def_inst := df.Define[inst_c.New.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstAtomicRMW:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.Dst.Ident()] != nil {
					def_inst := df.Define[inst_c.Dst.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstGetElementPtr:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.Src.Ident()] != nil {
					def_inst := df.Define[inst_c.Src.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst other
			case *ir.InstICmp:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFCmp:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstPhi:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				for _, inc := range inst_c.Incs {
					if df.Define[inc.X.Ident()] != nil {
						def_inst := df.Define[inc.X.Ident()]
						df.Use[inst] = append(df.Use[inst], def_inst)
						df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
					}
				}
			case *ir.InstSelect: //Fixme
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.ValueTrue.Ident()] != nil {
					def_inst := df.Define[inst_c.ValueTrue.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.ValueFalse.Ident()] != nil {
					def_inst := df.Define[inst_c.ValueFalse.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstFreeze:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstCall: //Fixme
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.Callee.Ident()] != nil {
					def_inst := df.Define[inst_c.Callee.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				for _, arg := range inst_c.Args {
					if df.Define[arg.Ident()] != nil {
						def_inst := df.Define[arg.Ident()]
						df.Use[inst] = append(df.Use[inst], def_inst)
						df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
					}
				}
				//TODO:based on debug info call @llvm.dbg.declare, get source variable information
				if strings.Contains(inst_c.Callee.Ident(), "llvm.dbg.declare") {
					ssaValue := fmt.Sprintf("%s:%s", f, inst_c.Args[0].Ident())
					// fmt.Println(ssaValue)
					debugIndex := inst_c.Args[1].Ident()
					if md, ok := dbg[debugIndex]; ok {
						switch value := md.(type) {
						case *metadata.DILocalVariable:
							df.VariableTable[ssaValue] = value.Name
						case *metadata.DIGlobalVariable:
							df.VariableTable[ssaValue] = value.Name
						default:
							continue
						}

					}
				}
			case *ir.InstVAArg:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.ArgList.Ident()] != nil {
					def_inst := df.Define[inst_c.ArgList.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstLandingPad:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				for _, clause := range inst_c.Clauses {
					if df.Define[clause.X.Ident()] != nil {
						def_inst := df.Define[clause.X.Ident()]
						df.Use[inst] = append(df.Use[inst], def_inst)
						df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
					}
				}
			case *ir.InstCleanupPad:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.ParentPad.Ident()] != nil {
					def_inst := df.Define[inst_c.ParentPad.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				for _, arg := range inst_c.Args {
					if df.Define[arg.Ident()] != nil {
						def_inst := df.Define[arg.Ident()]
						df.Use[inst] = append(df.Use[inst], def_inst)
						df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
					}
				}
			// inst unary
			case *ir.InstFNeg:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			// inst vector
			case *ir.InstExtractElement:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstInsertElement:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			case *ir.InstShuffleVector:
				df.Index[inst] = false
				df.Define[inst_c.Ident()] = inst
				if df.Define[inst_c.X.Ident()] != nil {

					def_inst := df.Define[inst_c.X.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
				if df.Define[inst_c.Y.Ident()] != nil {

					def_inst := df.Define[inst_c.Y.Ident()]
					df.Use[inst] = append(df.Use[inst], def_inst)
					df.DefChain[def_inst] = append(df.DefChain[def_inst], inst)
				}
			}
		}
		//Terminator node
		// switch term := block.Term.(type) {
		// case *ir.TermInvoke:
		// 	df.Define[term.Ident()] =
		// 	if df.Define[term.X.Ident()] != nil {
		// 		def_inst := df.Define[term.X.Ident()]
		// 		df.DefChain[def_inst] = append(df.DefChain[def_inst])

		// 	}
		// }

	}
	return df
}
