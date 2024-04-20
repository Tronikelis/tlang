package lexer

import (
	"fmt"
)

type TokenType = int

const (
	LET TokenType = iota
	IDENTIFIER

	FUNCTION
	IF
	WHILE
	FOR
	IN
	RETURN

	NUMBER
	BOOL
	STRING

	EQUALS
	PLUS
	MINUS
	SLASH
	STAR

	COMMA
	OPEN_PARENTHESES
	CLOSE_PARENTHESES
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_SQUARE
	CLOSE_SQUARE
)

type Token struct {
	Raw  string
	Type int
}

type Parser struct {
	script  string
	pointer uint
}

func NewParser(script string) *Parser {
	return &Parser{
		pointer: 0,
		script:  script,
	}
}

func isWhitespace(r byte) bool {
	return r == '\n' || r == ' '
}

func checkType(raw string) TokenType {
	switch raw {
	case "if":
		return IF
	case "for":
		return FOR
	case "let":
		return LET
	case "=":
		return EQUALS
	case "(":
		return OPEN_PARENTHESES
	case ")":
		return CLOSE_PARENTHESES
	case "{":
		return OPEN_CURLY
	case "}":
		return CLOSE_CURLY
	case "[":
		return OPEN_SQUARE
	case "]":
		return CLOSE_SQUARE
	case "-":
		return MINUS
	case "+":
		return PLUS
	case "*":
		return STAR
	case "/":
		return SLASH

	default:
		return IDENTIFIER
	}
}

func (parser *Parser) readNext() (uint, Token) {
	pointer := parser.pointer

	char := func() byte {
		return parser.script[pointer]
	}

	for isWhitespace(char()) {
		pointer += 1
	}

	raw := ""

	for !isWhitespace(char()) {
		raw += string(char())
		pointer += 1
	}

	token := Token{
		Raw:  raw,
		Type: 0,
	}

	return pointer, token
}

func (parser *Parser) readNextAndConsume() Token {
	read, token := parser.readNext()
	parser.pointer += read
	return token
}

func (parser *Parser) Parse() []Token {
	fmt.Println(parser.readNextAndConsume())
	return []Token{}
}
