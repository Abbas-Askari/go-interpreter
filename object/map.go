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

func (b Map) GetElementAtIndex(i Object) (Object, error) {
	index := i.String()
	val, ok := b.Map[index]
	if !ok {
		return Nil{}, nil
	}
	return val, nil
}

func (b Map) SetElementAtIndex(i Object, o Object) error {
	index := i.String()
	b.Map[index] = o
	return nil
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
