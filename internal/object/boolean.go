package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b Boolean) Add(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Boolean")
}

func (b Boolean) Sub(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Boolean")
}

func (b Boolean) Mul(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Boolean")
}

func (b Boolean) Div(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Boolean")
}

func (b Boolean) GetTruthy() Boolean {
	return b
}

func (b Boolean) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (b Boolean) Type() ObjectType {
	return BOOLEAN
}

func (b Boolean) GetPrototype() *Map {
	return nil
}
