package object

import (
	"Abbas-Askari/interpreter-v2/colors"
	"fmt"
)

type Closure struct {
	Function Function
	UpValues []*UpValue
	This     *Object
}

func NewClosure(function Function) Closure {
	UpValues := make([]*UpValue, 0, function.UpValueCount)
	return Closure{Function: function, UpValues: UpValues}
}

func (b Closure) String() string {
	return colors.Colorize(fmt.Sprintf("FUNC<%v>", b.Function.Name), colors.BLUE)
}

func (b Closure) Type() ObjectType {
	return CLOSURE
}

func (b Closure) Add(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Sub(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Mul(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Div(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) GetTruthy() Boolean {
	return Boolean{true}
}

func (b Closure) GetPrototype() *Map {
	return nil
}
