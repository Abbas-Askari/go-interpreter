package object

import (
	"Abbas-Askari/interpreter-v2/colors"
	"fmt"
)

type Map struct {
	Map map[string]Object
}

func (b Map) String() string {
	str := fmt.Sprint(colors.Colorize("{", colors.RESET))
	i := 0
	for k, v := range b.Map {
		str += fmt.Sprintf("%v: %v", k, v)
		if i != len(b.Map)-1 {
			str += ", "
		}
		i++
	}
	str += fmt.Sprint(colors.Colorize("}", colors.RESET))
	return str
}

func (b Map) GetElementAtIndex(i Object) Object {
	switch idx := i.(type) {
	case String:
		val, ok := b.Map[idx.Value]
		if !ok {
			return Nil{}
		}
		return val
	default:
		panic("Map index must be a string")
	}
}

func (b Map) SetElementAtIndex(i Object, o Object) {
	switch idx := i.(type) {
	case String:
		b.Map[idx.Value] = o
	default:
		panic("Map index must be a string")
	}
}

func (b Map) Type() ObjectType {
	return MAP
}

func (b Map) Add(o Object) Object {
	panic("Cannot add Maps")
}

func (b Map) Sub(o Object) Object {
	panic("Cannot add Maps")
}

func (b Map) Mul(o Object) Object {
	panic("Cannot add Maps")
}

func (b Map) Div(o Object) Object {
	panic("Cannot add Maps")
}

func (b Map) GetTruthy() Boolean {
	return Boolean{true}
}

func (b Map) GetPrototype() *Map {
	m, ok := b.Map["__proto__"]
	if !ok {
		return nil
	}
	mPtr, ok := m.(Map)
	if !ok {
		mPtr, ok := m.(*Map)
		if !ok {
			return nil
		}
		return mPtr
	}
	return &mPtr
}
