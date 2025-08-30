package interfaces

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/token"
)

type ICompiler interface {
	Emit(op.OpCode)
	SetOpCode(int, op.OpCode)
	GetOpCode(int) op.OpCode
	AddConstant(object.Object) int
	Declare(string)
	GetIdentifier(name token.Token)
	SetGlobal(name token.Token)
	EnterScope()
	ExitScope()
	GetBytecodeLength() int
	EnterTarget()
	ExitTarget() int
}
