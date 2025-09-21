package object

import "fmt"

type String struct {
	Value     string
	prototype *Map
}

var PrototypeString *Map = &Map{
	Map: map[string]Object{
		"width": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func NewString(value string) String {
	__proto__ := &Map{
		Map: map[string]Object{
			"length":    Number{Value: float64(len(value))},
			"__proto__": PrototypeString,
		},
	}
	s := String{Value: value, prototype: __proto__}
	return s
}

func (b String) String() string {
	return b.Value
}

func (b String) Type() ObjectType {
	return STRING
}

func (b String) Add(o Object) Object {
	return NewString(b.Value + fmt.Sprint(o))
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

func (b String) GetElementAtIndex(i Object) (Object, error) {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return nil, fmt.Errorf("String index out of range")
		}
		return NewString(string(b.Value[int(idx.Value)])), nil
	default:
		return nil, fmt.Errorf("String index must be a number")
	}
}

func (b String) SetElementAtIndex(i Object, o Object) error {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return fmt.Errorf("String index out of range")
		}
		// Set the element at the specified index
		str, ok := o.(String)
		if !ok || len(str.Value) != 1 {
			return fmt.Errorf("Can only assign single character strings to string indices")
		}
		b.Value = b.Value[:int(idx.Value)] + str.Value + b.Value[int(idx.Value)+1:]
	default:
		return fmt.Errorf("String index must be a number")
	}
	return nil
}

func (b String) GetPrototype() *Map {
	return b.prototype
}
