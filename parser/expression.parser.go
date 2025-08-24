package parser

import (
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

func (p *Parser) Expression() Expression {
	exp := p.Or()

	if assignee, ok := exp.(*IdentifierExpression); ok && p.consumeIfExists(token.ASSIGN) {
		exp = &AssignmentExpression{
			assignee:   *assignee,
			assignment: p.Expression(),
		}
	}
	return exp
}

func (p *Parser) Or() Expression {
	exp := p.And()
	// if p.consumeIfExists(token.)
	return exp
}

func (p *Parser) And() Expression {
	exp := p.Equality()
	return exp
}

func (p *Parser) Equality() Expression {
	exp := p.Comparison()
	tok := p.currentToken
	if p.consumeIfExists(token.EQUAL_EQUAL, token.NOT_EQUAL) {
		exp = &BinaryExpression{
			left:    exp,
			operand: token.Token{Type: tok.Type},
			right:   p.Comparison(),
		}
	}
	return exp
}

func (p *Parser) Comparison() Expression {
	exp := p.BinaryExpression()
	tok := p.currentToken
	if p.consumeIfExists(token.LESS_EQUAL, token.GREATER_EQUAL, token.LESS, token.GREATER) {
		exp = &BinaryExpression{
			left:    exp,
			operand: token.Token{Type: tok.Type},
			right:   p.BinaryExpression(),
		}
	}
	return exp
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
	if p.currentToken.Type == token.IDENTIFIER {
		exp := &IdentifierExpression{
			token: p.currentToken,
		}
		p.move()
		return exp
	}

	if p.currentToken.Type == token.NUMBER || p.currentToken.Type == token.STRING {
		exp := &LiteralExpression{
			token: p.currentToken,
		}
		p.move()
		return exp
	}

	panic(fmt.Errorf("Unexpected token: %v", p.currentToken))
}
