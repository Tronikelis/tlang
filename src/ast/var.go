package ast

type VarValue interface {
	Var()
}

type Identifier = string

type VarDec struct {
	Identifier
	Value VarValue
}

type VarString struct {
	value string
}

func (*VarString) Var() {}

type VarNumber struct {
	value float64
}

func (*VarNumber) Var() {}

type VarBool struct {
	value bool
}

func (*VarBool) Var() {}

type VarArr struct {
	value []VarValue
}

func (*VarArr) Var() {}

type VarMap struct {
	value map[Identifier]VarValue
}

func (*VarMap) Var() {}
