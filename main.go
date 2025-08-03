package main

import (
	"Abbas-Askari/interpreter-v2/lexer"
	"Abbas-Askari/interpreter-v2/parser"
	"fmt"
	"os"
)

func main() {
	filename := "./test.lox"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("File not found!\nUsage: run [filepath]")
	}
	fmt.Println(string(fileContent))

	tokens := lexer.Tokenize(string(fileContent))

	fmt.Println(tokens)

	p := parser.NewParser(tokens)
	ast := p.Parse()
	fmt.Println(ast)

	stream, constants := parser.Emit(ast)
	fmt.Println(parser.Emit(ast))

	parser.Decompile(stream, constants)
}
