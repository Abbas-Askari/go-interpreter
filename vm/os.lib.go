package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"os"
)

func getOs() *object.Map {
	osLib := &object.Map{Map: map[string]object.Object{}}

	osLib.Map["exit"] = NativeFunction{
		Name:  "exit",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.NUMBER, "exit", 0)
			code := args[0].(object.Number).Value
			os.Exit(int(code))
			return object.Nil{}
		},
	}

	return osLib
}
