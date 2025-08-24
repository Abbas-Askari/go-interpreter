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
		statements = append(statements, p.Statement())
	}
	return statements
}

func (p *Parser) Statement() Statement {
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
	} else if p.consumeIfExists(token.IF) {
		exp := p.Expression()
		thenStatement := p.Statement()
		ifStatement := IfStatement{
			condition:     exp,
			thenStatement: thenStatement,
		}

		if p.consumeIfExists(token.ELSE) {
			elseStatement := p.Statement()
			ifStatement.elseStatement = &elseStatement
		}

		statement = &ifStatement

	} else if p.consumeIfExists(token.LBRACE) {
		statements := []Statement{}
		for !p.consumeIfExists(token.RBRACE) {
			statements = append(statements, p.Statement())
		}
		// if p.index == len(p.tokens) &&
		statement = &BlockStatement{
			statements: statements,
		}
	} else if p.consumeIfExists(token.PRINT) {
		statement = p.printStatement()
	} else {
		statement = ExpressionStatement{expression: p.Expression()}
		p.consume(token.SEMICOLON, "Expected a semicolon")
	}

	return statement
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

func (p *Parser) consumeIfExists(types ...token.TokenType) bool {
	for _, t := range types {
		if p.currentToken.Type == t {
			p.move()
			return true
		}
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
