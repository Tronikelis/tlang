package ast

import (
	"tlang/src/tokens"
)

type Environment = map[Identifier]VarValue

type Body struct {
	environment Environment
	statements  []Statement
}

type Ast struct {
	root Body
}

func New([]tokens.Token) Ast {
	return Ast{}
}
