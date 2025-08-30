package lexer

import (
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"strings"
	"unicode"
)

func Tokenize(input string) []token.Token {
	// This is a placeholder implementation.
	// In a real lexer, you would parse the input string into tokens.

	var operators = []struct {
		literal string
		typ     token.TokenType
	}{
		{"==", token.EQUAL_EQUAL},
		{"!=", token.NOT_EQUAL},
		{"<=", token.LESS_EQUAL},
		{">=", token.GREATER_EQUAL},
		{"&&", token.AND},
		{"||", token.OR},
		{"=", token.ASSIGN},
		{"<", token.LESS},
		{">", token.GREATER},
		{"+", token.PLUS},
		{"-", token.MINUS},
		{"*", token.MULTIPLY},
		{"/", token.SLASH},
		{";", token.SEMICOLON},
		{"(", token.LPAREN},
		{")", token.RPAREN},
		{"{", token.LBRACE},
		{"}", token.RBRACE},
		{"%", token.PERCENT},
		{",", token.COMMA},
	}

	keywords := map[string]token.TokenType{
		"true":     token.TRUE,
		"false":    token.FALSE,
		"func":     token.FUNCTION,
		"print":    token.PRINT,
		"let":      token.LET,
		"if":       token.IF,
		"for":      token.FOR,
		"else":     token.ELSE,
		"break":    token.BREAK,
		"continue": token.CONTINUE,
	}

	tokens := []token.Token{}
	i := 0
	for i != len(input) {
		if i+1 < len(input) && string(input[i:i+2]) == "//" {
			for i < len(input) && input[i] != '\n' {
				i++
			}
			continue
		}

		foundOp := false
		for _, op := range operators {
			if strings.HasPrefix(input[i:], op.literal) {
				tokens = append(tokens, token.Token{
					Type:    op.typ,
					Literal: op.literal,
				})
				i += len(op.literal)
				foundOp = true
				break
			}
		}
		if foundOp {
			continue
		}

		foundFromKeywords := false
		for str, tokenType := range keywords {

			if i+len(str) > len(input) {
				continue
			}

			test := input[i : i+len(str)]

			if test != str {
				continue
			}

			if i+len(str) < len(input) {
				nextCharacter := rune(input[i+len(str)])
				if unicode.IsDigit(nextCharacter) || unicode.IsLetter(nextCharacter) || nextCharacter == '_' {
					continue
				}
			}

			tokens = append(tokens, token.Token{
				Type:    tokenType,
				Literal: str,
			})
			i += len(str)
			foundFromKeywords = true
			break
		}

		if foundFromKeywords {
			continue
		}

		c := input[i]

		if '0' <= c && c <= '9' {
			number := ""
			seenDot := false
			for '0' <= c && c <= '9' || (!seenDot && c == '.') {
				number = number + string(c)
				i++
				if c == '.' {
					seenDot = true
				}
				if i < len(input) {
					c = input[i]
				} else {
					break
				}
			}

			tokens = append(tokens, token.Token{
				Type:    "NUMBER",
				Literal: number,
			})

			continue
		}

		if c == '\'' || c == '"' {
			i++
			starting := c
			c = input[i]
			str := ""
			for c != starting {
				str = str + string(c)
				i++
				if i < len(input) {
					c = input[i]
				} else {
					panic("Incomplete string literal")
				}
			}
			i++

			tokens = append(tokens, token.Token{
				Type:    token.STRING,
				Literal: str,
			})

			continue
		}

		if unicode.IsLetter(rune(c)) || rune(c) == '_' {
			str := ""
			for unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || rune(c) == '_' {
				str = str + string(c)
				i++
				if i < len(input) {
					c = input[i]
				} else {
					break
				}
			}

			tokens = append(tokens, token.Token{
				Type:    token.IDENTIFIER,
				Literal: str,
			})

			continue
		}

		if c == ' ' || c == '\n' {
			i++
			continue
		}

		panic(fmt.Errorf("UNKNOWN TOKEN: '%v'", string(c)))
	}
	return tokens
}
