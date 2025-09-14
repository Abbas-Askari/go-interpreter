package object

import "fmt"

type Array struct {
	Value []Object
}

func (b Array) String() string {
	str := "["
	for i, v := range b.Value {
		// fmt.Println(i, v)
		str += fmt.Sprint(v)
		if i < len(b.Value)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

func (b Array) Type() ObjectType {
	return ARRAY
}

func (b Array) Add(o Object) Object {
	arr, ok := o.(Array)
	if !ok {
		panic("Can only add Array to Array")
	}
	return Array{Value: append(b.Value, arr.Value...)}
}

func (b Array) Sub(o Object) Object {
	panic("Cannot subtract Arrays")
}

func (b Array) Mul(o Object) Object {
	panic("Cannot multiply Arrays")
}

func (b Array) Div(o Object) Object {
	panic("Cannot divide Arrays")
}

func (b Array) GetTruthy() Boolean {
	return Boolean{len(b.Value) != 0}
}

func (b Array) GetElementAtIndex(i int) Object {
	if i < 0 || i >= len(b.Value) {
		panic("Array index out of range")
	}
	return b.Value[i]
}

func (b Array) SetElementAtIndex(i int, o Object) {
	if i < 0 || i >= len(b.Value) {
		panic("Array index out of range")
	}
	b.Value[i] = o
}

var PrototypeArray *Map = &Map{
	Map: map[string]Object{
		"width": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func (b Array) GetPrototype() *Map {
	return &Map{
		Map: map[string]Object{
			"length":    Number{Value: float64(len(b.Value))},
			"__proto__": PrototypeArray,
		},
	}
}
