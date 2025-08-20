package parser

import (
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

type Statement interface {
	Emit(p interfaces.ICompiler)
	String() string
}

type PrintStatement struct {
	expression Expression
}

func (ps PrintStatement) Emit(c interfaces.ICompiler) {
	ps.expression.Emit(c)
	c.Emit(op.OpPrint)
}

func (ps PrintStatement) String() string {
	return fmt.Sprintf("PRINT: %v\n", ps.expression)
}

type ExpressionStatement struct {
	expression Expression
}

func (e ExpressionStatement) Emit(c interfaces.ICompiler) {
	e.expression.Emit(c)
	c.Emit(op.OpPop)
}

func (ex ExpressionStatement) String() string {
	return fmt.Sprintf("Expression: %v\n", ex.expression)
}

type DeclarationStatement struct {
	name       token.Token
	expression Expression
	// globalIndex int
}

func (d *DeclarationStatement) Emit(c interfaces.ICompiler) {
	d.expression.Emit(c)
	c.Emit(op.OpSetGlobal)
	index := c.AddGlobal(d.name.Literal)
	c.Emit(op.OpCode(index))
	c.Emit(op.OpPop)
}

func (dx DeclarationStatement) String() string {
	return fmt.Sprintf("Declaration: %v - %v\n", dx.name, dx.expression)
}
