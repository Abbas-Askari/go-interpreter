package parser

import (
	"Abbas-Askari/interpreter-v2/colors"
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"strconv"
)

type ExpressionType string

const (
	BINARY_EXPRESSION  = "BINARY_EXPRESSION"
	LITERAL_EXPRESSION = "LITERAL_EXPRESSION"
)

type Expression interface {
	GetType() ExpressionType
	Emit(func(op.OpCode), func(object.Object) int)
}

type BinaryExpression struct {
	left    Expression
	operand token.Token
	right   Expression
}

func (b *BinaryExpression) GetType() ExpressionType {
	return BINARY_EXPRESSION
}

func (b *BinaryExpression) String() string {
	return fmt.Sprintf("%v(%v %v %v)", colors.Colorize(BINARY_EXPRESSION, colors.BLUE), b.left, b.operand, b.right)
}

func (b *BinaryExpression) Emit(emit func(op.OpCode), addConst func(object.Object) int) {
	b.left.Emit(emit, addConst)
	b.right.Emit(emit, addConst)
	mapping := map[token.TokenType]op.OpCode{
		token.PLUS:     op.OpAdd,
		token.MINUS:    op.OpSub,
		token.SLASH:    op.OpDiv,
		token.MULTIPLY: op.OpMul,
	}
	emit(mapping[b.operand.Type])
}

type LiteralExpression struct {
	token token.Token
}

func (l *LiteralExpression) GetType() ExpressionType {
	return LITERAL_EXPRESSION
}

func (l *LiteralExpression) String() string {
	return fmt.Sprintf("%v(%v)", colors.Colorize(LITERAL_EXPRESSION, colors.BLUE), l.token)
}

func (l *LiteralExpression) Emit(emit func(op.OpCode), addConstant func(object.Object) int) {
	emit(op.OpConstant)
	value, err := strconv.ParseFloat(l.token.Literal, 64)
	if err != nil {
		panic("This shouldn't have happened! Lexer is probably broken")
	}
	index := addConstant(object.Number{
		Value: value,
	})
	emit(op.OpCode(index))
}
