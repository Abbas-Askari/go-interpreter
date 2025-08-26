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
		// fmt.Println("Stack: ", vm.stack)
		switch opcode {

		case op.OpConstant:
			index := vm.bytecode[vm.ip+1]
			vm.ip++
			constant := vm.constants[index]
			vm.Push(constant)

		case op.OpAdd:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Add(right))

		case op.OpSub:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Sub(right))

		case op.OpMul:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(left.Mul(right))

		case op.OpMod:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.Mod(left, right))

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

		case op.OpSetLocal:
			index := int(vm.bytecode[vm.ip+1])
			vm.ip++
			vm.stack[index] = vm.Peek()

		case op.OpLoadGlobal:
			index := int(vm.bytecode[vm.ip+1])
			vm.ip++
			vm.Push(globals[index])

		case op.OpLoadLocal:
			index := int(vm.bytecode[vm.ip+1])
			vm.ip++
			vm.Push(vm.stack[index])

		case op.OpJump:
			jumpLength := int(vm.bytecode[vm.ip+1]) - 1 // -1 because we do a vm.ip++ at the end of the loop
			vm.ip++
			vm.ip += jumpLength

		case op.OpJumpIfFalse:
			jumpLength := int(vm.bytecode[vm.ip+1]) - 1 // -1 because we do a vm.ip++ at the end of the loop
			vm.ip++
			boolean := vm.Pop()
			if !boolean.GetTruthy().Value {
				vm.ip += jumpLength
			}

		case op.OpJumpIfTrue:
			jumpLength := int(vm.bytecode[vm.ip+1])
			vm.ip++
			boolean := vm.Pop()
			if boolean.GetTruthy().Value {
				vm.ip += jumpLength
			}

		case op.OpEqual:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.Equal(right, left))
		case op.OpNotEqual:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.NotEqual(right, left))
		case op.OpGreaterThan:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.Greater(right, left))
		case op.OpLessThan:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.Less(right, left))
		case op.OpGreaterEqual:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.GreaterOrEqual(right, left))
		case op.OpLessEqual:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.LessOrEqual(right, left))
		case op.OpNeg:
			val := vm.Pop()
			vm.Push(object.Neg(val))
		case op.OpNot:
			val := vm.Pop()
			vm.Push(object.Not(val))
		case op.OpAnd:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.And(right, left))
		case op.OpOr:
			right, left := vm.Pop(), vm.Pop()
			vm.Push(object.Or(right, left))

		case op.OpTrue:
			vm.Push(object.Boolean{Value: true})
		case op.OpFalse:
			vm.Push(object.Boolean{Value: false})

		default:
			log.Fatal("Unknown OpCode: ", opcode)

		}

		// fmt.Println("Stack: ", vm.stack)

		vm.ip++
	}
}
