package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"fmt"
	"net/http"
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
			Name:  "clock",
		},

		NativeFunction{
			Function: func(vm *VM, args ...object.Object) object.Object {
				// Return current time in seconds
				closure := args[0]
				port := args[1]
				if _, ok := closure.(object.Closure); !ok {
					panic("First argument to listen must be a function")
				}
				p, ok := port.(object.Number)
				if !ok {
					panic("Second argument to listen must be a number")
				}
				eventIndex := vm.RegisterEvent(closure.(object.Closure))
				helloHandler := func(w http.ResponseWriter, r *http.Request) {
					vm.FireEvent(eventIndex, object.String{Value: r.URL.Path}, object.String{Value: r.URL.Path})
					fmt.Fprintf(w, "Hello, World!")
				}
				http.HandleFunc("/", helloHandler)

				go func() {
					if err := http.ListenAndServe(fmt.Sprintf(":%d", int(p.Value)), nil); err != nil {
						panic(err)
					}
				}()
				return object.Nil{}
			},
			Arity: 2,
			Name:  "listen",
		},
	}
}
