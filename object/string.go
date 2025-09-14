package object

import "fmt"

type String struct {
	Value string
}

func (b String) String() string {
	return b.Value
}

func (b String) Type() ObjectType {
	return STRING
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

func (b String) GetElementAtIndex(i int) Object {
	if i < 0 || i >= len(b.Value) {
		panic("String index out of range")
	}
	return String{Value: string(b.Value[i])}
}

func (b String) SetElementAtIndex(i int, o Object) {
	if i < 0 || i >= len(b.Value) {
		panic("String index out of range")
	}
	str, ok := o.(String)
	if !ok || len(str.Value) != 1 {
		panic("Can only assign single character strings to string indices")
	}
	b.Value = b.Value[:i] + str.Value + b.Value[i+1:]
}

var PrototypeString *Map = &Map{
	Map: map[string]Object{
		"width": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func (b String) GetPrototype() *Map {
	return &Map{
		Map: map[string]Object{
			"length":    Number{Value: float64(len(b.Value))},
			"__proto__": PrototypeString,
		},
	}
}
