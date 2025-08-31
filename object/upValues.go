package object

import (
	"fmt"
)

type UpValue struct {
	Value *Object
}

func (b UpValue) String() string {
	return fmt.Sprintf("UPVALUE<%v>", b.Value)
}

func (b UpValue) Type() ObjectType {
	return UPVALUE
}

func (b UpValue) Add(o Object) Object {
	panic("Cannot add Functions")
}

func (b UpValue) Sub(o Object) Object {
	panic("Cannot add Functions")
}

func (b UpValue) Mul(o Object) Object {
	panic("Cannot add Functions")
}

func (b UpValue) Div(o Object) Object {
	panic("Cannot add Functions")
}

func (b UpValue) GetTruthy() Boolean {
	return Boolean{true}
}
