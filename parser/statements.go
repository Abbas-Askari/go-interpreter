package parser

import (
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/op"
	"go/token"
)

type Statement interface {
	Emit(p interfaces.ICompiler)
}

type PrintStatement struct {
	expression Expression
}

func (ps PrintStatement) Emit(c interfaces.ICompiler) {
	ps.expression.Emit(c)
	c.Emit(op.OpPrint)
}

type ExpressionStatement struct {
	expression Expression
}

func (e ExpressionStatement) Emit(c interfaces.ICompiler) {
	e.expression.Emit(c)
	c.Emit(op.OpPop)
}

type DeclarationStatement struct {
	name        token.Token
	expression  Expression
	globalIndex int
}

func (d *DeclarationStatement) Emit(c interfaces.ICompiler) {
	d.expression.Emit(c)
	c.Emit(op.OpPop)
}
