package parser

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

func Decompile(function object.Function) {
	constants := function.Constants
	fmt.Println("----------DeCompiler----------")
	fmt.Println("Stream:", function.Stream)
	fmt.Println("Constants:", constants)
	i := 0
	operandCount := map[op.OpCode]int{
		op.OpConstant:    1,
		op.OpSetGlobal:   1,
		op.OpLoadGlobal:  1,
		op.OpSetLocal:    1,
		op.OpLoadLocal:   1,
		op.OpGetUpValue:  1,
		op.OpSetUpValue:  1,
		op.OpJumpIfFalse: 1,
		op.OpJumpIfTrue:  1,
		op.OpJump:        1,
		op.OpCall:        2,
		op.OpSetProperty: 1,
		op.OpGetProperty: 1,
		op.OpArray:       1,
	}
	for i < len(function.Stream) {
		current := function.Stream[i]

		if current == op.OpClosure {
			functionIndex := function.Stream[i+1]
			closureFunction, ok := constants[functionIndex].(object.Function)
			if !ok {
				panic("Expected function")
			}
			upValuesCount := closureFunction.UpValueCount
			fmt.Printf("%04d %v %d\n", i, current, functionIndex)
			i += 2
			for j := 0; j < upValuesCount; j++ {
				isLocal := function.Stream[i]
				index := function.Stream[i+1]
				str := "????"
				if isLocal == 1 {
					str = "local"
				} else if isLocal == 0 {
					str = "upvalue"
				}
				fmt.Printf("%04d      |-- %s %d\n", i, str, index)
				i += 2
			}
			continue
		}

		if operandCount, ok := operandCount[current]; ok {
			fmt.Printf("%04d %v", i, current)
			for j := 0; j < operandCount; j++ {
				fmt.Printf(" %d", function.Stream[i+1+j])
			}
			fmt.Println()
			i += operandCount + 1
			continue
		}

		fmt.Printf("%04d %v\n", i, current)
		i++
	}
	fmt.Println("------------------------------")
}
