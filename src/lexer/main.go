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

func (lexer *Lexer) readAtPointer() byte {
	if len(lexer.script)-1 < lexer.pointer {
		return 0
	}

	b := lexer.script[lexer.pointer]

	return b
}

func (lexer *Lexer) consumeWhitespace() {
	for isWhitespace(lexer.readAtPointer()) {
		lexer.pointer++
	}
}

func (lexer *Lexer) consumeNonWhitespace() string {
	current := lexer.readAtPointer()
	read := ""

	for !isWhitespace(current) {
		read += string(current)
		lexer.pointer++
		current = lexer.readAtPointer()
	}

	return read
}

func (lexer *Lexer) readNext() *Token {
	lexer.consumeWhitespace()

	current := lexer.readAtPointer()
	if current == 0 {
		return nil
	}

	if current == '"' {
		// parse string
		return &Token{}
	}

	if isNumber(current) {
		// parse number
		return &Token{}
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
