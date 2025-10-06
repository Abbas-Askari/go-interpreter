package object

import "fmt"

type Buffer struct {
	Value     []byte
	__proto__ *Map
}

var PrototypeBuffer *Map = &Map{
	Map: map[string]Object{
		"width": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func (b Buffer) GetPrototype() *Map {
	return b.__proto__
}

func NewBuffer(value []byte) Buffer {
	__proto__ := &Map{
		Map: map[string]Object{
			"length":    Number{Value: float64(len(value))},
			"__proto__": PrototypeBuffer,
		},
	}
	arr := Buffer{Value: value, __proto__: __proto__}
	return arr
}

func (b Buffer) String() string {
	str := "Buffer["
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

func (b Buffer) Type() ObjectType {
	return BUFFER
}

func (b Buffer) Add(o Object) (Object, error) {
	arr, ok := o.(Buffer)
	if !ok {
		return Nil{}, fmt.Errorf("can only add Buffer to Buffer, got %s", o.Type())
	}
	return NewBuffer(append(b.Value, arr.Value...)), nil
}

func (b Buffer) Sub(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot subtract from Buffer")
}

func (b Buffer) Mul(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot multiply Buffer")
}

func (b Buffer) Div(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot divide Buffer")
}

func (b Buffer) GetTruthy() Boolean {
	return Boolean{len(b.Value) != 0}
}

func (b Buffer) GetElementAtIndex(i Object) (Object, error) {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return nil, fmt.Errorf("Buffer index out of range")
		}
		return Number{Value: float64(b.Value[int(idx.Value)])}, nil
	default:
		return nil, fmt.Errorf("Buffer index must be a number")
	}
}

func (b Buffer) SetElementAtIndex(i Object, o Object) error {
	switch idx := i.(type) {
	case Number:
		if idx.Value < 0 || idx.Value >= float64(len(b.Value)) {
			return fmt.Errorf("Buffer index out of range")
		}
		b.Value[int(idx.Value)] = byte(o.(Number).Value)
	default:
		return fmt.Errorf("Buffer index must be a number")
	}
	return nil
}
