// This example parses an LLVM IR assembly file and pretty-prints the data types
// of the parsed module to standard output.
package main

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

func main() {
	// Parse the LLVM IR assembly file `foo.ll`.
	m, err := asm.ParseFile("test.ll")
	// m1, err := asm.ParseFile("test1.ll")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Pretty-print the data types of the parsed LLVM IR module.
	pretty.Println(m)

	// locate func @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badD2Ev
	const funcName = "_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badD2Ev"
	// f, ok := findFunc(m, funcName)

	// if !ok {
	// 	log.Fatalf("unable to locate function %q", funcName)
	// }
	// modify func:
	// g, err := dupFunc(f, "g")
	// if err != nil {
	// 	log.Fatalf("unable to clone function %q; %v", f.Name(), err)
	// }

	// fmt.Println(asSha256(f), asSha256(g))
	// fmt.Println(asSha256(m), asSha256(m1))
	// add new fucntion
	// fmt.Println(f)
	// fmt.Println(m)
}

func findFunc(m *ir.Module, funcName string) (*ir.Func, bool) {
	for _, f := range m.Funcs {
		if f.Name() == funcName {
			// remove fun
			// m.Funcs = append(m.Funcs[:i], m.Funcs[i+1:]...)
			return f, true
		}
	}
	return nil, false
}

func modifyFunc(g *ir.Func) *ir.Func {
	for i, block := range g.Blocks {
		blockName := fmt.Sprintf("block_%d", i)
		block.SetName(blockName)
		for _, inst := range block.Insts {
			switch inst := inst.(type) {
			case *ir.InstAdd:
				// swap X and Y.
				inst.X, inst.Y = inst.Y, inst.X
			}
		}
	}
	return g
}

func dupFunc(origFunc *ir.Func, newFuncName string) (*ir.Func, error) {
	// Attempt using deepcopy.
	//
	//newFunc := deepcopy.Copy(origFunc).(*ir.Func)

	// Attempt using copier.
	//
	//newFunc := new(ir.Func)
	//if err := copier.Copy(newFunc, origFunc); err != nil {
	//	return nil, err
	//}

	// Attempt using copystructure.
	//v, err := copystructure.Copy(origFunc)
	//if err != nil {
	//	return nil, err
	//}
	//newFunc := v.(*ir.Func)

	// Use json marshal/unmarshal as workaround for deep copy.
	// This should work but may be slow.
	//newFunc := new(ir.Func)
	//if err := deepcopyWorkaround(newFunc, origFunc); err != nil {
	//	return nil, err
	//}

	// Attempt using getlantern/deepcopy.
	//newFunc := new(ir.Func)
	//if err := deepcopy.Copy(newFunc, origFunc); err != nil {
	//	return nil, err
	//}

	newFunc := cloneFunc(origFunc)

	// newFunc.SetName(newFuncName)
	return newFunc, nil
}

func cloneFunc(origFunc *ir.Func) *ir.Func {
	newFunc := new(ir.Func)
	*newFunc = *origFunc
	newFunc.Blocks = make([]*ir.Block, len(origFunc.Blocks))
	for i, origBlock := range origFunc.Blocks {
		newBlock := cloneBlock(origBlock)
		newFunc.Blocks[i] = newBlock
	}
	return newFunc
}

func cloneBlock(origBlock *ir.Block) *ir.Block {
	newBlock := new(ir.Block)
	*newBlock = *origBlock
	newBlock.Insts = make([]ir.Instruction, len(origBlock.Insts))
	for j, origInst := range origBlock.Insts {
		newInst := cloneInst(origInst)
		newBlock.Insts[j] = newInst
	}
	newTerm := cloneTerm(origBlock.Term)
	newBlock.Term = newTerm
	return newBlock
}

func cloneInst(origInst ir.Instruction) ir.Instruction {
	// var newInst ir.Instruction
	// switch origInst := origInst.(type) {
	// case *ir.InstAdd:
	// 	v := new(ir.InstAdd)
	// 	*v = *origInst
	// 	newInst = v
	// default:
	// 	panic(fmt.Errorf("support for instruction %T not yet implemented", origInst))
	// }
	newInst := origInst
	return newInst
}

func cloneTerm(origTerm ir.Terminator) ir.Terminator {
	var newTerm ir.Terminator
	switch origTerm := origTerm.(type) {
	case *ir.TermRet:
		v := new(ir.TermRet)
		*v = *origTerm
		newTerm = v
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", origTerm))
	}
	return newTerm
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
