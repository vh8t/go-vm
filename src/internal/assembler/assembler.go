package assembler

import (
	"fmt"
	"os"
)

type Token struct {
	Value string
	Row   int
	Col   int
}

func Lex(input string) []Token {
	var tokens []Token
	var buf string

	sRow, sCol := 1, 1
	row, col := 1, 1

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch ch {
		case '\n':
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
			tokens = append(tokens, Token{"\n", row, col})
			row++
			col = 0
		case ' ':
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
		case '"':
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
			tokens = append(tokens, Token{"\"", row, col})
			i++
			col++
			sRow, sCol = row, col
			for input[i] != '"' {
				if input[i] == '\n' {
					fmt.Printf("Unclosed string @ %d:%d\n", row, col)
					os.Exit(1)
				}
				buf += string(input[i])
				col++
				i++
			}
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
			tokens = append(tokens, Token{"\"", row, col})
		case ';':
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
			for input[i] != '\n' {
				col++
				i++
			}
			tokens = append(tokens, Token{"\n", row, col})
			row++
			col = 0
		case '[', ']', ':', ',':
			if buf != "" {
				tokens = append(tokens, Token{buf, sRow, sCol})
				buf = ""
			}
			tokens = append(tokens, Token{string(ch), row, col})
		default:
			if buf == "" {
				sRow, sCol = row, col
			}
			buf += string(ch)
		}

		col++
	}

	if buf != "" {
		tokens = append(tokens, Token{buf, sRow, sCol})
	}

	return tokens
}
