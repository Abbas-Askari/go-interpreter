package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
	"log"
)

type CallFrame struct {
	function object.Function
	// slots    []object.Object
	ip int
	bp int
}

func (f CallFrame) String() string {
	return fmt.Sprintf("CallFrame: { function: %v, ip: %d, bp: %d }",
		f.function, f.ip, f.bp)
}

const (
	STACK_SIZE = 1024
)

type VM struct {
	frames    []CallFrame
	constants []object.Object
	stack     []object.Object
	ip        int
}

func NewVM(function object.Function, constants []object.Object) *VM {
	stack := make([]object.Object, 0, STACK_SIZE)

	frames := []CallFrame{
		{
			function: function,
			// slots:    stack[:],
			ip: 0,
		},
	}

	return &VM{
		frames:    frames,
		constants: constants,
		stack:     stack,
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

	frame := &vm.frames[0]
	debug := false

	for frame.ip != len(frame.function.Stream) {
		opcode := frame.function.Stream[frame.ip]
		if debug {
			fmt.Println("Stack: ", vm.stack)
			// // fmt.Println("Slots: ", frame.slots)
			// fmt.Println("OpCode: ", opcode)
			// fmt.Println("Ip: ", frame.ip)
			// fmt.Println("Frame: ", vm.frames)
		}
		switch opcode {

		case op.OpConstant:
			index := frame.function.Stream[frame.ip+1]
			frame.ip++
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
			index := int(frame.function.Stream[frame.ip+1])
			frame.ip++
			if index >= len(globals) {
				globals = append(globals, vm.Peek())
			} else {
				globals[index] = vm.Peek()
			}

		case op.OpSetLocal:
			index := int(frame.function.Stream[frame.ip+1])
			frame.ip++
			vm.stack[frame.bp+index] = vm.Peek()

		case op.OpLoadGlobal:
			index := int(frame.function.Stream[frame.ip+1])
			frame.ip++
			vm.Push(globals[index])

		case op.OpLoadLocal:
			index := int(frame.function.Stream[frame.ip+1])
			frame.ip++
			vm.Push(vm.stack[frame.bp+index])

		case op.OpJump:
			jumpLength := int(frame.function.Stream[frame.ip+1]) - 1 // -1 because we do a frame.ip++ at the end of the loop
			frame.ip++
			frame.ip += jumpLength

		case op.OpJumpIfFalse:
			jumpLength := int(frame.function.Stream[frame.ip+1]) - 1 // -1 because we do a frame.ip++ at the end of the loop
			frame.ip++
			boolean := vm.Pop()
			if !boolean.GetTruthy().Value {
				frame.ip += jumpLength
			}

		case op.OpJumpIfTrue:
			jumpLength := int(frame.function.Stream[frame.ip+1])
			frame.ip++
			boolean := vm.Pop()
			if boolean.GetTruthy().Value {
				frame.ip += jumpLength
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

		case op.OpBreak:
			log.Fatal("Break statement not in loop")
		case op.OpContinue:
			log.Fatal("Continue statement not in loop")

		case op.OpCall:
			argCount := int(frame.function.Stream[frame.ip+1])
			frame.ip++

			// callee := vm.stack[frame.bp+len(vm.stack)-1-argCount]
			callee := vm.Pop()
			fn, ok := callee.(object.Function)
			if !ok {
				log.Fatal("Can only call functions. Got: ", callee.Type())
			}

			newFrame := CallFrame{
				function: fn,
				bp:       len(vm.stack) - argCount,
				// slots:    vm.stack[frame.bp+len(frame.slots)-argCount:], // pass the arguments and the space for local variables
				ip: 0,
			}
			vm.frames = append(vm.frames, newFrame)
			frame = &vm.frames[len(vm.frames)-1]
			continue

		case op.OpNil:
			vm.Push(object.Nil{})

		case op.OpReturn:
			returned := vm.Pop()
			vm.frames = vm.frames[:len(vm.frames)-1]
			vm.stack = vm.stack[:frame.bp]
			vm.Push(returned)
			frame = &vm.frames[len(vm.frames)-1]

		default:
			log.Fatal("Unknown OpCode: ", opcode)

		}

		frame.ip++
	}
	if debug {
		fmt.Println("Ending value of stack: ", vm.stack)
	}

}
