package lexer

import (
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

func Tokenize(input string) []token.Token {
	// This is a placeholder implementation.
	// In a real lexer, you would parse the input string into tokens.

	mapping := map[string]token.TokenType{
		"+":   token.PLUS,
		"-":   token.MINUS,
		"*":   token.MULTIPLY,
		"/":   token.SLASH,
		";":   token.SEMICOLON,
		"(":   token.LPAREN,
		")":   token.RPAREN,
		"fun": token.FUNCTION,
	}

	tokens := []token.Token{}
	i := 0
	for i != len(input) {
		for str, tokenType := range mapping {
			if i+len(str) >= len(input) || input[i:i+len(str)] != str {
				continue
			}

			tokens = append(tokens, token.Token{
				Type:    tokenType,
				Literal: str,
			})
			i += len(str)
			continue
		}

		c := input[i]

		if '0' <= c && c <= '9' {
			number := ""
			for '0' <= c && c <= '9' {
				number = number + string(c)
				i++
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

		if c == ' ' {
			i++
			continue
		}

		panic(fmt.Errorf("UNKNOWN TOKEN: %v", i))
	}
	return tokens
}
