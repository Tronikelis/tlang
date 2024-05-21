package ast

type VarValue interface {
	Var()
}

type Identifier = string

type VarString struct {
	value string
}

func (v VarString) Var() {}

type VarNumber struct {
	value float64
}

func (v VarNumber) Var() {}

type VarBool struct {
	value bool
}

func (v VarBool) Var() {}

type VarArr struct {
	value []VarValue
}

func (v VarArr) Var() {}

type VarMap struct {
	value map[Identifier]VarValue
}

func (v VarMap) Var() {}

type VarFunc struct {
	value Body
	args  []Identifier
}

func (v VarFunc) Var() {}
