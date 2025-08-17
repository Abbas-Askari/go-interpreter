package interfaces

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
)

type ICompiler interface {
	Emit(op.OpCode)
	AddConstant(object.Object) int
}
