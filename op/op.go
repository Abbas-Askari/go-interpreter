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
	OpMinus
	OpBang
	OpJump
	OpJumpNotTruthy
	OpNull
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
	}
	return opNames[o]
}
