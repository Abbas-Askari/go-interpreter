package token

import (
	"Abbas-Askari/interpreter-v2/colors"
	"fmt"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	SLASH    = "/"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	NUMBER     = "NUMBER"
	IDENTIFIER = "IDENTIFIER"

	PRINT = "PRINT"
	IF    = "IF"
	FOR   = "FOR"
	ELSE  = "ELSE"
)

type Token struct {
	Type    TokenType
	Literal string
	// Value   any
}

func (t Token) String() string {
	return fmt.Sprintf("<%v: %v>", colors.Colorize(string(t.Type), colors.GREEN), colors.Colorize(t.Literal, colors.GREEN))
}
