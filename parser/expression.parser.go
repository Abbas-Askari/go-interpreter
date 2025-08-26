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
	tok := p.currentToken
	if p.consumeIfExists(token.OR) {
		exp = &BinaryExpression{
			left:    exp,
			operand: token.Token{Type: tok.Type},
			right:   p.And(),
		}
	}
	return exp
}

func (p *Parser) And() Expression {
	exp := p.Equality()
	tok := p.currentToken
	if p.consumeIfExists(token.AND) {
		exp = &BinaryExpression{
			left:    exp,
			operand: token.Token{Type: tok.Type},
			right:   p.Equality(),
		}
	}
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
	exp := p.Term()
	tok := p.currentToken
	if p.consumeIfExists(token.LESS_EQUAL, token.GREATER_EQUAL, token.LESS, token.GREATER) {
		exp = &BinaryExpression{
			left:    exp,
			operand: token.Token{Type: tok.Type},
			right:   p.Term(),
		}
	}
	return exp
}

func (p *Parser) Term() Expression {
	left := p.Factor()

	operand := p.currentToken
	hasOperand := p.consumeIfExists(token.PLUS) || p.consumeIfExists(token.MINUS)
	for hasOperand {
		right := p.Factor()
		left = &BinaryExpression{
			left:    left,
			operand: operand,
			right:   right,
		}
		operand = p.currentToken
		hasOperand = p.consumeIfExists(token.PLUS) || p.consumeIfExists(token.MINUS)
	}

	return left
}

func (p *Parser) Factor() Expression {
	left := p.Unary()

	operand := p.currentToken
	hasOperand := p.consumeIfExists(token.SLASH) || p.consumeIfExists(token.MULTIPLY)
	for hasOperand {
		right := p.Unary()
		left = &BinaryExpression{
			left:    left,
			operand: operand,
			right:   right,
		}
		operand = p.currentToken
		hasOperand = p.consumeIfExists(token.SLASH) || p.consumeIfExists(token.MULTIPLY)
	}

	return left
}

func (p *Parser) Unary() Expression {
	tok := p.currentToken
	if p.consumeIfExists(token.NOT, token.MINUS, token.PLUS) {
		return &UnaryExpression{
			exp:     p.Unary(),
			operand: tok,
		}
	}

	return p.LiteralExpression()
}

func (p *Parser) LiteralExpression() Expression {
	if p.currentToken.Type == token.IDENTIFIER {
		exp := &IdentifierExpression{
			token: p.currentToken,
		}
		p.move()
		return exp
	}

	if p.match(token.NUMBER, token.STRING, token.TRUE, token.FALSE) {
		exp := &LiteralExpression{
			token: p.currentToken,
		}
		p.move()
		return exp
	}

	if p.consumeIfExists(token.LPAREN) {
		// Not introducing a Grouping struct, Will just return a normal expression because of laziness
		exp := p.Expression()
		p.consume(token.RPAREN, "Unclosed '(', Expected ')'")
		return exp
	}

	panic(fmt.Errorf("Unexpected token: %v", p.currentToken))
}
