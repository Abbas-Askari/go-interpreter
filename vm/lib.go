package vm

import "Abbas-Askari/interpreter-v2/object"

func GetLibraryMaps() map[string]*object.Map {
	return map[string]*object.Map{
		"fs":    getFileSystem(),
		"http":  getHttp(),
		"json":  getJson(),
		"os":    getOs(),
		"async": getAsync(),
	}
}

func (vm *VM) assertArgumentToType(arg object.Object, expectedType object.ObjectType, fnName string, argIndex int) {
	if arg.Type() != expectedType {
		vm.runtimeError("In function '%s', argument %d must be of type %s, got %s instead", fnName, argIndex, expectedType, arg.Type())
	}
}
