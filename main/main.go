// This example parses an LLVM IR assembly file and pretty-prints the data types
// of the parsed module to standard output.
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/llir/llvm/asm"
)

func main() {

	directory := "/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/"
	files, err := GetAllFiles(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
		mod, _ := asm.ParseFile(file)
		mapping := asm.NewASTMapping()
		mapping.FuncCompare(mod, file)
	}

	//Testcase
	// filePath := "/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d.ll"
	// src, _ := asm.ParseFile("/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d.ll")
	// dst, _ := asm.ParseFile("/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d.ll")
	// mapping := asm.NewASTMapping()
	// mapping.FuncCompare(src, filePath)
	// script := mapping.AstCompare(src, tgt)
	// fmt.Println(script)
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
