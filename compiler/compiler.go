package compiler

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/parser"
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

func (c *Compiler) Emit(op op.OpCode) {
	c.stream = append(c.stream, op)
}

func (c *Compiler) Compile(statements []parser.Statement) ([]op.OpCode, []object.Object) {
	for _, statement := range statements {
		statement.Emit(c)
	}

	return c.stream, c.constants
}
