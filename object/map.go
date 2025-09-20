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
	return nil
}
