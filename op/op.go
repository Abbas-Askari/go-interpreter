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
	OpMod
	OpCall
	OpNil
	OpReturn
	OpClosure
	OpGetUpValue
	OpSetUpValue
	OpCloseUpValue
	OpGetProperty
	OpSetProperty
	OpGetIndex
	OpSetIndex
	OpArray
)

func (o OpCode) String() string {
	opNames := map[OpCode]string{
		OpAdd:           "Add",
		OpConstant:      "Constant",
		OpSub:           "Sub",
		OpMul:           "Mul",
		OpDiv:           "Div",
		OpPop:           "Pop",
		OpTrue:          "True",
		OpFalse:         "False",
		OpEqual:         "Equal",
		OpNotEqual:      "NotEqual",
		OpGreaterThan:   "GreaterThan",
		OpLessThan:      "LessThan",
		OpGreaterEqual:  "GreaterEqual",
		OpLessEqual:     "LessEqual",
		OpAnd:           "And",
		OpOr:            "Or",
		OpMinus:         "Minus",
		OpBang:          "Bang",
		OpJump:          "Jump",
		OpJumpNotTruthy: "JumpNotTruthy",
		OpNull:          "Null",
		OpPrint:         "Print",
		OpSetGlobal:     "SetGlobal",
		OpLoadGlobal:    "LoadGlobal",
		OpLoadLocal:     "LoadLocal",
		OpSetLocal:      "SetLocal",
		OpJumpIfFalse:   "JumpIfFalse",
		OpJumpIfTrue:    "JumpIfTrue",
		OpNot:           "Not",
		OpNeg:           "Neg",
		OpBreak:         "Break",
		OpContinue:      "Continue",
		OpMod:           "Mod",
		OpCall:          "Call",
		OpNil:           "Nil",
		OpReturn:        "Return",
		OpClosure:       "Closure",
		OpGetUpValue:    "GetUpValue",
		OpSetUpValue:    "SetUpValue",
		OpCloseUpValue:  "CloseUpValue",
		OpGetProperty:   "GetProperty",
		OpSetProperty:   "SetProperty",
		OpGetIndex:      "GetIndex",
		OpSetIndex:      "SetIndex",
		OpArray:         "Array",
	}
	str, ok := opNames[o]
	if !ok {
		panic("Unknown OpCode")
	}
	return str
}
