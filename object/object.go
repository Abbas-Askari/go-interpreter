package object

import "fmt"

type Object interface {
	Add(Object) Object
	Sub(Object) Object
	Mul(Object) Object
	Div(Object) Object
}

type Number struct {
	Value float64
}

type Nil struct{}

func (n Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n Number) Add(o Object) Object {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value + v.Value}
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

func (n Nil) Add(o Object) Object {
	return Nil{}
}

func (n Nil) Sub(o Object) Object {
	return Nil{}
}

func (n Nil) Mul(o Object) Object {
	return Nil{}
}

func (n Nil) Div(o Object) Object {
	return Nil{}
}
