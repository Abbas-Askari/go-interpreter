package main

import (
	"Abbas-Askari/interpreter-v2/compiler"
	"Abbas-Askari/interpreter-v2/lexer"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/vm"
	"fmt"
	"os"
)

func main() {
	filename := "/home/abbas/repos/interpreter-v2/for.test.lox"
	// if len(os.Args) > 1 {
	// 	filename = os.Args[1]
	// }
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("File not found!\nUsage: run [filepath]"))
	}
	// fmt.Println(string(fileContent))

	tokens := lexer.Tokenize(string(fileContent))

	fmt.Println(tokens)

	p := parser.NewParser(tokens)
	statements := p.Parse()
	fmt.Println("------------AST------------")
	fmt.Println(statements)
	fmt.Println("---------------------------")

	compiler := compiler.NewCompiler()

	stream, constants := compiler.Compile(statements)
	fmt.Println(stream, constants)

	parser.Decompile(stream, constants)

	vm := vm.NewVM(stream, constants)
	fmt.Println("----------Output----------")
	vm.Run()
	fmt.Println("--------------------------")

	fmt.Println("Done!")
}
