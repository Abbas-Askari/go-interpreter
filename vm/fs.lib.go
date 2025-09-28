package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"os"
)

func getFileSystem() *object.Map {
	fs := &object.Map{Map: map[string]object.Object{}}

	fs.Map["open"] = NativeFunction{
		Name:  "open",
		Arity: 2,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "open", 0)
			vm.assertArgumentToType(args[1], object.STRING, "open", 1)
			path := args[0].(object.String).Value
			mode := args[1].(object.String).Value
			var file *os.File
			var err error
			if mode == "r" {
				file, err = os.Open(path)
			} else if mode == "w" {
				file, err = os.Create(path)
			} else if mode == "a" {
				file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			} else {
				vm.runtimeError("In function 'open', mode must be 'r', 'w', or 'a', got '%s' instead", mode)
			}
			if err != nil {
				vm.runtimeError("Error opening file '%s': %s", path, err.Error())
			}
			return object.Map{
				Map: map[string]object.Object{
					"name": NativeFunction{
						Name:  "name",
						Arity: 0,
						Function: func(vm *VM, args ...object.Object) object.Object {
							return object.NewString(file.Name())
						},
					},
					"readAll": NativeFunction{
						Name:  "readAll",
						Arity: 0,
						Function: func(vm *VM, args ...object.Object) object.Object {
							info, err := file.Stat()
							if err != nil {
								vm.runtimeError("Error stating file: %s", err.Error())
							}
							size := info.Size()
							data := make([]byte, size)
							_, err = file.Read(data)
							if err != nil {
								vm.runtimeError("Error reading file: %s", err.Error())
							}
							return object.NewString(string(data))
						},
					},
					"read": NativeFunction{
						Name:  "read",
						Arity: 1,
						Function: func(vm *VM, args ...object.Object) object.Object {
							vm.assertArgumentToType(args[0], object.NUMBER, "read", 0)
							size := int(args[0].(object.Number).Value)
							data := make([]byte, size)
							n, err := file.Read(data)
							var errObject object.Object = object.Nil{}
							if err != nil {
								// vm.runtimeError("Error reading file: %s", err.Error())
								errObject = object.NewString(err.Error())
							}
							return object.NewArray(
								[]object.Object{
									object.String{Value: string(data[:n])},
									errObject,
								},
							)
						},
					},
					"length": NativeFunction{
						Name:  "length",
						Arity: 0,
						Function: func(vm *VM, args ...object.Object) object.Object {
							info, err := file.Stat()
							if err != nil {
								vm.runtimeError("Error stating file: %s", err.Error())
							}
							return object.Number{Value: float64(info.Size())}
						},
					},
					"write": NativeFunction{
						Name:  "write",
						Arity: 1,
						Function: func(vm *VM, args ...object.Object) object.Object {
							vm.assertArgumentToType(args[0], object.STRING, "write", 0)
							data := args[0].(object.String).Value
							_, err := file.WriteString(data)
							if err != nil {
								vm.runtimeError("Error writing to file: %s", err.Error())
							}
							return object.Nil{}
						},
					},
					"writeBytes": NativeFunction{
						Name:  "writeBytes",
						Arity: 1,
						Function: func(vm *VM, args ...object.Object) object.Object {
							vm.assertArgumentToType(args[0], object.ARRAY, "write", 0)
							elements := args[0].(object.Array).Value
							data := make([]byte, len(elements))
							for i, elem := range elements {
								vm.assertArgumentToType(elem, object.NUMBER, "writeBytes", 0)
								data[i] = byte(int(elem.(object.Number).Value))
							}
							_, err := file.Write(data)
							if err != nil {
								vm.runtimeError("Error writing to file: %s", err.Error())
							}
							return object.Nil{}
						},
					},
					"close": NativeFunction{
						Name:  "close",
						Arity: 0,
						Function: func(vm *VM, args ...object.Object) object.Object {
							err := file.Close()
							if err != nil {
								vm.runtimeError("Error closing file: %s", err.Error())
							}
							return object.Nil{}
						},
					},
				},
			}
		},
	}

	fs.Map["exists"] = NativeFunction{
		Name:  "exists",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "exists", 0)
			path := args[0].(object.String).Value
			_, err := os.Stat(path)
			if err == nil {
				return object.Boolean{Value: true}
			}
			if os.IsNotExist(err) {
				return object.Boolean{Value: false}
			} else if err != nil {
				vm.runtimeError("Error checking if file exists: %s", err.Error())
			}
			return object.Boolean{Value: false}
		},
	}

	fs.Map["remove"] = NativeFunction{
		Name:  "remove",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "remove", 0)
			path := args[0].(object.String).Value
			err := os.Remove(path)
			if err != nil {
				vm.runtimeError("Error removing file: %s", err.Error())
			}
			return object.Nil{}
		},
	}

	fs.Map["getwd"] = NativeFunction{
		Name:  "getwd",
		Arity: 0,
		Function: func(vm *VM, args ...object.Object) object.Object {
			dir, err := os.Getwd()
			if err != nil {
				vm.runtimeError("Error getting current working directory: %s", err.Error())
			}
			return object.NewString(dir)
		},
	}

	return fs
}
