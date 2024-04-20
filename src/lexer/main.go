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

var keywordMap = map[string]TokenType{
	"if":  IF,
	"for": FOR,
	"let": LET,
	"fn":  FUNCTION,
	"=":   EQUALS,
	"(":   OPEN_PARENTHESES,
	")":   CLOSE_PARENTHESES,
	"{":   OPEN_CURLY,
	"}":   CLOSE_CURLY,
	"[":   OPEN_SQUARE,
	"]":   CLOSE_SQUARE,
	"-":   MINUS,
	"+":   PLUS,
	"*":   STAR,
	"/":   SLASH,
	",":   COMMA,
}

var maxKeywordLen = func() int {
	count := 0
	for key := range keywordMap {
		if len(key) < count {
			continue
		}

		count = len(key)
	}

	return count
}()

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
	keyword, exists := keywordMap[s]
	if !exists {
		return 0
	}

	return keyword
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

	if parseKeyword(string(current)) != 0 {
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
		lexer.pointer++

		return &Token{
			Raw:  raw,
			Type: NUMBER,
		}
	}

	// raw := ""
	//
	// for !isWhitespace(lexer.peekCurrent()) && parseKeyword(string(lexer.peekNext())) == 0 {
	// 	raw += string(lexer.peekCurrent())
	// 	lexer.pointer++
	//
	// 	maybeKeyword := parseKeyword(raw)
	//
	// 	if len(raw) <= maxKeywordLen && maybeKeyword != 0 {
	// 		return &Token{
	// 			Raw:  raw,
	// 			Type: maybeKeyword,
	// 		}
	// 	}
	// }
	//
	// return &Token{
	// 	Raw:  raw,
	// 	Type: IDENTIFIER,
	// }

	// todo: parse multi character keywords and lastly, identifiers

	return nil
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
