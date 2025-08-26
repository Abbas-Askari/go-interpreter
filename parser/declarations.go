package parser

import (
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
)

type DeclarationType int

const (
	VariableDeclarationType DeclarationType = iota
	StatementDeclarationType
)

type Declaration interface {
	Emit(p interfaces.ICompiler)
	String() string
	Type() DeclarationType
}

type VariableDeclaration struct {
	name       token.Token
	expression Expression
	// globalIndex int
}

func (d *VariableDeclaration) Emit(c interfaces.ICompiler) {
	d.expression.Emit(c)
	c.Declare(d.name.Literal)
}

func (dx VariableDeclaration) String() string {
	return fmt.Sprintf("Declaration: %v = %v\n", dx.name, dx.expression)
}

func (d *VariableDeclaration) Type() DeclarationType {
	return VariableDeclarationType
}
