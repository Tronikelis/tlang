package tokens

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

type Token struct {
	Raw  string
	Type TokenType
}
