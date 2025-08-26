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
	OpAnd
	OpOr
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
	OpNot
	OpNeg
	OpBreak
	OpContinue
)

func (o OpCode) String() string {
	opNames := map[OpCode]string{
		OpAdd:           "OpAdd",
		OpConstant:      "OpConstant",
		OpSub:           "OpSub",
		OpMul:           "OpMul",
		OpDiv:           "OpDiv",
		OpPop:           "OpPop",
		OpTrue:          "OpTrue",
		OpFalse:         "OpFalse",
		OpEqual:         "OpEqual",
		OpNotEqual:      "OpNotEqual",
		OpGreaterThan:   "OpGreaterThan",
		OpLessThan:      "OpLessThan",
		OpGreaterEqual:  "OpGreaterEqual",
		OpLessEqual:     "OpLessEqual",
		OpAnd:           "OpAnd",
		OpOr:            "OpOr",
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
		OpNot:           "OpNot",
		OpNeg:           "OpNeg",
		OpBreak:         "OpBreak",
		OpContinue:      "OpContinue",
	}
	str, ok := opNames[o]
	if !ok {
		panic("Unknown OpCode")
	}
	return str
}
