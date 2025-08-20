package interfaces

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
)

type ICompiler interface {
	Emit(op.OpCode)
	AddConstant(object.Object) int
	AddGlobal(string) int
	GetGlobal(name token.Token)
	SetGlobal(name token.Token)
}
