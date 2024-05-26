package ast

type Statement interface {
	Statement()
}

type VarStatement struct {
	identifier Identifier
	value      VarValue
}

func (v VarStatement) Statement()

type IfStatement struct {
	condition VarValue
	tru       Body
	fal       Body
}

func (v IfStatement) Statement()

type WhileStatement struct {
	condition VarValue
	body      Body
}

func (v WhileStatement) Statement()

type ForStatement struct {
	key  VarNumber
	item VarValue
	over VarValue
	body Body
}

func (v ForStatement) Statement()
