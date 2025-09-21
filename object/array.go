package object

import "fmt"

type Array struct {
	Value     []Object
	__proto__ *Map
}

func NewArray(value []Object) Array {
	__proto__ := &Map{
		Map: map[string]Object{
			"length":    Number{Value: float64(len(value))},
			"__proto__": PrototypeArray,
		},
	}
	arr := Array{Value: value, __proto__: __proto__}
	return arr
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
	return NewArray(append(b.Value, arr.Value...))
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

func (b Array) GetElementAtIndex(i Object) (Object, error) {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return nil, fmt.Errorf("Array index out of range")
		}
		return b.Value[int(idx.Value)], nil
	default:
		return nil, fmt.Errorf("Array index must be a number")
	}
}

func (b Array) SetElementAtIndex(i Object, o Object) error {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return fmt.Errorf("Array index out of range")
		}
		b.Value[int(idx.Value)] = o
	default:
		return fmt.Errorf("Array index must be a number")
	}
	return nil
}

var PrototypeArray *Map = &Map{
	Map: map[string]Object{
		"width": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func (b Array) GetPrototype() *Map {
	return b.__proto__
}
