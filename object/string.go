package object

import "fmt"

type String struct {
	Value string
}

func (b String) String() string {
	return b.Value
}

func (b String) Add(o Object) Object {
	return String{Value: b.Value + fmt.Sprint(o)}
}

func (b String) Sub(o Object) Object {
	panic("Cannot add Strings")
}

func (b String) Mul(o Object) Object {
	panic("Cannot add Strings")
}

func (b String) Div(o Object) Object {
	panic("Cannot add Strings")
}

func (b String) GetTruthy() Boolean {
	return Boolean{len(b.Value) != 0}
}
