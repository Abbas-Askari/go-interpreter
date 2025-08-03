package parser

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
)

type Parser struct {
	tokens       []token.Token
	currentToken token.Token
	index        int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:       tokens,
		currentToken: tokens[0],
		index:        0,
	}
}

func (p *Parser) Parse() Expression {
	return p.Expression()
}

func (p *Parser) move() {
	p.index++
	if len(p.tokens) > p.index {
		p.currentToken = p.tokens[p.index]
	}
}

func (p *Parser) consumeIfExists(t token.TokenType) bool {
	if p.currentToken.Type == t {
		p.move()
		return true
	}
	return false
}

func (p *Parser) Expression() Expression {
	return p.BinaryExpression()
}

func (p *Parser) BinaryExpression() Expression {
	left := p.LiteralExpression()

	operand := p.currentToken
	hasOperand := p.consumeIfExists(token.PLUS) || p.consumeIfExists(token.MINUS) || p.consumeIfExists(token.SLASH) || p.consumeIfExists(token.MULTIPLY)
	for hasOperand {
		right := p.LiteralExpression()
		left = &BinaryExpression{
			left:    left,
			operand: operand,
			right:   right,
		}
		operand = p.currentToken
		hasOperand = p.consumeIfExists(token.PLUS) || p.consumeIfExists(token.MINUS) || p.consumeIfExists(token.SLASH) || p.consumeIfExists(token.MULTIPLY)
	}

	return left
}

func (p *Parser) LiteralExpression() Expression {
	if p.currentToken.Type != token.NUMBER {
		panic("Expected to be a number")
	}

	exp := &LiteralExpression{
		token: p.currentToken,
	}
	p.move()
	return exp
}

func Emit(ast Expression) ([]op.OpCode, []object.Object) {
	stream := []op.OpCode{}
	constants := []object.Object{}

	ast.Emit(
		func(oc op.OpCode) {
			stream = append(stream, oc)
		},
		func(o object.Object) int {
			constants = append(constants, o)
			return len(constants) - 1
		},
	)

	return stream, constants
}
