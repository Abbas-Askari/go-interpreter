package object

import (
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

type Function struct {
	Value        string
	Stream       []op.OpCode
	Arity        int
	UpValueCount int
}

func (b Function) String() string {
	return fmt.Sprintf("FUNCTION")
}

func (b Function) Type() ObjectType {
	return FUNCTION
}

func (b Function) Add(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Sub(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Mul(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Div(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) GetTruthy() Boolean {
	return Boolean{true}
}

func (b Function) GetPrototype() *Map {
	return nil
}
