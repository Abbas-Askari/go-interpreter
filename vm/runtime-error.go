package vm

import (
	"Abbas-Askari/interpreter-v2/colors"
	"fmt"
	"os"
)

func (vm *VM) runtimeError(format string, args ...interface{}) {
	errMessage := fmt.Sprintf(format+"\n", args...)
	for len(vm.frames) > 0 {
		frame := vm.frames[len(vm.frames)-1]
		ip := frame.ip
		line := frame.closure.Function.LineInfo[ip-1]
		column := frame.closure.Function.ColumnInfo[ip-1]
		errMessage = fmt.Sprintf("%s\t[line %3d:%3d of %s] at %s\n", errMessage, line, column, frame.closure.Function.ScriptName, frame.closure.Function.Name)
		vm.frames = vm.frames[:len(vm.frames)-1]
	}
	fmt.Println(colors.Colorize("Runtime Error:", colors.RED))
	fmt.Print(errMessage)
	os.Exit(1)
}
