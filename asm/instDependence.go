package asm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

type DataFlow struct {
	Index         map[*ir.Instruction]bool              // use for matcher.go::ambCompare()
	IdentTable    map[*ir.Instruction]string            //use for map Inst to IdentID
	InstTable     map[string]*ir.Instruction            // Define map IdentID to inst
	BlockTable    map[string]*ir.Block                  // map block's IdentID to itself
	DefChain      map[*ir.Instruction][]*ir.Instruction // Use map from use-inst to def-inst, which is definition instruction.
	UseChain      map[*ir.Instruction][]*ir.Instruction // DefChain map from def-inst to use-inst
	VariableTable map[string]string                     // map ssaValue to its variable name in source code
	VariableTrace map[string][]*ir.Instruction
	TermToInst    map[*ir.Terminator][]*ir.Instruction
	TermToBlock   map[*ir.Terminator][]*ir.Block
}

func NewDataFlow() *DataFlow {
	return &DataFlow{
		Index:         make(map[*ir.Instruction]bool),
		IdentTable:    make(map[*ir.Instruction]string),
		InstTable:     make(map[string]*ir.Instruction),
		BlockTable:    make(map[string]*ir.Block),
		DefChain:      make(map[*ir.Instruction][]*ir.Instruction),
		UseChain:      make(map[*ir.Instruction][]*ir.Instruction),
		VariableTable: make(map[string]string),
		VariableTrace: make(map[string][]*ir.Instruction),
		TermToInst:    make(map[*ir.Terminator][]*ir.Instruction),
		TermToBlock:   make(map[*ir.Terminator][]*ir.Block),
	}
}

func InstDependence(f *ir.Func, dbg map[string]metadata.Definition) *DataFlow {
	df := NewDataFlow()
	for _, block := range f.Blocks {
		df.BlockTable[block.LocalIdent.Ident()] = block //Index ir.Block by it's IdentID, for example %42
		for k := range block.Insts {                    //For instruction
			inst := &(block.Insts)[k]
			switch inst_c := (*inst).(type) {
			// inst aggregate 0-2:
			case *ir.InstExtractValue:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstInsertValue:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst binary 2-12:
			case *ir.InstAdd:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFAdd:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstSub:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFSub:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstMul:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFMul:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstUDiv:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstSDiv:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {
					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFDiv:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstURem:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstSRem:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFRem:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst bitwise 14-20
			case *ir.InstShl:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstLShr:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstAShr:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstAnd:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstOr:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstXor:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst conversion 20-33
			case *ir.InstTrunc:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstZExt:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstSExt:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFPTrunc:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFPExt:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFPToUI:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFPToSI:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstUIToFP:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstSIToFP:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstPtrToInt:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstIntToPtr:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstBitCast:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstAddrSpaceCast:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.From.Ident()] != nil {
					def_inst := df.InstTable[inst_c.From.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst memory 33-40
			case *ir.InstAlloca:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
			case *ir.InstLoad:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.Src.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Src.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}

			case *ir.InstStore:
				df.Index[inst] = false
				if df.InstTable[inst_c.Src.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Src.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Dst.Ident()] != nil {
					df.VariableTrace[inst_c.Dst.Ident()] = append(df.VariableTrace[inst_c.Dst.Ident()], inst)
					def_inst := df.InstTable[inst_c.Dst.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFence: //Fixme
				df.Index[inst] = false
			case *ir.InstCmpXchg: //Fixme
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.Ptr.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Ptr.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Cmp.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Cmp.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.New.Ident()] != nil {
					def_inst := df.InstTable[inst_c.New.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstAtomicRMW:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.Dst.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Dst.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstGetElementPtr:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.Src.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Src.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst other
			case *ir.InstICmp:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFCmp:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstPhi:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				for _, inc := range inst_c.Incs {
					if df.InstTable[inc.X.Ident()] != nil {
						def_inst := df.InstTable[inc.X.Ident()]
						df.DefChain[inst] = append(df.DefChain[inst], def_inst)
						df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
					}
				}
			case *ir.InstSelect: //Fixme
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.ValueTrue.Ident()] != nil {
					def_inst := df.InstTable[inst_c.ValueTrue.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.ValueFalse.Ident()] != nil {
					def_inst := df.InstTable[inst_c.ValueFalse.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstFreeze:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstCall: //Fixme
				//some call instructions' type is void
				if !inst_c.Type().Equal(types.Void) {
					df.Index[inst] = false
					df.IdentTable[inst] = inst_c.Ident()
					df.InstTable[inst_c.Ident()] = inst
					df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				}
				if df.InstTable[inst_c.Callee.Ident()] != nil {
					def_inst := df.InstTable[inst_c.Callee.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				for _, arg := range inst_c.Args {
					if df.InstTable[arg.Ident()] != nil {
						def_inst := df.InstTable[arg.Ident()]
						df.VariableTrace[arg.Ident()] = append(df.VariableTrace[arg.Ident()], inst)
						df.DefChain[inst] = append(df.DefChain[inst], def_inst)
						df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
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
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.ArgList.Ident()] != nil {
					def_inst := df.InstTable[inst_c.ArgList.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstLandingPad:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				for _, clause := range inst_c.Clauses {
					if df.InstTable[clause.X.Ident()] != nil {
						def_inst := df.InstTable[clause.X.Ident()]
						df.DefChain[inst] = append(df.DefChain[inst], def_inst)
						df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
					}
				}
			case *ir.InstCleanupPad:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.ParentPad.Ident()] != nil {
					def_inst := df.InstTable[inst_c.ParentPad.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				for _, arg := range inst_c.Args {
					if df.InstTable[arg.Ident()] != nil {
						def_inst := df.InstTable[arg.Ident()]
						df.DefChain[inst] = append(df.DefChain[inst], def_inst)
						df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
					}
				}
			// inst unary
			case *ir.InstFNeg:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			// inst vector
			case *ir.InstExtractElement:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstInsertElement:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			case *ir.InstShuffleVector:
				df.Index[inst] = false
				df.IdentTable[inst] = inst_c.Ident()
				df.InstTable[inst_c.Ident()] = inst
				df.VariableTrace[inst_c.Ident()] = append(df.VariableTrace[inst_c.Ident()], inst)
				if df.InstTable[inst_c.X.Ident()] != nil {

					def_inst := df.InstTable[inst_c.X.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
				if df.InstTable[inst_c.Y.Ident()] != nil {

					def_inst := df.InstTable[inst_c.Y.Ident()]
					df.DefChain[inst] = append(df.DefChain[inst], def_inst)
					df.UseChain[def_inst] = append(df.UseChain[def_inst], inst)
				}
			}
		}
		//Terminator node
		p_term := &block.Term
		switch term := block.Term.(type) {
		case *ir.TermRet:
			if term.X != nil {
				if Inst, ok := df.InstTable[term.X.Ident()]; ok {
					df.TermToInst[p_term] = append(df.TermToInst[p_term], Inst)
				}
			}
		case *ir.TermBr:
			if df.BlockTable[term.Target.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.Target.Ident()])
			}
		case *ir.TermCondBr:
			// Case1: Compare *ir.Instrution
			if df.InstTable[term.Cond.Ident()] != nil {
				df.TermToInst[p_term] = append(df.TermToInst[p_term], df.InstTable[term.Cond.Ident()])
			}

			// Case2: Compare *ir.Block
			if df.BlockTable[term.TargetTrue.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.TargetTrue.Ident()])
			}

			// Case2: Compare *ir.Block
			if df.BlockTable[term.TargetFalse.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.TargetFalse.Ident()])
			}
		case *ir.TermSwitch:
			// Case1: Compare *ir.Instrution
			if df.InstTable[term.X.Ident()] != nil {
				df.TermToInst[p_term] = append(df.TermToInst[p_term], df.InstTable[term.X.Ident()])
			}

			// Case2: Compare *ir.Block
			if df.BlockTable[term.TargetDefault.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.TargetDefault.Ident()])
			}

			//Case4 Compare []*ir.Block
			for _, src_case := range term.Cases {
				target := df.BlockTable[src_case.Target.Ident()] //*ir.Block
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], target)
			}

		case *ir.TermIndirectBr:
			//Case4 Compare []*ir.Block
			for _, src_target := range term.ValidTargets {
				src_target := df.BlockTable[src_target.Ident()] //*ir.Block
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], src_target)
			}

		case *ir.TermInvoke: //Fixme：该指令有ssa赋值，先忽略

			// Case2: Compare *ir.Block
			if df.BlockTable[term.NormalRetTarget.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.NormalRetTarget.Ident()])
			}

			// Case2: Compare *ir.Block
			if df.BlockTable[term.ExceptionRetTarget.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.ExceptionRetTarget.Ident()])
			}
		case *ir.TermCallBr:
			// Case2: Compare *ir.Block
			if df.BlockTable[term.NormalRetTarget.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.NormalRetTarget.Ident()])
			}

		case *ir.TermResume:
			// Case1: Compare *ir.Instrution
			if df.InstTable[term.X.Ident()] != nil {
				df.TermToInst[p_term] = append(df.TermToInst[p_term], df.InstTable[term.X.Ident()])
			}
		case *ir.TermCatchSwitch:
			//Case4 Compare []*ir.Block
			for _, src_target := range term.Handlers {
				src_target := df.BlockTable[src_target.Ident()] //*ir.Block
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], src_target)
			}

			// Case2: Compare *ir.Block
			if df.BlockTable[term.DefaultUnwindTarget.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.DefaultUnwindTarget.Ident()])
			}
		case *ir.TermCatchRet:

			// Case2: Compare *ir.Block
			if df.BlockTable[term.Target.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.Target.Ident()])
			}
		case *ir.TermCleanupRet:
			// Case2: Compare *ir.Block
			if df.BlockTable[term.UnwindTarget.Ident()] != nil {
				df.TermToBlock[p_term] = append(df.TermToBlock[p_term], df.BlockTable[term.UnwindTarget.Ident()])
			}
		case *ir.TermUnreachable:

		}

	}
	// opt for memory dependence: concatenate inst trace
	for _, block := range f.Blocks {
		df.BlockTable[block.LocalIdent.Ident()] = block //Index ir.Block by it's IdentID, for example %42
		for k := range block.Insts {                    //For instruction
			inst := &(block.Insts)[k]
			switch inst_c := (*inst).(type) {
			case *ir.InstLoad:
				if reflect.TypeOf(inst_c.ElemType).String() == "*types.PointerType" {
					df.VariableTrace[inst_c.Src.Ident()] = append(df.VariableTrace[inst_c.Src.Ident()], df.VariableTrace[inst_c.Ident()]...)
				}
			case *ir.InstGetElementPtr:
				df.VariableTrace[inst_c.Src.Ident()] = append(df.VariableTrace[inst_c.Src.Ident()], df.VariableTrace[inst_c.Ident()]...)
			case *ir.InstBitCast:
				df.VariableTrace[inst_c.From.Ident()] = append(df.VariableTrace[inst_c.From.Ident()], df.VariableTrace[inst_c.Ident()]...)
			case *ir.InstStore:
				df.VariableTrace[inst_c.Dst.Ident()] = append(df.VariableTrace[inst_c.Dst.Ident()], df.VariableTrace[inst_c.Src.Ident()]...)
			}
		}
	}
	// print variable-instTrace maps
	// for key, trace := range df.VariableTrace {
	// 	fmt.Println(key)
	// 	for _, inst := range trace {
	// 		fmt.Println((*inst).LLString())
	// 	}
	// }

	return df
}
