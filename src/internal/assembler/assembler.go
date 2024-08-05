package assembler

import (
	"fmt"
	"go-vm/src/internal/cpu"
	"go-vm/src/internal/utils"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Token struct {
	Value string
	Row   int
	Col   int
}

func isNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

func parseReg(token Token) (byte, error) {
	val := token.Value
	row, col := token.Row, token.Col

	if !strings.HasPrefix(val, "%") {
		return 0, fmt.Errorf("%d:%d Expected register got `%s` instead", row, col, val)
	}

	if !isNumber(val[1:]) {
		return 0, fmt.Errorf("%d:%d `%s` is not a valid register", row, col, val)
	}

	num, err := strconv.Atoi(val[1:])
	if err != nil || (num < 0 || num > 255) {
		return 0, fmt.Errorf("%d:%d `%s` is not a valid register", row, col, val)
	}

	return byte(num), nil
}

func parseMov(tokens []Token) ([]byte, int) {
	if len(tokens) < 4 {
		fmt.Printf("`mov` missing arguments on lin %d\n", tokens[0].Row)
		os.Exit(1)
	}

	reg1, err := parseReg(tokens[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if tokens[2].Value != "," {
		fmt.Printf("%d:%d Expected `,` got `%s` instead\n", tokens[2].Row, tokens[2].Col, tokens[2].Value)
		os.Exit(1)
	}

	if tokens[3].Value == "[" {
		if len(tokens) < 6 {
			fmt.Printf("`mov` missing arguments on lin %d\n", tokens[0].Row)
			os.Exit(1)
		}

		if tokens[5].Value != "]" {
			fmt.Printf("%d:%d Expected `]` got `%s` instead\n", tokens[5].Row, tokens[5].Col, tokens[5].Value)
			os.Exit(1)
		}

		reg2, err := parseReg(tokens[4])
		if err != nil {
			num, err := strconv.Atoi(tokens[4].Value)
			if err != nil {
				fmt.Printf("%d:%d Expected value or address got `%s` instead\n", tokens[4].Row, tokens[4].Col, tokens[4].Value)
				os.Exit(1)
			}
			return append([]byte{cpu.LOAD_VAL, reg1}, utils.IntToBytes(int64(num))...), 6
		}
		return []byte{cpu.MOV_VAL, reg1, reg2}, 6
	} else {
		reg2, err := parseReg(tokens[3])
		if err != nil {
			num, err := strconv.Atoi(tokens[3].Value)
			if err != nil {
				fmt.Printf("%d:%d Expected value or address got `%s` instead\n", tokens[3].Row, tokens[3].Col, tokens[3].Value)
				os.Exit(1)
			}
			return append([]byte{cpu.LOAD, reg1}, utils.IntToBytes(int64(num))...), 4
		}
		return []byte{cpu.MOV, reg1, reg2}, 4
	}
}

func Parse(tokens []Token) []byte {
	var bytecode []byte

	for len(tokens) != 0 {
		val := tokens[0].Value
		switch val {
		case "mov":
			parsed, skip := parseMov(tokens)
			tokens = tokens[skip:]
			bytecode = append(bytecode, parsed...)
		case "push":
			fallthrough
		case "store":
			fallthrough
		case "inc":
			fallthrough
		case "add":
			fallthrough
		case "sub":
			fallthrough
		case "mul":
			fallthrough
		case "div":
			fallthrough
		case "jmp":
			fallthrough
		case "cmp":
			fallthrough
		case "je":
			fallthrough
		case "jne":
			fallthrough
		case "jg":
			fallthrough
		case "jl":
			fallthrough
		case "jge":
			fallthrough
		case "jle":
			fallthrough
		case "neg":
			fallthrough
		case "and":
			fallthrough
		case "or":
			fallthrough
		case "xor":
			fallthrough
		case "not":
			fallthrough
		case "syscall":
			fallthrough
		case "halt":
			fallthrough
		default:
			tokens = tokens[1:]
		}

		// if len(tokens) != 0 {
		// 	if tokens[0].Value != "\n" {
		// 		fmt.Printf("%d:%d Expected new line got `%s` instead\n", tokens[0].Row, tokens[0].Col, tokens[0].Value)
		// 		os.Exit(1)
		// 	}
		// 	tokens = tokens[1:]
		// }
	}

	return bytecode
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
