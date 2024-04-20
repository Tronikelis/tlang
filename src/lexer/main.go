package lexer

type TokenType = int

const (
	LET TokenType = iota + 1
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

// 0 means unknown keyword
func parseKeyword(s string) TokenType {
	switch s {
	case "if":
		return IF
	case "for":
		return FOR
	case "let":
		return LET
	case "fn":
		return FUNCTION
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
	}

	return 0
}

func (lexer *Lexer) parseString() string {
	raw := ""

	shouldContinue := func() bool {
		if lexer.peekCurrent() == '\\' && lexer.peekNext() == '"' {
			return true
		}

		return lexer.peekNext() != '"'
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
		lexer.pointer++

		return &Token{
			Raw:  raw,
			Type: NUMBER,
		}
	}

	nonWhitespace := lexer.consumeNonWhitespace()

	keyword := parseKeyword(nonWhitespace)
	if keyword != 0 {
		return &Token{
			Raw:  nonWhitespace,
			Type: keyword,
		}
	}

	return &Token{
		Raw:  nonWhitespace,
		Type: IDENTIFIER,
	}
}

func (lexer *Lexer) Parse() []Token {
	return []Token{}
}
