package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
	"log"
	"slices"
)

type CallFrame struct {
	closure object.Closure
	// slots    []object.Object
	ip int
	bp int
}

func (f CallFrame) String() string {
	return fmt.Sprintf("CallFrame: { closure: %v, ip: %d, bp: %d }",
		f.closure, f.ip, f.bp)
}

const (
	STACK_SIZE = 1024
)

type VM struct {
	frames       []CallFrame
	constants    []object.Object
	stack        []object.Object
	ip           int
	openUpValues []*object.UpValue
	nativeFuncs  []object.NativeFunction
}

func NewVM(function object.Function, constants []object.Object, nativeFunctions []object.NativeFunction) *VM {
	stack := make([]object.Object, 0, STACK_SIZE)

	frames := []CallFrame{
		{
			closure: object.Closure{Function: function},
			// slots:    stack[:],
			ip: 0,
		},
	}

	return &VM{
		frames:      frames,
		constants:   constants,
		stack:       stack,
		nativeFuncs: nativeFunctions,
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

func (vm *VM) CloseUpValues(last int) {
	closedIndexes := []int{}
	// fmt.Println("Closing upvalues...")
	for i := len(vm.stack) - 1; i >= 0; i-- {
		if i < last {
			break
		}
		for _, upvalue := range vm.openUpValues {
			if upvalue.Value != &vm.stack[i] {
				continue
			}
			// Move the value from the stack to the closed field
			// and point the upvalue to the closed field
			upvalue.Closed = *upvalue.Value
			upvalue.Value = &upvalue.Closed
			// fmt.Println("Closed upvalue:", upvalue)
			closedIndexes = append(closedIndexes, i)
		}
	}
	// fmt.Println("Done closing upvalues.")

	// Remove closed upvalues from the openUpValues list
	newOpenUpValues := []*object.UpValue{}
	for _, upvalue := range vm.openUpValues {
		shouldClose := false
		for _, closedIndex := range closedIndexes {
			if upvalue.Value == &vm.stack[closedIndex] {
				shouldClose = true
				break
			}
		}
		if !shouldClose {
			newOpenUpValues = append(newOpenUpValues, upvalue)
		}
	}
	vm.openUpValues = newOpenUpValues
}

func (vm *VM) CaptureUpValues(o *object.Object) *object.UpValue {
	index := slices.IndexFunc(vm.openUpValues, func(u *object.UpValue) bool {
		return u.Value == o
	})
	if index != -1 {
		return vm.openUpValues[index]
	}
	upvalue := &object.UpValue{
		Value:  o,
		Closed: nil,
	}
	vm.openUpValues = append(vm.openUpValues, upvalue)
	return upvalue
}

func (vm *VM) Run() {

	globals := []object.Object{}

	for _, fn := range vm.nativeFuncs {
		globals = append(globals, fn)
	}

	frame := &vm.frames[0]
	debug := false
	stream := frame.closure.Function.Stream
	for frame.ip != len(stream) {
		opcode := stream[frame.ip]
		if debug {
			fmt.Print("=======================\n")
			fmt.Println("Stack: ", vm.stack)
			// // fmt.Println("Slots: ", frame.slots)
			fmt.Println("OpCode: ", opcode)
			fmt.Println("Ip: ", frame.ip)
			fmt.Println("Frame: ", vm.frames)
		}
		switch opcode {

		case op.OpConstant:
			index := stream[frame.ip+1]
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
			index := int(stream[frame.ip+1])
			frame.ip++
			if index >= len(globals) {
				globals = append(globals, vm.Peek())
			} else {
				globals[index] = vm.Peek()
			}

		case op.OpSetLocal:
			index := int(stream[frame.ip+1])
			frame.ip++
			vm.stack[frame.bp+index] = vm.Peek()

		case op.OpSetUpValue:
			index := int(stream[frame.ip+1])
			frame.ip++
			*frame.closure.UpValues[index].Value = vm.Peek()

		case op.OpLoadGlobal:
			index := int(stream[frame.ip+1])
			frame.ip++
			vm.Push(globals[index])

		case op.OpLoadLocal:
			index := int(stream[frame.ip+1])
			frame.ip++
			vm.Push(vm.stack[frame.bp+index])

		case op.OpGetUpValue:
			index := int(stream[frame.ip+1])
			frame.ip++
			vm.Push(*frame.closure.UpValues[index].Value)

		case op.OpSetProperty:
			index := int(stream[frame.ip+1])
			frame.ip++
			property := vm.constants[index]
			obj := vm.Pop()
			value := vm.Peek()
			str, ok := property.(object.String)
			if !ok {
				log.Fatal("Property name must be a string. Got: ", property.Type())
			}
			Map, ok := obj.(object.Map)
			if !ok {
				log.Fatal("Only maps have properties. Got: ", obj.Type())
			}
			Map.Map[str.Value] = value

		case op.OpGetProperty:
			index := int(stream[frame.ip+1])
			frame.ip++
			property := vm.constants[index]
			obj := vm.Pop()
			str, ok := property.(object.String)
			if !ok {
				log.Fatal("Property name must be a string. Got: ", property.Type())
			}
			Map, ok := obj.(*object.Map)
			if !ok {
				Map = obj.GetPrototype()
			}

			for Map != nil {
				value, ok := Map.Map[str.Value]
				if ok {
					vm.Push(value)
					break
				}
				Map = Map.GetPrototype()
			}
			if Map == nil {
				// log.Fatalf("Property %s not found on object of type %s\n", str.Value, obj.Type())
				vm.Push(object.Nil{})
			}

		case op.OpJump:
			jumpLength := int(stream[frame.ip+1]) - 1 // -1 because we do a frame.ip++ at the end of the loop
			frame.ip++
			frame.ip += jumpLength

		case op.OpJumpIfFalse:
			jumpLength := int(stream[frame.ip+1]) - 1 // -1 because we do a frame.ip++ at the end of the loop
			frame.ip++
			boolean := vm.Pop()
			if !boolean.GetTruthy().Value {
				frame.ip += jumpLength
			}

		case op.OpJumpIfTrue:
			jumpLength := int(stream[frame.ip+1])
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

		case op.OpClosure:
			index := stream[frame.ip+1]
			frame.ip++
			f := vm.constants[index]
			closure := object.NewClosure(f.(object.Function))
			vm.Push(closure)
			for i := 0; i < closure.Function.UpValueCount; i++ {
				isLocal := stream[frame.ip+1]
				frame.ip++
				index := int(stream[frame.ip+1])
				frame.ip++
				if isLocal == 1 {
					closure.UpValues = append(closure.UpValues, vm.CaptureUpValues(&vm.stack[frame.bp+index]))
				} else {
					closure.UpValues = append(closure.UpValues, frame.closure.UpValues[index])
				}
			}
			vm.stack[len(vm.stack)-1] = closure

		case op.OpCall:
			argCount := int(stream[frame.ip+1])
			frame.ip++

			// callee := vm.stack[frame.bp+len(vm.stack)-1-argCount]
			callee := vm.stack[len(vm.stack)-1-argCount]
			fn, ok := callee.(object.Closure)
			if !ok {
				nfn, ok := callee.(object.NativeFunction)
				if ok {
					if argCount != nfn.Arity {
						log.Fatalf("Wrong number of arguments for native function. Expected %d, got %d\n", nfn.Arity, argCount)
					}
					// Call the native function
					args := vm.stack[len(vm.stack)-argCount:]
					result := nfn.Function(args...)
					// Pop the arguments and the native function from the stack
					vm.stack = vm.stack[:len(vm.stack)-argCount-1]
					// Push the result onto the stack
					vm.Push(result)
					// Continue to the next instruction
					frame.ip++
					continue
				}
				log.Fatal("Can only call functions. Got: ", callee.Type())

			}

			if argCount != fn.Function.Arity {
				log.Fatalf("Wrong number of arguments. Expected %d, got %d\n", fn.Function.Arity, argCount)
			}

			newFrame := CallFrame{
				closure: fn,
				bp:      len(vm.stack) - 1 - argCount,
				// slots:    vm.stack[frame.bp+len(frame.slots)-argCount:], // pass the arguments and the space for local variables
				ip: 0,
			}
			vm.frames = append(vm.frames, newFrame)
			frame = &vm.frames[len(vm.frames)-1]
			stream = frame.closure.Function.Stream
			continue

		case op.OpNil:
			vm.Push(object.Nil{})

		case op.OpReturn:
			returned := vm.Pop()
			vm.frames = vm.frames[:len(vm.frames)-1]
			vm.CloseUpValues(frame.bp)
			vm.stack = vm.stack[:frame.bp]
			vm.Push(returned)
			frame = &vm.frames[len(vm.frames)-1]
			stream = frame.closure.Function.Stream

		case op.OpCloseUpValue:
			vm.CloseUpValues(len(vm.stack) - 1)
			vm.Pop()

		default:
			fmt.Println("Stack: ", vm.stack)
			log.Fatal("Unknown OpCode: ", opcode)

		}

		frame.ip++
	}
	if debug {
		fmt.Println("Ending value of stack: ", vm.stack)
	}

}
