package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"time"
)

func getAsync() *object.Map {
	async := &object.Map{Map: map[string]object.Object{}}

	async.Map["setInterval"] = NativeFunction{
		Name:  "setInterval",
		Arity: 2,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.CLOSURE, "setInterval", 0)
			vm.assertArgumentToType(args[1], object.NUMBER, "setInterval", 1)

			fn := args[0].(object.Closure)
			delay := int(args[1].(object.Number).Value)
			vm.RegisterEvent()

			go func() {
				for {
					time.Sleep(time.Duration(delay) * time.Millisecond)
					vm.FireEvent(fn)
				}
			}()

			return object.Nil{}
		},
	}

	async.Map["setTimeout"] = NativeFunction{
		Name:  "setTimeout",
		Arity: 2,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.CLOSURE, "setTimeout", 0)
			vm.assertArgumentToType(args[1], object.NUMBER, "setTimeout", 1)

			fn := args[0].(object.Closure)
			delay := int(args[1].(object.Number).Value)
			vm.RegisterEvent()

			go func() {
				defer vm.DetachEvent()
				time.Sleep(time.Duration(delay) * time.Millisecond)
				vm.FireEvent(fn)
			}()

			return object.Nil{}
		},
	}

	return async
}
