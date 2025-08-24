package op

type OpCode int

const (
	OpAdd OpCode = iota
	OpConstant
	OpSub
	OpMul
	OpDiv
	OpPop
	OpTrue
	OpFalse
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpLessThan
	OpGreaterEqual
	OpLessEqual
	OpMinus
	OpBang
	OpJump
	OpJumpNotTruthy
	OpNull
	OpPrint
	OpSetGlobal
	OpLoadGlobal
	OpLoadLocal
	OpSetLocal
	OpJumpIfFalse
	OpJumpIfTrue
)

func (o OpCode) String() string {
	opNames := map[OpCode]string{
		OpConstant:      "OpConstant",
		OpAdd:           "OpAdd",
		OpSub:           "OpSub",
		OpMul:           "OpMul",
		OpDiv:           "OpDiv",
		OpPop:           "OpPop",
		OpTrue:          "OpTrue",
		OpFalse:         "OpFalse",
		OpEqual:         "OpEqual",
		OpNotEqual:      "OpNotEqual",
		OpGreaterThan:   "OpGreaterThan",
		OpMinus:         "OpMinus",
		OpBang:          "OpBang",
		OpJump:          "OpJump",
		OpJumpNotTruthy: "OpJumpNotTruthy",
		OpNull:          "OpNull",
		OpPrint:         "OpPrint",
		OpSetGlobal:     "OpSetGlobal",
		OpLoadGlobal:    "OpLoadGlobal",
		OpLoadLocal:     "OpLoadLocal",
		OpSetLocal:      "OpSetLocal",
		OpJumpIfFalse:   "OpJumpIfFalse",
		OpJumpIfTrue:    "OpJumpIfTrue",
	}
	return opNames[o]
}
