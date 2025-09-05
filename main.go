package main

import (
	"Abbas-Askari/interpreter-v2/compiler"
	"Abbas-Askari/interpreter-v2/lexer"
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/vm"
	"fmt"
	"os"
)

func main() {
	filename := "/home/abbas/repos/interpreter-v2/prototypes.test.lox"
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
	nativeFunctions := object.GetNativeFunctions()
	for _, fun := range nativeFunctions {
		compiler.DefineConstant(fun.Name, fun)
	}

	function, constants := compiler.Compile(statements)
	fmt.Println(function, constants)

	parser.Decompile(function, constants)

	for i, c := range constants {
		if fn, ok := c.(object.Function); ok {
			fmt.Println("Function at constant index:", i)
			parser.Decompile(fn, constants)
		}
	}

	vm := vm.NewVM(function, constants, nativeFunctions)
	fmt.Println("----------Output----------")
	vm.Run()
	fmt.Println("--------------------------")

	fmt.Println("Done!")
}
