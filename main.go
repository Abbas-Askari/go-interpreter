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
	"path/filepath"
)

// TODO: Clean this ugly main file

func runFile(filename string, debug bool) *object.Map {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("File not found!\nUsage: run [filepath]"))
	}

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

	lib := vm.GetLibraryMaps()
	for _, stmt := range statements {
		if imp, ok := stmt.(*parser.ImportDeclaration); ok {
			if debug {
				fmt.Println("Importing module:", imp.Module.Literal)
			}
			module := lib[imp.Module.Literal]
			if module != nil {
				imp.Exports = module
				continue
			}
			dir := filepath.Dir(filename)
			entries, err := os.ReadDir(dir)
			if err != nil {
				panic(err)
			}
			target := imp.Module.Literal + ".turtle"
			modulePath := ""
			for _, e := range entries {
				if !e.IsDir() && e.Name() == target {
					modulePath = filepath.Join(dir, e.Name())
					break
				}
			}
			if err != nil {
				panic(err)
			}
			if modulePath == "" {
				panic(fmt.Sprintf("Module not found: %s", modulePath))
			}
			imp.Exports = runFile(modulePath, debug)
		}
	}

	compiler := compiler.NewCompiler(filename)

	globals := vm.GetNativeFunctions()

	// TODO: Fix this hacky way of adding prototypes and Map
	compiler.DefineConstant("exports", object.Map{})
	compiler.DefineConstant("Array", object.Map{})
	compiler.DefineConstant("String", object.Map{})
	for _, fun := range globals {
		compiler.DefineConstant(fun.(vm.NativeFunction).Name, fun)
	}
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
	filename := "/home/abbas/repos/interpreter-v2/scripts/tcp-server.test.turtle"
	runFile(filename, false)
	fmt.Println(colors.Colorize("Program finished successfully!", colors.GREEN))
}
