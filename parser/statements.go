package parser

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
)

type Statement interface {
	Emit(func(op.OpCode), func(object.Object) int)
}

type PrintStatement struct {
	expression Expression
}

func (p PrintStatement) Emit(emit func(op.OpCode), addConst func(object.Object) int) {
	p.expression.Emit(emit, addConst)
	emit(op.OpPrint)
}

type ExpressionStatement struct {
	expression Expression
}

func (e ExpressionStatement) Emit(emit func(op.OpCode), addConst func(object.Object) int) {
	e.expression.Emit(emit, addConst)
	emit(op.OpPop)
}

type DeclarationStatement struct {
	expression Expression
}
