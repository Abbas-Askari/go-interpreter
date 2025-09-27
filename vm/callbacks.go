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

func (vm *VM) RegisterEvent() {
	vm.eventMu.Lock()
	defer vm.eventMu.Unlock()
	vm.pendingEvents++
}

func (vm *VM) DetachEvent() {
	vm.eventMu.Lock()
	defer vm.eventMu.Unlock()
	vm.pendingEvents--
}

func (vm *VM) FireEvent(function object.Closure, args ...object.Object) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.callbackQueue = append(vm.callbackQueue, QueueElement{
		function: function,
		args:     args,
	})
	vm.cond.Signal()
}

func (vm *VM) HasPendingEvents() bool {
	vm.eventMu.Lock()
	defer vm.eventMu.Unlock()
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.pendingEvents > 0 || len(vm.callbackQueue) > 0
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
		closure: object.NewClosure(object.NewFunction(
			0,
			"<callback>",
			"<callback>",
			[]op.OpCode{
				op.OpCall,
				op.OpCode(argCount),
				op.OpCode(0),
				op.OpPop,
			},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
			[]object.Object{},
		)),
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
