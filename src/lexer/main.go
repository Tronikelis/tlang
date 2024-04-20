package lexer

import (
	"fmt"
	"log"
)

type TokenType = string

const (
	LET               TokenType = "let"
	IDENTIFIER                  = "identifier"
	FUNCTION                    = "function"
	IF                          = "if"
	WHILE                       = "while"
	FOR                         = "for"
	IN                          = "in"
	RETURN                      = "return"
	NUMBER                      = "number"
	BOOL                        = "bool"
	STRING                      = "string"
	EQUALS                      = "equals"
	PLUS                        = "plus"
	MINUS                       = "minus"
	SLASH                       = "slash"
	STAR                        = "star"
	COMMA                       = "comma"
	OPEN_PARENTHESES            = "open_parentheses"
	CLOSE_PARENTHESES           = "close_parentheses"
	OPEN_CURLY                  = "open_curly"
	CLOSE_CURLY                 = "close_curly"
	OPEN_SQUARE                 = "open_square"
	CLOSE_SQUARE                = "close_square"
	TRUE                        = "true"
	FALSE                       = "false"
	ELSE                        = "else"
)

var keywordMap = map[string]TokenType{
	"in":     IN,
	"if":     IF,
	"for":    FOR,
	"let":    LET,
	"fn":     FUNCTION,
	"=":      EQUALS,
	"(":      OPEN_PARENTHESES,
	")":      CLOSE_PARENTHESES,
	"{":      OPEN_CURLY,
	"}":      CLOSE_CURLY,
	"[":      OPEN_SQUARE,
	"]":      CLOSE_SQUARE,
	"-":      MINUS,
	"+":      PLUS,
	"*":      STAR,
	"/":      SLASH,
	",":      COMMA,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
}

type Token struct {
	Raw  string
	Type TokenType
}

type Lexer struct {
	script  string
	pointer int
}

func NewLexer(script string) *Lexer {
	return &Lexer{
		pointer: 0,
		script:  script,
	}
}

func isWhitespace(b byte) bool {
	return b == '\n' || b == ' ' || b == '\r'
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return !isNumber(b) && !isWhitespace(b)
}

// "" means unknown keyword
func parseKeyword(s string) TokenType {
	keyword, exists := keywordMap[s]
	if !exists {
		return ""
	}

	return keyword
}

func (lexer *Lexer) parseString() string {
	raw := ""

	shouldContinue := func() bool {
		if lexer.peekCurrent() == 0 {
			log.Fatal("string literal did not end")
		}

		if lexer.peekCurrent() == '\\' && lexer.peekNext() == '"' {
			lexer.pointer++
			return true
		}

		return lexer.peekCurrent() != '"'
	}

	for shouldContinue() {
		raw += string(lexer.peekCurrent())
		lexer.pointer++
	}

	return raw
}

func (lexer *Lexer) parseNumber() string {
	raw := ""

	for isNumber(lexer.peekCurrent()) || lexer.peekCurrent() == '.' {
		raw += string(lexer.peekCurrent())
		lexer.pointer++
	}

	return raw
}

func (lexer *Lexer) peekNext() byte {
	lexer.pointer++
	peeked := lexer.peekCurrent()
	lexer.pointer--
	return peeked

}

func (lexer *Lexer) peekCurrent() byte {
	if len(lexer.script)-1 < lexer.pointer {
		return 0
	}

	b := lexer.script[lexer.pointer]

	return b
}

func (lexer *Lexer) consumeWhitespace() {
	for isWhitespace(lexer.peekCurrent()) {
		lexer.pointer++
	}
}

func (lexer *Lexer) consumeNonWhitespace() string {
	current := lexer.peekCurrent()
	read := ""

	for !isWhitespace(current) {
		read += string(current)
		lexer.pointer++
		current = lexer.peekCurrent()
	}

	return read
}

func (lexer *Lexer) readNext() *Token {
	lexer.consumeWhitespace()

	current := lexer.peekCurrent()
	if current == 0 {
		return nil
	}

	if parseKeyword(string(current)) != "" {
		lexer.pointer++
		return &Token{
			Raw:  string(current),
			Type: parseKeyword(string(current)),
		}
	}

	if current == '"' {
		lexer.pointer++
		raw := lexer.parseString()
		lexer.pointer++

		return &Token{
			Raw:  raw,
			Type: STRING,
		}
	}

	if isNumber(current) {
		raw := lexer.parseNumber()

		return &Token{
			Raw:  raw,
			Type: NUMBER,
		}
	}

	raw := string(current)

	for
	// next exists
	lexer.peekNext() != 0 &&
		// next is not a keyword
		parseKeyword(string(lexer.peekNext())) == "" &&
		// next is not whitespace
		!isWhitespace(lexer.peekNext()) {

		raw += string(lexer.peekNext())
		lexer.pointer++
	}

	lexer.pointer++

	maybeKeyword := parseKeyword(raw)
	if maybeKeyword != "" {
		return &Token{
			Raw:  raw,
			Type: maybeKeyword,
		}
	}

	return &Token{
		Raw:  raw,
		Type: IDENTIFIER,
	}

}

func (lexer *Lexer) Parse() []Token {
	tokens := []Token{}

	nextToken := lexer.readNext()
	for nextToken != nil {
		tokens = append(tokens, *nextToken)
		nextToken = lexer.readNext()
	}

	return tokens
}

func PrintTokens(tokens []Token) {
	for _, value := range tokens {
		fmt.Printf("%+v\n", value)
	}
}
