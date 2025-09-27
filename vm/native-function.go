package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type NativeFunction struct {
	Function func(VM *VM, args ...object.Object) object.Object
	Arity    int
	Name     string
}

func (b NativeFunction) String() string {
	return fmt.Sprintf("NativeFUNCTION<%s>", b.Name)
}

func (b NativeFunction) Type() object.ObjectType {
	return object.FUNCTION
}

func (b NativeFunction) Add(o object.Object) object.Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Sub(o object.Object) object.Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Mul(o object.Object) object.Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) Div(o object.Object) object.Object {
	panic("Cannot add Functions")
}

func (b NativeFunction) GetTruthy() object.Boolean {
	return object.Boolean{true}
}

func (b NativeFunction) GetPrototype() *object.Map {
	return nil
}

func GetNativeFunctions() []object.Object {
	return []object.Object{
		NativeFunction{
			Function: func(vm *VM, args ...object.Object) object.Object {
				// Return current time in seconds
				x := object.Number{Value: float64(time.Now().Unix())}
				return x
			},
			Arity: 0,
			Name:  "now",
		},
		NativeFunction{
			Function: func(vm *VM, args ...object.Object) object.Object {
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					vm.runtimeError("Error reading input: %v", err)
					return object.NewString("")
				}
				input = strings.TrimRight(input, "\r\n")
				return object.NewString(input)
			},
			Arity: 0,
			Name:  "readLine",
		},
	}
}
