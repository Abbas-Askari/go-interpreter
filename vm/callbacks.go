package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"log"
)

type QueueElement struct {
	function object.Closure
	args     []object.Object
}

func (vm *VM) RegisterEvent(closure object.Closure) int {
	vm.eventQueue = append(vm.eventQueue, closure)
	return len(vm.eventQueue) - 1
}

func (vm *VM) FireEvent(index int, args ...object.Object) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	if index < 0 || index >= len(vm.eventQueue) {
		panic("Invalid event index")
	}
	vm.callbackQueue = append(vm.callbackQueue, QueueElement{
		function: vm.eventQueue[index],
		args:     args,
	})
	vm.cond.Signal()
}

func (vm *VM) HadPendingEvents() bool {
	return len(vm.eventQueue) > 0
}

func (vm *VM) HasPendingCallbacks() bool {
	return len(vm.callbackQueue) > 0
}

func (vm *VM) DetachEvent(index int) {
	if index < 0 || index >= len(vm.eventQueue) {
		panic("Invalid event index")
	}
	vm.eventQueue = append(vm.eventQueue[:index], vm.eventQueue[index+1:]...)
}

func (vm *VM) ExecuteNextCallback() {
	vm.mu.Lock()
	if len(vm.callbackQueue) == 0 {
		vm.cond.Wait()
	}
	callback := vm.callbackQueue[0]
	vm.callbackQueue = vm.callbackQueue[1:]
	vm.mu.Unlock()
	vm.Push(callback.function)
	for _, arg := range callback.args {
		vm.Push(arg)
	}
	fn := callback.function
	argCount := len(callback.args)

	if argCount != fn.Function.Arity {
		log.Fatalf("Expected %d arguments but got %d\n", fn.Function.Arity, argCount)
	}

	// caller frame

	callerFrame := CallFrame{
		closure: object.NewClosure(object.Function{
			Name:       "<callback>",
			ScriptName: "<callback>",
			Stream: []op.OpCode{
				op.OpCall,
				op.OpCode(argCount),
				op.OpPop,
			},
			LineInfo:   []int{0, 0, 0},
			ColumnInfo: []int{0, 0, 0},
		}),
		bp: len(vm.stack) - 1 - argCount,
		ip: 0,
	}

	// newFrame := CallFrame{
	// 	closure: callback.function,
	// 	bp:      len(vm.stack) - 1 - argCount,
	// 	ip:      0,
	// }
	vm.frames = append(vm.frames, callerFrame)

	// if fn.This != nil {
	// 	vm.stack[len(vm.stack)-1-argCount] = *fn.This
	// } else {
	// 	vm.stack[len(vm.stack)-1-argCount] = object.Nil{}
	// }
}
