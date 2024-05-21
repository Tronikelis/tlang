package ast

type Environment = map[Identifier]VarValue

type Body struct {
	Environment
}

type Ast struct {
}
