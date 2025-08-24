package parser

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

func Decompile(stream []op.OpCode, constants []object.Object) {
	fmt.Println("----------DeCompiler----------")
	i := 0
	operandCount := map[op.OpCode]int{
		op.OpConstant:   1,
		op.OpSetGlobal:  1,
		op.OpLoadGlobal: 1,
		op.OpSetLocal:   1,
		op.OpLoadLocal:  1,
	}
	for i < len(stream) {
		current := stream[i]

		if _, ok := operandCount[current]; ok {
			fmt.Printf("%04d %v %d\n", i, current, stream[i+1])
			i += 2
			continue
		}

		fmt.Printf("%04d %v\n", i, current)
		i++
	}
	fmt.Println("------------------------------")
}
