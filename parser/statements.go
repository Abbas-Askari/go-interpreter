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
	statements []Statement
}

func (b *BlockStatement) Emit(c interfaces.ICompiler) {
	c.EnterScope()
	for _, statement := range b.statements {
		statement.Emit(c)
	}
	c.ExitScope()
}

func (b BlockStatement) String() string {
	return fmt.Sprintf("Block: {\n%v}\n", b.statements)
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

type ForStatement struct {
	initialization Declaration
	condition      Expression
	advancement    Expression
	body           Statement
}

func (f *ForStatement) Emit(c interfaces.ICompiler) {
	if f.initialization != nil {
		c.EnterScope()
		defer c.ExitScope()
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
	if f.advancement != nil {
		f.advancement.Emit(c)
	}
	c.Emit(op.OpJump)
	c.Emit(op.OpCode(startIndex - c.GetBytecodeLength()))

	if f.condition != nil {
		jumpLength := c.GetBytecodeLength() - jumpLengthIndex
		c.SetOpCode(jumpLengthIndex, op.OpCode(jumpLength))
	}
}

func (f ForStatement) String() string {
	return fmt.Sprintf("for: (%v;%v;%v) {\n%v}\n", f.initialization, f.condition, f.advancement, f.body)
}

func (f ForStatement) Type() DeclarationType {
	return StatementDeclarationType
}
