package vm

import (
	"Abbas-Askari/interpreter-v2/internal/colors"
	"fmt"
	"os"
)

func (vm *VM) runtimeError(format string, args ...interface{}) {
	errMessage := fmt.Sprintf(format+"\n", args...)
	for len(vm.frames) > 0 {
		frame := vm.frames[len(vm.frames)-1]
		ip := frame.ip
		line := frame.closure.Function.LineInfo[ip-1]
		errMessage = fmt.Sprintf("%s\t[line %3d of %s:%d] at %s\n", errMessage, line, frame.closure.Function.ScriptName, line, frame.closure.Function.Name)
		vm.frames = vm.frames[:len(vm.frames)-1]
	}
	fmt.Println(colors.Colorize("Runtime Error:", colors.RED))
	fmt.Print(errMessage)
	os.Exit(1)
}
