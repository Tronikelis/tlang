package ast

import (
	"tlang/src/tokens"
)

type Body struct {
	statements []Statement
}

type Ast struct {
	root Body
}

func New([]tokens.Token) Ast {
	return Ast{}
}
