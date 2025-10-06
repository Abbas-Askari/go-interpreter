package object

import "fmt"

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
	return b.Function.String()
}

func (b Closure) Type() ObjectType {
	return CLOSURE
}

func (b Closure) Add(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Closure")
}

func (b Closure) Sub(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Closure")
}

func (b Closure) Mul(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Closure")
}

func (b Closure) Div(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on Closure")
}

func (b Closure) GetTruthy() Boolean {
	return Boolean{true}
}

func (b Closure) GetPrototype() *Map {
	return b.Function.GetPrototype()
}
