package parser

import (
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

func (p *Parser) Parse() []Statement {
	statements := []Statement{}
	for p.index < len(p.tokens) {
		var statement Statement

		if p.consumeIfExists(token.LET) {
			name := p.currentToken
			p.move()
			var exp Expression = &LiteralExpression{}
			if p.consumeIfExists(token.ASSIGN) {
				exp = p.Expression()
			}
			statement = &DeclarationStatement{
				name: name, expression: exp,
			}
			p.consume(token.SEMICOLON, "Expected a semicolon")
		} else if p.consumeIfExists(token.PRINT) {
			statement = p.printStatement()
		} else {
			statement = ExpressionStatement{expression: p.Expression()}
			p.consume(token.SEMICOLON, "Expected a semicolon")
		}

		statements = append(statements, statement)
	}
	return statements
}

func (p *Parser) printStatement() Statement {
	exp := p.Expression()
	p.consume(token.SEMICOLON, "Expected a semicolon")
	return PrintStatement{
		expression: exp,
	}
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

func (p *Parser) consume(t token.TokenType, err string) {
	if p.currentToken.Type == t {
		p.move()
	} else {
		panic(err)
	}
}

func (p *Parser) Expression() Expression {
	exp := p.BinaryExpression()

	if assignee, ok := exp.(*IdentifierExpression); ok && p.consumeIfExists(token.ASSIGN) {
		exp = &AssignmentExpression{
			assignee:   *assignee,
			assignment: p.Expression(),
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

	if p.currentToken.Type != token.NUMBER {
		// _ := fmt.Errorf("Unknown token: %v", p.currentToken)
		panic("Error")
	}

	exp := &LiteralExpression{
		token: p.currentToken,
	}
	p.move()
	return exp
}
