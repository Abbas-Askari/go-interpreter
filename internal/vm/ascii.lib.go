package vm

import (
	"Abbas-Askari/interpreter-v2/internal/object"
)

func getAscii() *object.Map {
	asciiLib := &object.Map{Map: map[string]object.Object{}}

	asciiLib.Map["charCode"] = NativeFunction{
		Name:  "charCode",
		Arity: 2,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "charCode", 0)
			vm.assertArgumentToType(args[1], object.NUMBER, "charCode", 1)

			return object.Number{Value: float64(args[0].(object.String).Value[int(args[1].(object.Number).Value)])}
		},
	}

	asciiLib.Map["fromCharCode"] = NativeFunction{
		Name:  "fromCharCode",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.NUMBER, "fromCharCode", 0)
			return object.NewString(string(rune(int(args[0].(object.Number).Value))))
		},
	}

	return asciiLib
}
