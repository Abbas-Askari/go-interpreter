package lexer

import (
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"unicode"
)

func Tokenize(input string) []token.Token {
	// This is a placeholder implementation.
	// In a real lexer, you would parse the input string into tokens.

	mapping := map[string]token.TokenType{
		"+":     token.PLUS,
		"-":     token.MINUS,
		"*":     token.MULTIPLY,
		"/":     token.SLASH,
		";":     token.SEMICOLON,
		"(":     token.LPAREN,
		"=":     token.ASSIGN,
		")":     token.RPAREN,
		"fun":   token.FUNCTION,
		"print": token.PRINT,
		"let":   token.LET,
	}

	tokens := []token.Token{}
	i := 0
	for i != len(input) {
		foundFromMapping := false
		for str, tokenType := range mapping {

			if i+len(str) > len(input) {
				continue
			}

			test := input[i : i+len(str)]

			if test != str {
				continue
			}

			tokens = append(tokens, token.Token{
				Type:    tokenType,
				Literal: str,
			})
			i += len(str)
			foundFromMapping = true
			break
		}

		if foundFromMapping {
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
