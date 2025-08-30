package object

import (
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

type Function struct {
	Value  string
	Stream []op.OpCode
}

func (b Function) String() string {
	return fmt.Sprintf("FN<%v>", b.Stream)
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
