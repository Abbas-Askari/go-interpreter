package object

import (
	"fmt"
)

type UpValue struct {
	Value  *Object
	Closed Object
}

func (b UpValue) String() string {
	if b.Value == nil {
		return fmt.Sprintf("UPVALUE<closed: %v>", b.Closed)
	}
	return fmt.Sprintf("UPVALUE<%p>", b.Value)
}

func (b UpValue) Type() ObjectType {
	return UPVALUE
}

func (b UpValue) Add(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on UpValue")
}

func (b UpValue) Sub(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on UpValue")
}

func (b UpValue) Mul(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on UpValue")
}

func (b UpValue) Div(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on UpValue")
}

func (b UpValue) GetTruthy() Boolean {
	return Boolean{true}
}

func (b UpValue) GetPrototype() *Map {
	return nil
}
