package object

import "fmt"

type Number struct {
	Value float64
}

func (n Number) String() string {
	// if v is int, print as int
	if n.Value == float64(int(n.Value)) {
		return fmt.Sprintf("%d", int(n.Value))
	}
	return fmt.Sprintf("%v", n.Value)
}

func (n Number) Type() ObjectType {
	return NUMBER
}

func (n Number) Add(o Object) (Object, error) {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value + v.Value}, nil
	case String:
		return NewString(n.String() + v.Value), nil
	default:
		return Nil{}, fmt.Errorf("cannot add Number and %s", o.Type())
	}
}

func (n Number) Sub(o Object) (Object, error) {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value - v.Value}, nil
	default:
		return Nil{}, fmt.Errorf("cannot subtract %s from Number", o.Type())
	}
}

func (n Number) Mul(o Object) (Object, error) {
	switch v := o.(type) {
	case Number:
		return Number{Value: n.Value * v.Value}, nil
	default:
		return Nil{}, fmt.Errorf("cannot multiply Number by %s", o.Type())
	}
}

func (n Number) Div(o Object) (Object, error) {
	switch v := o.(type) {
	case Number:
		if v.Value == 0 {
			return Nil{}, fmt.Errorf("division by zero")
		}
		return Number{Value: n.Value / v.Value}, nil
	default:
		return Nil{}, fmt.Errorf("cannot divide Number by %s", o.Type())
	}
}

func (n Number) GetTruthy() Boolean {
	return Boolean{Value: n.Value != 0}
}

func (n Number) GetPrototype() *Map {
	return nil
}
