package asm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func terminatorCompare(src_block, dst_block *ir.Block, src_f, dst_f *ir.Func, src_df, dst_df *DataFlow, container_mapping map[*ir.Block]*ir.Block) operation {

	src := src_block.Term
	dst := dst_block.Term
	if reflect.TypeOf(src) != reflect.TypeOf(dst) {
		return UPDATE
	} else {
		switch src_term := src.(type) {
		case *ir.TermRet:
			dst_term := dst.(*ir.TermRet)

			// Case1: Compare *ir.Instrution
			if src_term.X != nil && dst_term.X != nil {
				src_defInst, ok1 := src_df.InstTable[src_term.X.Ident()]
				dst_defInst, ok2 := dst_df.InstTable[dst_term.X.Ident()]
				if ok1 && ok2 {
					if !instCompare(src_defInst, dst_defInst, src_f, dst_f, src_df, dst_df) { //引用的inst相同
						return UPDATE
					}
				}
			}
			//default
			return MOVE
		case *ir.TermBr:
			dst_term := dst.(*ir.TermBr)
			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.Target.Ident()] != nil && dst_df.BlockTable[dst_term.Target.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.Target.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.Target.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermCondBr:
			dst_term := dst.(*ir.TermCondBr)

			// Case1: Compare *ir.Instrution
			if src_df.InstTable[src_term.Cond.Ident()] != nil && dst_df.InstTable[dst_term.Cond.Ident()] != nil {
				src_defInst := src_df.InstTable[src_term.Cond.Ident()]
				dst_defInst := dst_df.InstTable[dst_term.Cond.Ident()]
				if !instCompare(src_defInst, dst_defInst, src_f, dst_f, src_df, dst_df) { //引用的inst相同
					return UPDATE
				}
				// Condition equal
			}
			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.TargetTrue.Ident()] != nil && dst_df.BlockTable[dst_term.TargetTrue.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.TargetTrue.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.TargetTrue.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
				//TargetTrue equal
			}
			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.TargetFalse.Ident()] != nil && dst_df.BlockTable[dst_term.TargetFalse.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.TargetFalse.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.TargetFalse.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermSwitch:
			dst_term := dst.(*ir.TermSwitch)

			// Case1: Compare *ir.Instrution
			if src_df.InstTable[src_term.X.Ident()] != nil && dst_df.InstTable[dst_term.X.Ident()] != nil {
				src_defInst := src_df.InstTable[src_term.X.Ident()]
				dst_defIsnt := dst_df.InstTable[dst_term.X.Ident()]
				if !instCompare(src_defInst, dst_defIsnt, src_f, dst_f, src_df, dst_df) { //引用的inst相同
					return UPDATE
				}
				// Control equal
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.TargetDefault.Ident()] != nil && dst_df.BlockTable[dst_term.TargetDefault.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.TargetDefault.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.TargetDefault.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}

			//Case4 Compare []*ir.Block
			for i, src_case := range src_term.Cases {
				src_target := src_df.BlockTable[src_case.Target.Ident()]          //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.Cases[i].Target.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermIndirectBr:
			dst_term := dst.(*ir.TermIndirectBr)

			// Case3 Compare constant
			if src_term.Addr != dst_term.Addr {
				return UPDATE
			}

			//Case4 Compare []*ir.Block
			for i, src_target := range src_term.ValidTargets {
				src_target := src_df.BlockTable[src_target.Ident()]               //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.ValidTargets[i].Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			//ALL equal
			return MOVE
		case *ir.TermInvoke: //Fixme：该指令有ssa赋值，先忽略
			dst_term := dst.(*ir.TermInvoke)

			if src_term.Invokee.Ident() != dst_term.Invokee.Ident() { //Fixme 直接比较字符串相等
				return UPDATE
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.NormalRetTarget.Ident()] != nil && dst_df.BlockTable[dst_term.NormalRetTarget.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.NormalRetTarget.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.NormalRetTarget.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.ExceptionRetTarget.Ident()] != nil && dst_df.BlockTable[dst_term.ExceptionRetTarget.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.ExceptionRetTarget.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.ExceptionRetTarget.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}

			// Case5: compare string
			src_buf := &strings.Builder{}
			for _, arg := range src_term.Args {
				src_buf.WriteString(arg.String())
			}
			dst_buf := &strings.Builder{}
			for _, arg := range dst_term.Args {
				dst_buf.WriteString(arg.String())
			}
			if src_buf.String() != dst_buf.String() {
				return UPDATE
			}
			//ALL equal
			return MOVE
		case *ir.TermCallBr:
			dst_term := dst.(*ir.TermCallBr)

			//Callee
			if src_term.Callee.Ident() != dst_term.Callee.Ident() { //FIXME
				return UPDATE
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.NormalRetTarget.Ident()] != nil && dst_df.BlockTable[dst_term.NormalRetTarget.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.NormalRetTarget.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.NormalRetTarget.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}

			// Case5: compare string
			src_buf := &strings.Builder{}
			for _, arg := range src_term.Args {
				src_buf.WriteString(arg.String())
			}
			dst_buf := &strings.Builder{}
			for _, arg := range dst_term.Args {
				dst_buf.WriteString(arg.String())
			}
			if src_buf.String() != dst_buf.String() {
				return UPDATE
			}
			//ALL equal
			return MOVE
		case *ir.TermResume:
			dst_term := dst.(*ir.TermResume)

			// Case1: Compare *ir.Instrution
			if src_df.InstTable[src_term.X.Ident()] != nil && dst_df.InstTable[dst_term.X.Ident()] != nil {
				src_defInst := src_df.InstTable[src_term.X.Ident()]
				dst_defIsnt := dst_df.InstTable[dst_term.X.Ident()]
				if !instCompare(src_defInst, dst_defIsnt, src_f, dst_f, src_df, dst_df) { //引用的inst相同
					return UPDATE
				}
			}
			//default
			return MOVE
		case *ir.TermCatchSwitch:
			dst_term := dst.(*ir.TermCatchSwitch)

			//ParentPad
			if src_term.ParentPad.Ident() != dst_term.ParentPad.Ident() { //Fixme 应该在Instdependence里写term的def-use
				return UPDATE
			}

			//Case4 Compare []*ir.Block
			for i, src_target := range src_term.Handlers {
				src_target := src_df.BlockTable[src_target.Ident()]           //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.Handlers[i].Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.DefaultUnwindTarget.Ident()] != nil && dst_df.BlockTable[dst_term.DefaultUnwindTarget.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.DefaultUnwindTarget.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.DefaultUnwindTarget.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermCatchRet:
			dst_term := dst.(*ir.TermCatchRet)

			//CatchPad
			if src_term.CatchPad.Ident() != dst_term.CatchPad.Ident() { //Fixme 应该在Instdependence里写term的def-use
				return UPDATE
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.Target.Ident()] != nil && dst_df.BlockTable[dst_term.Target.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.Target.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.Target.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermCleanupRet:
			dst_term := dst.(*ir.TermCleanupRet)

			//CleanupPad
			if src_term.CleanupPad.Ident() != dst_term.CleanupPad.Ident() { //Fixme 应该在Instdependence里写term的def-use
				return UPDATE
			}

			// Case2: Compare *ir.Block
			if src_df.BlockTable[src_term.UnwindTarget.Ident()] != nil && dst_df.BlockTable[dst_term.UnwindTarget.Ident()] != nil { //for example %42
				src_target := src_df.BlockTable[src_term.UnwindTarget.Ident()] //*ir.Block
				dst_target := dst_df.BlockTable[dst_term.UnwindTarget.Ident()] //*ir.Block
				if container_mapping[src_target] != dst_target {
					return UPDATE
				}
			}
			return MOVE
		case *ir.TermUnreachable:
			return MOVE

		default:
			return UPDATE
		}
	}

}

func calculateHash(inst *ir.Instruction, df *DataFlow) string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s", (*inst).Hash())
	if df.UseChain[inst] != nil {
		for _, use_inst := range df.UseChain[inst] {
			fmt.Fprintf(buf, "%s", calculateHash(use_inst, df))
		}
	}
	return buf.String()
}
func instDiceSimilary(src, dst *ir.Instruction, src_df, dst_df *DataFlow) float64 {
	common := 0
	srcNumInst := 0
	dstNumInst := 0
	// srcIgnored := make(map[*ir.Instruction]bool)
	// dstIgnored := make(map[*ir.Instruction]bool)
	if src_df.UseChain[src] != nil {
		srcNumInst = len(src_df.UseChain[src])
	}
	if dst_df.UseChain[dst] != nil {
		dstNumInst = len(dst_df.UseChain[dst])
	}
	if srcNumInst > 0 && dstNumInst > 0 {

	}
	return float64(common * 2)
}

func ambCompare(src, dst *ir.Instruction, src_df, dst_df *DataFlow) bool { //For now, only use for search update action
	if reflect.TypeOf(*src) != reflect.TypeOf(*dst) {
		return false
	} else if len(src_df.DefChain[src]) != len(dst_df.DefChain[dst]) { //Fixme maybe never reached this branch
		return false
	} else if len(src_df.DefChain[src]) == 0 {
		switch src_inst := (*src).(type) {
		case *ir.InstAlloca:
			dst_inst := (*dst).(*ir.InstAlloca)
			if !src_inst.ElemType.Equal(dst_inst.ElemType) {
				return false
			}
			if src_inst.Align != dst_inst.Align {
				return false
			}
			//TODO: add check def-use chain
			dst_df.Index[dst] = true
			return true
		default:
			return false
		}
	} else {
		for i, dst_inst := range dst_df.DefChain[dst] {
			if !dst_df.Index[dst_inst] || !ambCompare(src_df.DefChain[src][i], dst_inst, src_df, dst_df) { //false
				return false
			}
			dst_df.Index[dst] = true
			return true
		}
		return false
	}

}

func instCompare(src, dst *ir.Instruction, src_f, dst_f *ir.Func, src_df, dst_df *DataFlow) bool {
	if reflect.TypeOf(*src) != reflect.TypeOf(*dst) {
		return false
	} else if len(src_df.DefChain[src]) != len(dst_df.DefChain[dst]) { //Fixme maybe never reached this branch
		return false
	} else if len(dst_df.DefChain[dst]) == 0 { //没有引用变量的指令比较

		if src_value, ok := (*src).(*ir.InstAlloca); ok { //单独处理变量申请的inst指令比较，根据debug信息中的变量名是否一致来比较
			dst_value := (*dst).(*ir.InstAlloca)
			ssaSrcValue := fmt.Sprintf("%s:%s", src_f, src_value.String())
			ssaDstValue := fmt.Sprintf("%s:%s", dst_f, dst_value.String())
			if src_df.VariableTable[ssaSrcValue] == dst_df.VariableTable[ssaDstValue] {
				dst_df.Index[dst] = true
				return true
			} else if src_value.GetTypeDefs() != nil && dst_value.GetTypeDefs() != nil { //自定义变量类型
				if types.IsPointer(src_value.ElemType) && types.IsPointer(dst_value.ElemType) {
					if (src_value.ElemType).(*types.PointerType).ElemType.LLString() == (dst_value.ElemType).(*types.PointerType).ElemType.LLString() {
						dst_df.Index[dst] = true
						return true
					}
				} else if src_value.ElemType.LLString() == dst_value.ElemType.LLString() {
					dst_df.Index[dst] = true
					return true
				}
			}
			return false
		}
		if src_value, ok := (*src).(*ir.InstCall); ok && strings.Contains(src_value.Callee.Ident(), "@llvm.dbg.declare") {
			dst_value := (*dst).(*ir.InstCall)
			if src_value.Callee == dst_value.Callee { //ignore two @llvm.dbg.declare call compare, return ture
				return true
			}
		}
		if (*src).Hash() == (*dst).Hash() {
			dst_df.Index[dst] = true
			return true
		}
		return false
	} else { //len(src_df.Use[src]) == len(dst_df.Use[dst]) >0 引用变量的指令比较

		//Step1. 检查当前指令所处的系统状态(跟该指令有数据流关系)是否相同
		for i, dst_inst := range dst_df.DefChain[dst] {
			//递归前溯 变量定义是否相同
			if !dst_df.Index[dst_inst] || !instCompare(src_df.DefChain[src][i], dst_inst, src_f, dst_f, src_df, dst_df) { //false
				// if !dst_df.Index[dst_inst] || !ambCompare(src_df.DefChain[src][i], dst_inst, src_df, dst_df) { //false
				return false
			}

		}
		//Step2. 检查两个指令改变的系统状态量是否相同
		// fmt.Println("Src:"+(*dst).LLString(), "Dst:"+(*src).LLString())
		if (*src).Hash() == (*dst).Hash() {
			dst_df.Index[dst] = true
			return true
		}
		return false
	}
}
