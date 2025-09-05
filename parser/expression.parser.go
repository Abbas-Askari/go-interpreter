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

	if assignee, ok := exp.(*PropertyExpression); ok && p.consumeIfExists(token.ASSIGN) {
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
	hasOperand := p.consumeIfExists(token.SLASH, token.MULTIPLY, token.PERCENT)
	for hasOperand {
		right := p.Unary()
		left = &BinaryExpression{
			left:    left,
			operand: operand,
			right:   right,
		}
		operand = p.currentToken
		hasOperand = p.consumeIfExists(token.SLASH, token.MULTIPLY, token.PERCENT)
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
		var exp Expression = &IdentifierExpression{
			token: p.currentToken,
		}
		p.move()
		for p.currentToken.Type == token.LPAREN || p.currentToken.Type == token.DOT {
			if p.consumeIfExists(token.LPAREN) {
				// Function Call
				args := []Expression{}
				if !p.consumeIfExists(token.RPAREN) {
					for {
						args = append(args, p.Expression())
						if p.consumeIfExists(token.RPAREN) {
							break
						}
						p.consume(token.COMMA, "Expected ',' between function call arguments")
					}
				}
				exp = &CallExpression{
					callee:    exp,
					arguments: args,
				}
			}
			if p.consumeIfExists(token.DOT) {
				if p.currentToken.Type != token.IDENTIFIER {
					panic("Expected property name after '.'")
				}
				property := p.currentToken
				p.move()
				exp = &PropertyExpression{
					object:   exp,
					property: property.Literal,
				}
			}
		}

		return exp
	}

	if p.match(token.NUMBER, token.STRING, token.TRUE, token.FALSE, token.NIL) {
		exp := &LiteralExpression{
			token: p.currentToken,
		}
		p.move()
		return exp
	}

	if p.consumeIfExists(token.LBRACE) {
		exp := &MapExpression{
			pairs: map[Expression]Expression{},
		}
		if !p.consumeIfExists(token.RBRACE) {
			for {
				key := p.Expression()
				if _, ok := key.(*IdentifierExpression); !ok {
					if _, ok := key.(*LiteralExpression); !ok {
						panic("Only identifiers and literals can be map keys")
					}
				}
				p.consume(token.COLON, "Expected ':' between key and value in map")
				value := p.Expression()
				exp.pairs[key] = value
				if p.consumeIfExists(token.RBRACE) {
					break
				} else {
					p.consume(token.COMMA, "Expected ',' between map pairs")
				}
			}
		}
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
