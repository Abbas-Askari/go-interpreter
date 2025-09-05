package object

import "fmt"

type Map struct {
	Map map[string]Object
}

func (b Map) String() string {
	return fmt.Sprint(b.Map)
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
