package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
	"log"
)

type VM struct {
	bytecode  []op.OpCode
	constants []object.Object
	stack     []object.Object
	ip        int
}

func NewVM(bytecode []op.OpCode, constants []object.Object) *VM {
	return &VM{
		bytecode:  bytecode,
		constants: constants,
		stack:     []object.Object{},
		ip:        0,
	}
}

func (vm *VM) Pop() object.Object {
	top := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return top
}

func (vm *VM) Peek() object.Object {
	top := vm.stack[len(vm.stack)-1]
	return top
}

func (vm *VM) Push(o object.Object) {
	vm.stack = append(vm.stack, o)
}

func (vm *VM) Run() {

	globals := []object.Object{}

	for vm.ip != len(vm.bytecode) {
		opcode := vm.bytecode[vm.ip]
		switch opcode {

		case op.OpConstant:
			index := vm.bytecode[vm.ip+1]
			vm.ip++
			constant := vm.constants[index]
			vm.Push(constant)

		case op.OpAdd:
			left, right := vm.Pop(), vm.Pop()
			vm.Push(left.Add(right))

		case op.OpSub:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Sub(right))

		case op.OpMul:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Mul(right))

		case op.OpDiv:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Div(right))

		case op.OpPop:
			vm.Pop()

		case op.OpPrint:
			o := vm.Pop()
			fmt.Println(o)

		case op.OpSetGlobal:
			index := int(vm.bytecode[vm.ip+1])
			vm.ip++
			if index >= len(globals) {
				globals = append(globals, vm.Peek())
			} else {
				globals[index] = vm.Peek()
			}

		case op.OpLoadGlobal:
			index := int(vm.bytecode[vm.ip+1])
			vm.ip++
			vm.Push(globals[index])

		default:
			log.Fatal("Unknown OpCode: ", opcode)

		}

		// fmt.Println(globals)

		vm.ip++
	}
}
