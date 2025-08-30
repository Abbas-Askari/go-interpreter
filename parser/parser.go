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

func (p *Parser) Parse() []Declaration {
	statements := []Declaration{}
	for p.index < len(p.tokens) {
		statements = append(statements, p.Declaration())
	}
	return statements
}

func (p *Parser) Declaration() Declaration {
	var statement Declaration

	if p.consumeIfExists(token.LET) {
		name := p.currentToken
		p.move()
		var exp Expression = &LiteralExpression{}
		if p.consumeIfExists(token.ASSIGN) {
			exp = p.Expression()
		}
		statement = &VariableDeclaration{
			name: name, expression: exp,
		}
		p.consume(token.SEMICOLON, "Expected a semicolon")
	} else if p.consumeIfExists(token.FUNCTION) {
		name := p.currentToken
		p.move()
		p.consume(token.LPAREN, "Expected '(' after function name")
		parameters := []IdentifierExpression{}
		if !p.consumeIfExists(token.RPAREN) {
			param, ok := p.Expression().(*IdentifierExpression)
			if !ok {
				panic("Expected parameter name")
			}
			parameters = append(parameters, *param)
			for p.consumeIfExists(token.COMMA) {
				param, ok := p.Expression().(*IdentifierExpression)
				if !ok {
					panic("Expected parameter name")
				}
				parameters = append(parameters, *param)
			}
		}
		p.consume(token.RPAREN, "Expected ')' after argument list")
		p.consume(token.LBRACE, "Expected '{' after argument list")
		declarations := []Declaration{}
		for !p.consumeIfExists(token.RBRACE) {
			declarations = append(declarations, p.Declaration())
		}
		// if p.index == len(p.tokens) &&
		block := BlockStatement{
			declarations: declarations,
		}
		statement = &FunctionDeclaration{
			name: name, body: block,
			parameters: parameters,
		}
	} else {
		statement = p.Statement()
	}

	return statement
}

func (p *Parser) Statement() Statement {
	var statement Statement

	if p.consumeIfExists(token.IF) {
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

	} else if p.consumeIfExists(token.FOR) {
		var init Declaration = nil
		var cond Expression = nil
		var adv Expression = nil
		if p.match(token.LET) {
			init = p.Declaration()
		} else {
			p.consumeIfExists(token.SEMICOLON)
		}
		if !p.match(token.LBRACE) {
			if !p.consumeIfExists(token.SEMICOLON) {
				cond = p.Expression()
			}
		}
		p.consumeIfExists(token.SEMICOLON)
		if !p.match(token.LBRACE) {
			adv = p.Expression()
		}
		statement = &ForStatement{
			initialization: init,
			condition:      cond,
			advancement:    adv,
			body:           p.Statement(),
		}
	} else if p.consumeIfExists(token.LBRACE) {
		declarations := []Declaration{}
		for !p.consumeIfExists(token.RBRACE) {
			declarations = append(declarations, p.Declaration())
		}
		// if p.index == len(p.tokens) &&
		statement = &BlockStatement{
			declarations: declarations,
		}
	} else if p.consumeIfExists(token.PRINT) {
		statement = p.printStatement()
	} else if p.consumeIfExists(token.BREAK) {
		statement = &BreakStatement{}
		p.consume(token.SEMICOLON, "Expected a semicolon")
	} else if p.consumeIfExists(token.CONTINUE) {
		statement = &ContinueStatement{}
		p.consume(token.SEMICOLON, "Expected a semicolon")
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

func (p *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.currentToken.Type == t {
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
