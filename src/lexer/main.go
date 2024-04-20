package lexer

import (
	"fmt"
	"log"
	"tlang/src/tokens"
)

var keywordMap = map[string]tokens.TokenType{
	"in":     tokens.IN,
	"if":     tokens.IF,
	"for":    tokens.FOR,
	"let":    tokens.LET,
	"fn":     tokens.FUNCTION,
	"=":      tokens.EQUALS,
	"(":      tokens.OPEN_PARENTHESES,
	")":      tokens.CLOSE_PARENTHESES,
	"{":      tokens.OPEN_CURLY,
	"}":      tokens.CLOSE_CURLY,
	"[":      tokens.OPEN_SQUARE,
	"]":      tokens.CLOSE_SQUARE,
	"-":      tokens.MINUS,
	"+":      tokens.PLUS,
	"*":      tokens.STAR,
	"/":      tokens.SLASH,
	",":      tokens.COMMA,
	"return": tokens.RETURN,
	"true":   tokens.TRUE,
	"false":  tokens.FALSE,
	"else":   tokens.ELSE,
	".":      tokens.DOT,
	"while":  tokens.WHILE,
	"nil":    tokens.NIL,
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
	return b == '\n' || b == ' ' || b == '\r' || b == '\t'
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return !isNumber(b) && !isWhitespace(b)
}

// "" means unknown keyword
func parseKeyword(s string) tokens.TokenType {
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

func (lexer *Lexer) consumeUntilWith(b byte) {
	for lexer.peekCurrent() != b {
		lexer.pointer++
	}

	lexer.pointer++
}

func (lexer *Lexer) consumeWhitespace() {
	removed := 0

	for isWhitespace(lexer.peekCurrent()) {
		lexer.pointer++
		removed++
	}

	if removed == 0 {
		return
	}

	for lexer.peekCurrent() == '/' && lexer.peekNext() == '/' {
		lexer.consumeUntilWith('\n')
		lexer.consumeWhitespace()
	}
}

func (lexer *Lexer) readNext() *tokens.Token {
	lexer.consumeWhitespace()

	current := lexer.peekCurrent()
	if current == 0 {
		return nil
	}

	if parseKeyword(string(current)) != "" {
		lexer.pointer++
		return &tokens.Token{
			Raw:  string(current),
			Type: parseKeyword(string(current)),
		}
	}

	if current == '"' {
		lexer.pointer++
		raw := lexer.parseString()
		lexer.pointer++

		return &tokens.Token{
			Raw:  raw,
			Type: tokens.STRING,
		}
	}

	if isNumber(current) {
		raw := lexer.parseNumber()

		return &tokens.Token{
			Raw:  raw,
			Type: tokens.NUMBER,
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
		return &tokens.Token{
			Raw:  raw,
			Type: maybeKeyword,
		}
	}

	return &tokens.Token{
		Raw:  raw,
		Type: tokens.IDENTIFIER,
	}

}

func (lexer *Lexer) Parse() []tokens.Token {
	tokens := []tokens.Token{}

	nextToken := lexer.readNext()
	for nextToken != nil {
		tokens = append(tokens, *nextToken)
		nextToken = lexer.readNext()
	}

	return tokens
}

func PrintTokens(tokens []tokens.Token) {
	for _, value := range tokens {
		fmt.Printf("%+v\n", value)
	}
}
