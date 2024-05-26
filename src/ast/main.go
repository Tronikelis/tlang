package ast

import (
	"strconv"
	"tlang/src/tokens"
)

type Body struct {
	statements []Statement
}

type Ast struct {
	root Body
}

func NewAst(t []tokens.Token) (Ast, error) {
	traverser := NewTraverser(t)

	ast := Ast{}

	for {
		token, exists := traverser.Next()
		if !exists {
			break
		}

		if token.Type == tokens.LET {
			if err := traverser.ExpectNext(tokens.EQUALS); err != nil {
				return Ast{}, err
			}

			// equal
			traverser.Next()

			if err := traverser.ExpectNext(tokens.IDENTIFIER); err != nil {
				return Ast{}, err
			}

			// identifier
			identifier, exists := traverser.Next()
			if !exists {
				panic("bruh what")
			}

			if err := traverser.ExpectNextOneOf([]tokens.TokenType{
				tokens.NUMBER,
				tokens.STRING,
				tokens.BOOL,
			}); err != nil {
				return Ast{}, err
			}

			// value
			value, exists := traverser.Next()
			if !exists {
				break
			}

			statement := VarStatement{identifier: identifier.Raw}

			switch value.Type {
			case tokens.NUMBER:
				num, err := strconv.ParseFloat(value.Raw, 64)
				if err != nil {
					return Ast{}, err
				}

				statement.value = VarNumber{value: num}
			}

			ast.root.statements = append(ast.root.statements, statement)
		}
	}

	return Ast{}, nil
}
