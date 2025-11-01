package main

import (
	"Abbas-Askari/interpreter-v2/internal/compiler"
	"Abbas-Askari/interpreter-v2/internal/lexer"
	"Abbas-Askari/interpreter-v2/internal/object"
	"Abbas-Askari/interpreter-v2/internal/parser"
	"Abbas-Askari/interpreter-v2/internal/vm"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

// TODO: Clean this ugly main file

func runFile(filename string, debug bool, safemode bool) *object.Map {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("File not found!\nUsage: run [filepath]"))
	}

	tokens := lexer.Tokenize(string(fileContent))

	if debug {
		fmt.Println(tokens)
	}

	p := parser.NewParser(tokens, filename)
	statements := p.Parse()
	if debug {
		fmt.Println("------------AST------------")
		fmt.Println(statements)
		fmt.Println("---------------------------")
	}

	lib := vm.GetLibraryMaps()
	unsafeLibs := []string{
		"os",
		"fs",
		"net",
		"http",
		"tcp",
	}
	for _, stmt := range statements {
		if imp, ok := stmt.(*parser.ImportDeclaration); ok {
			if debug {
				fmt.Println("Importing module:", imp.Module.Literal)
			}
			module := lib[imp.Module.Literal]
			if module != nil {
				if safemode && slices.Contains(unsafeLibs, imp.Module.Literal) {
					panic(fmt.Sprintf("Cannot import unsafe module '%s' in safe mode", imp.Module.Literal))
				}
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
			imp.Exports = runFile(modulePath, debug, safemode)
		}
	}

	compiler := compiler.NewCompiler(filename)

	globals := vm.GetNativeFunctions()

	// TODO: Fix this hacky way of adding prototypes and Map
	compiler.DefineConstant("exports", object.Map{})
	compiler.DefineConstant("Array", object.Map{})
	compiler.DefineConstant("String", object.Map{})
	compiler.DefineConstant("Buffer", object.Map{})
	for _, fun := range globals {
		compiler.DefineConstant(fun.(vm.NativeFunction).Name, fun)
	}
	globals = append([]object.Object{*object.PrototypeBuffer}, globals...)
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

	vm := vm.NewVM(function, constants, globals, safemode)
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
	verbose := flag.Bool("v", false, "Enable verbose mode")
	safemode := flag.Bool("s", false, "Enable safe mode")
	flag.Parse()

	// After parsing, flag.Args() gives the remaining non-flag arguments
	args := flag.Args()

	if len(args) < 1 {
		fmt.Printf("Usage: %s [filename] [-v]\n", os.Args[0])
		os.Exit(1)
	}

	filename := args[0]
	if *verbose {
		fmt.Println("Running in verbose mode")
	}
	runFile(filename, *verbose, *safemode)
	// fmt.Println(colors.Colorize("Program finished successfully!", colors.GREEN))
}
