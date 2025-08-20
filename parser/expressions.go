package parser

import (
	"Abbas-Askari/interpreter-v2/colors"
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"strconv"
)

type ExpressionType string

const (
	BINARY_EXPRESSION     = "BINARY_EXPRESSION"
	LITERAL_EXPRESSION    = "LITERAL_EXPRESSION"
	IDENTIFIER_EXPRESSION = "IDENTIFIER_EXPRESSION"
	ASSIGNMENT_EXPRESSION = "ASSIGNMENT_EXPRESSION"
)

type Expression interface {
	// GetType() ExpressionType
	Emit(p interfaces.ICompiler)
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

func (b *BinaryExpression) Emit(c interfaces.ICompiler) {
	b.left.Emit(c)
	b.right.Emit(c)
	mapping := map[token.TokenType]op.OpCode{
		token.PLUS:     op.OpAdd,
		token.MINUS:    op.OpSub,
		token.SLASH:    op.OpDiv,
		token.MULTIPLY: op.OpMul,
	}
	c.Emit(mapping[b.operand.Type])
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

func (l *LiteralExpression) Emit(c interfaces.ICompiler) {
	c.Emit(op.OpConstant)
	value, err := strconv.ParseFloat(l.token.Literal, 64)
	if err != nil {
		panic("This shouldn't have happened! Lexer is probably broken")
	}
	index := c.AddConstant(object.Number{
		Value: value,
	})
	c.Emit(op.OpCode(index))
}

type IdentifierExpression struct {
	token token.Token
}

func (l *IdentifierExpression) GetType() ExpressionType {
	return IDENTIFIER_EXPRESSION
}

func (l *IdentifierExpression) String() string {
	return fmt.Sprintf("%v(%v)", colors.Colorize(IDENTIFIER_EXPRESSION, colors.BLUE), l.token)
}

func (l *IdentifierExpression) Emit(c interfaces.ICompiler) {
	c.GetGlobal(l.token)
}

type AssignmentExpression struct {
	assignee   IdentifierExpression
	assignment Expression
}

// func (l *AssignmentExpression) GetType() ExpressionType {
// 	return IDENTIFIER_EXPRESSION
// }

func (l *AssignmentExpression) String() string {
	return fmt.Sprintf("%v(%v = %v)", colors.Colorize(ASSIGNMENT_EXPRESSION, colors.BLUE), l.assignee, l.assignment)
}

func (l *AssignmentExpression) Emit(c interfaces.ICompiler) {
	l.assignment.Emit(c)
	c.SetGlobal(l.assignee.token)
}
