package asm

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/llir/llvm/ir"
)

type Vocabulary struct {
	NodeType map[string]int
	hasField map[string]int
}

func NewVocabulary() *Vocabulary {
	new := &Vocabulary{
		NodeType: make(map[string]int),
		hasField: make(map[string]int),
	}
	// initial set UNKOWN type at first
	new.NodeType["UNKOWN"] = len(new.NodeType)
	new.hasField["UNKOWN"] = len(new.hasField)
	return new
}

func (vocab *Vocabulary) GetInstType() map[string]int {
	return vocab.NodeType
}

func (vocab *Vocabulary) GetFiedType() map[string]int {
	return vocab.hasField
}

func (vocab *Vocabulary) GetVocabulary() []string {
	dict := make([]string, len(vocab.NodeType))
	for k, v := range vocab.NodeType {
		dict[v] = k
	}
	return dict
}

func (vocab *Vocabulary) ToString() string {
	buf := &strings.Builder{}
	for n_type, n_index := range vocab.NodeType {
		buf.WriteString(fmt.Sprintf("%d:%s\n", n_index, n_type))
	}
	for field_type, field_index := range vocab.hasField {
		buf.WriteString(fmt.Sprintf("%d:%s\n", field_index, field_type))
	}
	return buf.String()
}

func (vocab *Vocabulary) PrintToFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(vocab.ToString()), n)
		defer f.Close()
	}
	return err
}

func (vocab *Vocabulary) CollectTypeFromFiles(files []string) {
	for _, file := range files {
		mod, err := ParseFile(file)

		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, f := range mod.Funcs {
			for _, block := range f.Blocks {
				block_type := reflect.TypeOf(block).Elem()
				if _, ok := vocab.NodeType[block_type.String()]; !ok {
					vocab.NodeType[block_type.String()] = len(vocab.NodeType)
				}
				term := reflect.TypeOf(block.Term).Elem()
				if _, ok := vocab.NodeType[term.String()]; !ok {
					vocab.NodeType[term.String()] = len(vocab.NodeType)
				}
				for i := 0; i < term.NumField(); i++ {
					fieldType := term.Field(i)
					if _, ok := vocab.hasField[fieldType.Name]; !ok {
						vocab.hasField[fieldType.Name] = len(vocab.hasField)
					}
				}

				for _, inst := range block.Insts {
					inst_type := reflect.TypeOf(inst).Elem()
					if _, ok := vocab.NodeType[inst_type.String()]; !ok {
						vocab.NodeType[inst_type.String()] = len(vocab.NodeType)
					}
					for i := 0; i < inst_type.NumField(); i++ {
						fieldType := inst_type.Field(i)
						if _, ok := vocab.hasField[fieldType.Name]; !ok {
							vocab.hasField[fieldType.Name] = len(vocab.hasField)
						}
					}
				}
			}
		}
	}
}

func (vocab *Vocabulary) IndexOfInst(inst *ir.Instruction) int {

	inst_type := reflect.TypeOf(*inst).Elem()
	if index, ok := vocab.NodeType[inst_type.String()]; ok {
		return index
	} else {
		return 0
	}
}
