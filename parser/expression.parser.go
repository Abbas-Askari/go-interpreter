package parser

import (
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

func (p *Parser) Expression() Expression {
	exp := p.Or()

	assignmentToken := p.currentToken
	if assignee, ok := exp.(*IdentifierExpression); ok && p.consumeIfExists(token.ASSIGN) {
		exp = &AssignmentExpression{
			assignee:        *assignee,
			assignment:      p.Expression(),
			assignmentToken: assignmentToken,
		}
	}

	if assignee, ok := exp.(*PropertyExpression); ok && p.consumeIfExists(token.ASSIGN) {
		exp = &AssignmentExpression{
			assignee:        *assignee,
			assignment:      p.Expression(),
			assignmentToken: assignmentToken,
		}
	}

	if assignee, ok := exp.(*IndexExpression); ok && p.consumeIfExists(token.ASSIGN) {
		exp = &AssignmentExpression{
			assignee:        *assignee,
			assignment:      p.Expression(),
			assignmentToken: assignmentToken,
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

	return p.LiteralExpression(false)
}

func (p *Parser) LiteralExpression(isConstructorCall bool) Expression {
	if p.currentToken.Type == token.IDENTIFIER {
		var exp Expression = &IdentifierExpression{
			token: p.currentToken,
		}
		p.move()
		for p.currentToken.Type == token.LPAREN || p.currentToken.Type == token.DOT || p.currentToken.Type == token.LBRACKET {
			// Function Call
			if p.consumeIfExists(token.LPAREN) {
				args := []Expression{}
				parenToken := p.tokens[p.index-1]
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
					paren:     parenToken,
				}
			}
			// Property Access
			if p.consumeIfExists(token.DOT) {
				dotToken := p.tokens[p.index-1]
				if p.currentToken.Type != token.IDENTIFIER {
					panic("Expected property name after '.'")
				}
				property := p.currentToken
				p.move()
				exp = &PropertyExpression{
					object:   exp,
					property: property.Literal,
					dotToken: dotToken,
				}
			}
			// Property Access with Index
			if p.consumeIfExists(token.LBRACKET) {
				indexToken := p.tokens[p.index-1]
				index := p.Expression()
				p.consume(token.RBRACKET, "Expected ']' after index")
				exp = &IndexExpression{
					object:     exp,
					index:      index,
					indexToken: indexToken,
				}
			}
		}

		if isConstructorCall {
			call, ok := exp.(*CallExpression)
			if !ok {
				panic("Expected function call after 'new'")
			}
			call.isConstructorCall = true
			return call
		}

		return exp
	}

	if isConstructorCall {
		panic("Expected identifier after 'new'")
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

	if p.consumeIfExists(token.LBRACKET) {
		exp := &ArrayExpression{
			elements: []Expression{},
		}
		if !p.consumeIfExists(token.RBRACKET) {
			for {
				exp.elements = append(exp.elements, p.Expression())
				if p.consumeIfExists(token.RBRACKET) {
					break
				}
				p.consume(token.COMMA, "Expected ',' between array elements")
			}
		}
		return exp
	}

	if p.consumeIfExists(token.LPAREN) {
		// Not introducing a Grouping struct, Will just return a normal expression because of laziness
		parameters := []IdentifierExpression{}
		if !p.consumeIfExists(token.RPAREN) {
			exp := p.Expression()
			if p.currentToken.Type == token.RPAREN && p.tokens[p.index+1].Type != token.ARROW {
				p.consume(token.RPAREN, "Expected ')' after argument list")
				return exp
			}
			if ident, ok := exp.(*IdentifierExpression); ok {
				parameters = append(parameters, *ident)
			} else {
				panic("Expected parameter name for arrow function")
			}
			for p.consumeIfExists(token.COMMA) {
				param, ok := p.Expression().(*IdentifierExpression)
				if !ok {
					panic("Expected parameter name for arrow function")
				}
				parameters = append(parameters, *param)
			}
			p.consume(token.RPAREN, "Expected ')' after argument list")
		}
		p.consume(token.ARROW, "Expected '=>' after parameter list")

		var body BlockStatement
		if p.consumeIfExists(token.LBRACE) {
			body = p.blockStatement()
		} else {
			body = BlockStatement{
				declarations: []Declaration{&ReturnStatement{
					exp: p.Expression(),
				}},
			}
		}
		return &FunctionDeclaration{
			name:       token.Token{Type: token.IDENTIFIER, Literal: "<anonymous>"},
			body:       body,
			parameters: parameters,
		}
	}

	if p.consumeIfExists(token.NEW) {
		return p.LiteralExpression(true)
	}
	panic(fmt.Errorf("Unexpected token: %v,\ntokens: %v,\nindex: %d\nlength: %d", p.currentToken, p.tokens, p.index, len(p.tokens)))
}
