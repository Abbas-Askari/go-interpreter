package parser

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

func Decompile(stream []op.OpCode, constants []object.Object) {
	fmt.Println("----------DeCompiler----------")
	i := 0
	for i < len(stream) {
		current := stream[i]

		if current == op.OpConstant {
			fmt.Printf("%04d %v %d\n", i, current, stream[i+1])
			i += 2
			continue
		}

		fmt.Printf("%04d %v\n", i, current)
		i++
	}
	fmt.Println("------------------------------")
}
