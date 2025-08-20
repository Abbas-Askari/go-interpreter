package compiler

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"slices"
)

type Compiler struct {
	stream    []op.OpCode
	constants []object.Object
	globals   []string
}

func NewCompiler() *Compiler {
	return &Compiler{
		stream:    []op.OpCode{},
		constants: []object.Object{},
		globals:   []string{},
	}
}

func (c *Compiler) AddConstant(o object.Object) int {
	c.constants = append(c.constants, o)
	return len(c.constants) - 1
}

func (c *Compiler) AddGlobal(name string) int {
	c.globals = append(c.globals, name)
	return len(c.globals) - 1
}

func (c *Compiler) GetGlobal(name token.Token) {
	// c.globals
	i := slices.Index(c.globals, name.Literal)
	if i == -1 {
		panic(fmt.Errorf("Undeclared identifier: %v", name))
	}
	c.Emit(op.OpLoadGlobal)
	c.Emit(op.OpCode(i))
	// return index
}

func (c *Compiler) SetGlobal(name token.Token) {
	i := slices.Index(c.globals, name.Literal)
	if i == -1 {
		panic(fmt.Errorf("Undeclared identifier: %v", name))
	}
	c.Emit(op.OpSetGlobal)
	c.Emit(op.OpCode(i))
}

func (c *Compiler) Emit(op op.OpCode) {
	c.stream = append(c.stream, op)
}

func (c *Compiler) Compile(statements []parser.Statement) ([]op.OpCode, []object.Object) {
	for _, statement := range statements {
		statement.Emit(c)
	}

	return c.stream, c.constants
}
