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
			Name:  "BufferFrom",
			Arity: 1,
			Function: func(vm *VM, args ...object.Object) object.Object {
				switch val := args[0].(type) {
				case object.String:
					return object.NewBuffer([]byte(val.Value))
				case object.Array:
					bytes := make([]byte, len(val.Value))
					for i, v := range val.Value {
						switch num := v.(type) {
						case object.Number:
							bytes[i] = byte(num.Value)
						default:
							vm.runtimeError("Array elements must be numbers to convert to Buffer")
							return object.Nil{}
						}
					}
					return object.NewBuffer(bytes)
				default:
					vm.runtimeError("Argument to Buffer.from must be a string or an array of numbers")
					return object.Nil{}
				}
			},
		},
		NativeFunction{
			Name:  "BufferSlice",
			Arity: 3,
			Function: func(vm *VM, args ...object.Object) object.Object {
				vm.assertArgumentToType(args[0], object.BUFFER, "Buffer.slice", 0)
				vm.assertArgumentToType(args[1], object.NUMBER, "Buffer.slice", 1)
				vm.assertArgumentToType(args[2], object.NUMBER, "Buffer.slice", 2)
				buf := args[0].(object.Buffer)
				start := int(args[1].(object.Number).Value)
				end := int(args[2].(object.Number).Value)
				if start < 0 || end > len(buf.Value) || start > end {
					vm.runtimeError("Invalid slice indices")
					return object.Nil{}
				}
				newBuf := make([]byte, end-start)
				copy(newBuf, buf.Value[start:end])
				return object.NewBuffer(newBuf)
			},
		},
		NativeFunction{
			Name:  "BufferToString",
			Arity: 1,
			Function: func(vm *VM, args ...object.Object) object.Object {
				vm.assertArgumentToType(args[0], object.BUFFER, "Buffer.toString", 0)
				buf := args[0].(object.Buffer)
				return object.NewString(string(buf.Value))
			},
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
			Name:  "readLineSync",
		},
		NativeFunction{
			Function: func(vm *VM, args ...object.Object) object.Object {
				vm.assertArgumentToType(args[0], object.CLOSURE, "readLine", 0)
				closure := args[0].(object.Closure)
				vm.RegisterEvent()
				go func() {
					reader := bufio.NewReader(os.Stdin)
					input, err := reader.ReadString('\n')
					if err != nil {
						vm.FireEvent(closure, object.Nil{}, object.NewString(fmt.Sprintf("Error reading input: %v", err)))
						vm.DetachEvent()
					}
					input = strings.TrimRight(input, "\r\n")
					vm.FireEvent(closure, object.NewString(input), object.Nil{})
					vm.DetachEvent()
				}()
				return object.Nil{}
			},
			Arity: 1,
			Name:  "readLine",
		},
	}
}
