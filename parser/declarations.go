package parser

import (
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

type DeclarationType int

const (
	VariableDeclarationType DeclarationType = iota
	StatementDeclarationType
	FunctionDeclarationType
	ImportDeclarationType
)

type Declaration interface {
	Emit(p interfaces.ICompiler)
	String() string
	Type() DeclarationType
}

type VariableDeclaration struct {
	name       token.Token
	expression Expression
	// globalIndex int
}

func (d *VariableDeclaration) Emit(c interfaces.ICompiler) {
	d.expression.Emit(c)
	c.Declare(d.name.Literal)
}

func (dx VariableDeclaration) String() string {
	return fmt.Sprintf("Declaration: %v = %v\n", dx.name, dx.expression)
}

func (d *VariableDeclaration) Type() DeclarationType {
	return VariableDeclarationType
}

type FunctionDeclaration struct {
	name       token.Token
	body       BlockStatement
	parameters []IdentifierExpression
	// globalIndex int
}

func (d *FunctionDeclaration) Emit(c interfaces.ICompiler) {
	c.Emit(op.OpClosure, d.name.Line, d.name.Column) // placeholder
	c.Emit(op.OpCode(0), d.name.Line, d.name.Column)
	indexIndex := c.GetBytecodeLength() - 1
	c.Declare(d.name.Literal)
	c.EnterTarget(d.name.Literal)
	c.EnterScope()
	for _, param := range d.parameters {
		c.Declare(param.token.Literal)
	}
	d.body.Emit(c)
	c.ExitScope()
	c.Emit(op.OpNil, d.name.Line, d.name.Column)
	c.Emit(op.OpReturn, d.name.Line, d.name.Column)
	// d.expression.Emit(c)
	index := c.ExitTarget(len(d.parameters))
	c.SetOpCode(indexIndex, op.OpCode(index))
}

func (dx FunctionDeclaration) String() string {
	return fmt.Sprintf("Function: %v = %v\n", dx.parameters, dx.body)
}

func (d *FunctionDeclaration) Type() DeclarationType {
	return FunctionDeclarationType
}

type ImportDeclaration struct {
	Module  token.Token
	Exports *object.Map
}

func (d *ImportDeclaration) Emit(c interfaces.ICompiler) {
	index := c.AddConstant(*d.Exports)
	c.Emit(op.OpConstant, d.Module.Line, d.Module.Column)
	c.Emit(op.OpCode(index), d.Module.Line, d.Module.Column)
	c.Declare(d.Module.Literal)
}

func (dx ImportDeclaration) String() string {
	return fmt.Sprintf("Import: %v\n", dx.Module)
}

func (d *ImportDeclaration) Type() DeclarationType {
	return ImportDeclarationType
}
