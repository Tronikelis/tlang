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

type Parser struct {
	script  string
	pointer int
}

func NewParser(script string) *Parser {
	return &Parser{
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

func (parser *Parser) readAtPointer() byte {
	if len(parser.script)-1 < parser.pointer {
		return 0
	}

	b := parser.script[parser.pointer]

	return b
}

func (parser *Parser) consumeWhitespace() {
	for isWhitespace(parser.readAtPointer()) {
		parser.pointer++
	}
}

func (parser *Parser) readNext() *Token {
	parser.consumeWhitespace()

	current := parser.readAtPointer()
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

	raw := ""

	for !isWhitespace(current) {
		raw += string(current)
		parser.pointer++
		current = parser.readAtPointer()
	}

	keyword := parseKeyword(raw)
	if keyword != 0 {
		return &Token{
			Raw:  raw,
			Type: keyword,
		}
	}

	return &Token{
		Raw:  raw,
		Type: IDENTIFIER,
	}
}

func (parser *Parser) Parse() []Token {
	return []Token{}
}
