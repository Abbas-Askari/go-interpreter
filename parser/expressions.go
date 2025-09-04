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
	UNARY_EXPRESSION      = "UNARY_EXPRESSION"
	CALL_EXPRESSION       = "CALL_EXPRESSION"
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
		token.PLUS:          op.OpAdd,
		token.MINUS:         op.OpSub,
		token.SLASH:         op.OpDiv,
		token.MULTIPLY:      op.OpMul,
		token.PERCENT:       op.OpMod,
		token.EQUAL_EQUAL:   op.OpEqual,
		token.NOT_EQUAL:     op.OpNotEqual,
		token.GREATER:       op.OpGreaterThan,
		token.LESS:          op.OpLessThan,
		token.GREATER_EQUAL: op.OpGreaterEqual,
		token.LESS_EQUAL:    op.OpLessEqual,
		token.AND:           op.OpAnd,
		token.OR:            op.OpOr,
	}
	c.Emit(mapping[b.operand.Type])
}

type UnaryExpression struct {
	operand token.Token
	exp     Expression
}

func (b *UnaryExpression) GetType() ExpressionType {
	return UNARY_EXPRESSION
}

func (b *UnaryExpression) String() string {
	return fmt.Sprintf("%v(%v %v)", colors.Colorize(UNARY_EXPRESSION, colors.BLUE), b.operand, b.exp)
}

func (b *UnaryExpression) Emit(c interfaces.ICompiler) {
	b.exp.Emit(c)
	mapping := map[token.TokenType]op.OpCode{
		token.NOT:   op.OpNot,
		token.MINUS: op.OpNeg,
		// token.PLUS:  op.OpConvNum,
	}
	op, ok := mapping[b.operand.Type]

	if !ok {
		// No need to eval plus right on numbers.
		// Will deal with Unary Plus on string later.
		return
	}

	c.Emit(op)
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
	if l.token.Type == token.TRUE {
		c.Emit(op.OpTrue)
		return
	} else if l.token.Type == token.FALSE {
		c.Emit(op.OpFalse)
		return
	}

	c.Emit(op.OpConstant)
	var index int

	if l.token.Type == token.NUMBER {
		value, err := strconv.ParseFloat(l.token.Literal, 64)
		if err != nil {
			panic("This shouldn't have happened! Lexer is probably broken")
		}
		index = c.AddConstant(object.Number{
			Value: value,
		})
	} else if l.token.Type == token.STRING {
		index = c.AddConstant(object.String{
			Value: l.token.Literal,
		})
	} else if l.token.Type == token.NIL {
		index = c.AddConstant(object.Nil{})
	} else {
		panic("This shouldn't have happened! Lexer is probably broken")
	}

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
	c.GetIdentifier(l.token)
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

type CallExpression struct {
	callee    Expression
	arguments []Expression
}

// func (l *CallExpression) GetType() ExpressionType {
// 	return IDENTIFIER_EXPRESSION
// }

func (l *CallExpression) String() string {
	return fmt.Sprintf("%v(%v = %v)", colors.Colorize(CALL_EXPRESSION, colors.BLUE), l.callee, l.arguments)
}

func (l *CallExpression) Emit(c interfaces.ICompiler) {
	l.callee.Emit(c)
	for _, arg := range l.arguments {
		arg.Emit(c)
	}
	c.Emit(op.OpCall)
	c.Emit(op.OpCode(len(l.arguments)))
}
