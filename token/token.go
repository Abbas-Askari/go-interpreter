package token

import (
	"Abbas-Askari/interpreter-v2/colors"
	"fmt"
)

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 1343456

	ASSIGN   TokenType = "="
	MULTIPLY TokenType = "*"
	SLASH    TokenType = "/"
	PERCENT  TokenType = "%"
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	NOT      TokenType = "!"
	AND      TokenType = "&&"
	OR       TokenType = "||"

	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	DOT       TokenType = "."

	EQUAL_EQUAL   TokenType = "=="
	NOT_EQUAL     TokenType = "!="
	LESS_EQUAL    TokenType = "<="
	GREATER_EQUAL TokenType = ">="
	LESS          TokenType = "<"
	GREATER       TokenType = ">"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	FUNCTION TokenType = "FUNCTION"
	RETURN   TokenType = "RETURN"
	LET      TokenType = "LET"

	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
	NIL   TokenType = "NIL"

	NUMBER     TokenType = "NUMBER"
	STRING     TokenType = "STRING"
	IDENTIFIER TokenType = "IDENTIFIER"

	PRINT    TokenType = "PRINT"
	IF       TokenType = "IF"
	FOR      TokenType = "FOR"
	BREAK    TokenType = "BREAK"
	CONTINUE TokenType = "CONTINUE"
	ELSE     TokenType = "ELSE"
	IMPORT   TokenType = "IMPORT"
)

type Token struct {
	Type    TokenType
	Literal string
	// Value   any
}

func (t Token) String() string {
	return fmt.Sprintf("<%v: %v>", colors.Colorize(string(t.Type), colors.GREEN), colors.Colorize(t.Literal, colors.GREEN))
}
