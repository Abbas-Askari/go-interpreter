package object

import (
	"fmt"
	"time"
)

type NativeFunction struct {
	Function func(args ...Object) Object
	Arity    int
	Name     string
}

func (b NativeFunction) String() string {
	return fmt.Sprintf("NativeFUNCTION<%s>", b.Name)
}

func (b NativeFunction) Type() ObjectType {
	return FUNCTION
}

func (b NativeFunction) Add(o Object) Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Sub(o Object) Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Mul(o Object) Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Div(o Object) Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) GetTruthy() Boolean {
	return Boolean{true}
}

func (b NativeFunction) GetPrototype() *Map {
	return nil
}

func GetNativeFunctions() []Object {
	return []Object{
		NativeFunction{
			Function: func(args ...Object) Object {
				// Return current time in seconds
				x := Number{Value: float64(time.Now().Unix())}
				return x
			},
			Arity: 0,
			Name:  "clock",
		},
	}
}
