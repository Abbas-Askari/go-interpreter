package object

import "fmt"

type ObjectType string

const (
	NUMBER   = "NUMBER"
	BOOLEAN  = "BOOLEAN"
	STRING   = "STRING"
	NIL      = "NIL"
	FUNCTION = "FUNCTION"
	CLOSURE  = "CLOSURE"
	UPVALUE  = "UPVALUE"
	MAP      = "MAP"
)

type Object interface {
	Add(Object) Object
	Sub(Object) Object
	Mul(Object) Object
	Div(Object) Object
	GetTruthy() Boolean
	String() string
	Type() ObjectType
	GetPrototype() *Map
}

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
		return String{Value: n.String() + v.Value}
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

type Nil struct{}

func (n Nil) Type() ObjectType {
	return NIL
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

func (n Nil) String() string {
	return "nil"
}

func (n Nil) GetTruthy() Boolean {
	return Boolean{Value: false}
}

func (n Nil) GetPrototype() *Map {
	return nil
}
