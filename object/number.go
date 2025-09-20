package object

import "fmt"

type Number struct {
	Value float64
}

func (n Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n Number) Type() ObjectType {
	return NUMBER
}

func (n Number) Add(o Object) Object {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value + v.Value}
	case String:
		return NewString(n.String() + v.Value)
	default:
		return Nil{}
	}
}

func (n Number) Sub(o Object) Object {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value - v.Value}
	default:
		return Nil{}
	}
}

func (n Number) Mul(o Object) Object {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value * v.Value}
	default:
		return Nil{}
	}
}

func (n Number) Div(o Object) Object {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value / v.Value}
	default:
		return Nil{}
	}
}

func (n Number) GetTruthy() Boolean {
	return Boolean{Value: n.Value != 0}
}

func (n Number) GetPrototype() *Map {
	return nil
}
