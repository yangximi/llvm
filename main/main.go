// This example parses an LLVM IR assembly file and pretty-prints the data types
// of the parsed module to standard output.
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

func testcase() (*ir.Module, []string) {
	//Testcase
	// testcase := "/Users/ys/Project/llvm/main/CWE121_Stack_Based_Buffer_Overflow__src_char_declare_cat_41.ll"
	testcase := "/Users/ys/Desktop/CWE121_Stack_Based_Buffer_Overflow__src_char_declare_cpy_05.ll"
	mod, err := asm.ParseFile(testcase)
	if err != nil {
		panic(err)
	}
	// mapping := asm.NewASTMapping()
	// mapping.FuncCompare(mod, testcase)
	// dst, _ := asm.ParseFile(testcase)
	mapping := asm.NewASTMapping()
	filePath := "/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d.ll"
	mapping.FuncCompare(mod, filePath)
	// script := mapping.AstCompare(src, dst)
	// fmt.Println(script)
	files := []string{testcase}
	return mod, files
}

func testgraphviz() {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	output := graph.String()
	fmt.Println(output)
}

func main() {
	_, _ = testcase()
	// testgraphviz()
	//Step1 get samples
	directory := "/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow"
	files, err := GetAllFiles(directory)
	if err != nil {
		panic(err)
	}
	GenerateSamples(files)
	GenerateVocabulary(files)

}

func GenerateSamples(files []string) {
	for _, file := range files {
		mod, err := asm.ParseFile(file)
		fmt.Println("\n" + file + "\n")
		if err != nil {
			fmt.Println(err)
			continue
		}
		mapping := asm.NewASTMapping()
		mapping.FuncCompare(mod, file)
	}
}

func GenerateVocabulary(files []string) *asm.Vocabulary {
	vocab := asm.NewVocabulary()
	vocab.CollectTypeFromFiles(files)
	//Print to File
	filepath := "data/vocab.txt"
	vocab.PrintToFile(filepath)
	fmt.Println(len(vocab.NodeType))
	fmt.Println(vocab)
	return vocab
}

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			// GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".ll")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
