package parser

import (
	"Abbas-Askari/interpreter-v2/interfaces"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

type Statement interface {
	Emit(p interfaces.ICompiler)
	String() string
	Type() DeclarationType
}

type PrintStatement struct {
	expression Expression
}

func (ps PrintStatement) Emit(c interfaces.ICompiler) {
	ps.expression.Emit(c)
	c.Emit(op.OpPrint)
}

func (ps PrintStatement) String() string {
	return fmt.Sprintf("PRINT: %v\n", ps.expression)
}

func (ps PrintStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type ExpressionStatement struct {
	expression Expression
}

func (e ExpressionStatement) Emit(c interfaces.ICompiler) {
	e.expression.Emit(c)
	c.Emit(op.OpPop)
}

func (ex ExpressionStatement) String() string {
	return fmt.Sprintf("Expression: %v\n", ex.expression)
}

func (ex ExpressionStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type BlockStatement struct {
	declarations []Declaration
}

func (b *BlockStatement) Emit(c interfaces.ICompiler) {
	c.EnterScope()
	for _, statement := range b.declarations {
		statement.Emit(c)
	}
	c.ExitScope()
}

func (b BlockStatement) String() string {
	return fmt.Sprintf("Block: {\n%v}\n", b.declarations)
}

func (b BlockStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type IfStatement struct {
	condition     Expression
	thenStatement Statement
	elseStatement *Statement
}

func (b *IfStatement) Emit(c interfaces.ICompiler) {
	b.condition.Emit(c)

	c.Emit(op.OpJumpIfFalse)
	c.Emit(op.OpCode(0))

	jumpLengthIndex := c.GetBytecodeLength() - 1

	b.thenStatement.Emit(c)
	jumpElseBlockLengthIndex := 0
	if b.elseStatement != nil {
		c.Emit(op.OpJump)
		c.Emit(op.OpCode(0))
		jumpElseBlockLengthIndex = c.GetBytecodeLength() - 1
	}

	jumpLength := c.GetBytecodeLength() - jumpLengthIndex
	c.SetOpCode(jumpLengthIndex, op.OpCode(jumpLength))

	if b.elseStatement != nil {
		(*b.elseStatement).Emit(c)
		jumpLength := c.GetBytecodeLength() - jumpElseBlockLengthIndex
		c.SetOpCode(jumpElseBlockLengthIndex, op.OpCode(jumpLength))
	}

}

func (b IfStatement) String() string {
	if b.elseStatement != nil {
		return fmt.Sprintf("If: (%v) {\n%v} else {\n%v}\n", b.condition, b.thenStatement, *b.elseStatement)
	}
	return fmt.Sprintf("If: (%v) {\n%v} else {\n%v}\n", b.condition, b.thenStatement, b.elseStatement)
}

func (b IfStatement) Type() DeclarationType {
	return StatementDeclarationType
}

const (
	BREAK_JUMP_LENGTH    op.OpCode = -1
	CONTINUE_JUMP_LENGTH op.OpCode = -2
)

type ForStatement struct {
	initialization Declaration
	condition      Expression
	advancement    Expression
	body           Statement
}

func (f *ForStatement) Emit(c interfaces.ICompiler) {
	fmt.Println("Entering for Statement")
	c.EnterScope()
	if f.initialization != nil {
		f.initialization.Emit(c)
	}

	startIndex := c.GetBytecodeLength()

	if f.condition != nil {
		f.condition.Emit(c)
		c.Emit(op.OpJumpIfFalse)
		c.Emit(op.OpCode(0))
	}

	jumpLengthIndex := c.GetBytecodeLength() - 1

	f.body.Emit(c)
	advancementIndex := c.GetBytecodeLength()
	if f.advancement != nil {
		f.advancement.Emit(c)
		c.Emit(op.OpPop)
	}
	c.Emit(op.OpJump)
	c.Emit(op.OpCode(startIndex - c.GetBytecodeLength()))

	loopEndTarget := c.GetBytecodeLength()

	if f.condition != nil {
		jumpLength := loopEndTarget - jumpLengthIndex
		c.SetOpCode(jumpLengthIndex, op.OpCode(jumpLength))
	}

	// Hacky fix not seen anywhere for breaks and continue.
	for i := startIndex; i < loopEndTarget; i++ {
		if c.GetOpCode(i) == op.OpBreak && c.GetOpCode(i+1) == op.OpBreak {
			c.SetOpCode(i, op.OpJump)
			c.SetOpCode(i+1, op.OpCode(loopEndTarget-(i+1)))
		}
		if c.GetOpCode(i) == op.OpContinue && c.GetOpCode(i+1) == op.OpContinue {
			c.SetOpCode(i, op.OpJump)
			c.SetOpCode(i+1, op.OpCode(advancementIndex-(i+1)))
			// panic("LOL")
		}
	}

	c.ExitScope()
	fmt.Println("Exiting for Statement")
}

func (f ForStatement) String() string {
	return fmt.Sprintf("for: (%v;%v;%v) {\n%v}\n", f.initialization, f.condition, f.advancement, f.body)
}

func (f ForStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type BreakStatement struct{}

func (b *BreakStatement) Emit(c interfaces.ICompiler) {
	c.Emit(op.OpBreak)
	c.Emit(op.OpBreak)
}

func (b BreakStatement) String() string {
	return fmt.Sprintf("Break\n")
}

func (b BreakStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type ContinueStatement struct{}

func (b *ContinueStatement) Emit(c interfaces.ICompiler) {
	c.Emit(op.OpContinue)
	c.Emit(op.OpContinue)
}

func (b ContinueStatement) String() string {
	return fmt.Sprintf("Continue\n")
}

func (b ContinueStatement) Type() DeclarationType {
	return StatementDeclarationType
}

type ReturnStatement struct {
	exp Expression
}

func (r *ReturnStatement) Emit(c interfaces.ICompiler) {
	if r.exp != nil {
		r.exp.Emit(c)
	} else {
		c.Emit(op.OpNil)
	}
	c.Emit(op.OpReturn)
}

func (r ReturnStatement) String() string {
	if r.exp != nil {
		return fmt.Sprintf("Return: %v\n", r.exp)
	}
	return fmt.Sprintf("Return: <no expression>\n")
}

func (r ReturnStatement) Type() DeclarationType {
	return StatementDeclarationType
}
