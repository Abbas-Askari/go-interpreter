package main

import (
	"Abbas-Askari/interpreter-v2/colors"
	"Abbas-Askari/interpreter-v2/compiler"
	"Abbas-Askari/interpreter-v2/lexer"
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/vm"
	"fmt"
	"os"
)

func runFile(filename string, debug bool) *object.Map {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("File not found!\nUsage: run [filepath]"))
	}
	// fmt.Println(string(fileContent))

	tokens := lexer.Tokenize(string(fileContent))

	if debug {
		fmt.Println(tokens)
	}

	p := parser.NewParser(tokens)
	statements := p.Parse()
	if debug {
		fmt.Println("------------AST------------")
		fmt.Println(statements)
		fmt.Println("---------------------------")
	}

	for _, stmt := range statements {
		if imp, ok := stmt.(*parser.ImportDeclaration); ok {
			if debug {
				fmt.Println("Importing module:", imp.Module.Literal)
			}
			modulePath := "./" + imp.Module.Literal + ".lox"
			if _, err := os.Stat(modulePath); os.IsNotExist(err) {
				panic(fmt.Sprintf("Module not found: %s", modulePath))
			}
			imp.Exports = runFile(modulePath, debug)
		}
	}

	compiler := compiler.NewCompiler(filename)

	globals := vm.GetNativeFunctions()
	compiler.DefineConstant("exports", object.Map{})
	compiler.DefineConstant("Array", object.Map{})
	compiler.DefineConstant("String", object.Map{})
	for _, fun := range globals {
		compiler.DefineConstant(fun.(vm.NativeFunction).Name, fun)
	}
	// put object.Map{} in globals as "exports" as index 0
	// so user can do exports["key"] = "value"
	// and access it from other files by import
	globals = append([]object.Object{*object.PrototypeString}, globals...)
	globals = append([]object.Object{*object.PrototypeArray}, globals...)
	globals = append([]object.Object{object.Map{Map: map[string]object.Object{}}}, globals...)

	function, constants := compiler.Compile(statements)
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Println(function, constants)
		parser.Decompile(function)
	}

	if debug {
		for i, c := range constants {
			if fn, ok := c.(object.Function); ok {
				fmt.Println("Function at constant index:", i)
				parser.Decompile(fn)
			}
		}
	}

	vm := vm.NewVM(function, constants, globals)
	if debug {
		fmt.Println("----------Output----------")
	}
	vm.Run()
	if debug {
		fmt.Println("--------------------------")
	}
	m := vm.Globals[0].(object.Map)
	return &m
}

func main() {
	filename := "/home/abbas/repos/interpreter-v2/exports.test.lox"
	runFile(filename, false)
	fmt.Println(colors.Colorize("Program finished successfully!", colors.GREEN))
}
