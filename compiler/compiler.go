package compiler

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"slices"
)

type Compiler struct {
	stream     []op.OpCode
	constants  []object.Object
	globals    []string
	scope      *SymbolTable
	scopeDepth int
}

func NewCompiler() *Compiler {
	return &Compiler{
		stream:    []op.OpCode{},
		constants: []object.Object{},
		globals:   []string{},
		scope: &SymbolTable{
			Store: []Symbol{},
		},
		scopeDepth: 0,
	}
}

func (c *Compiler) AddConstant(o object.Object) int {
	c.constants = append(c.constants, o)
	return len(c.constants) - 1
}

func (c *Compiler) Declare(name string) {
	scope := LocalScope
	if c.scopeDepth == 0 {
		scope = GlobalScope
		c.Emit(op.OpSetGlobal)

		index := slices.Index(c.globals, name)
		if index != -1 {
			panic(fmt.Errorf("Error: %v is already declared", name))
		}
		c.globals = append(c.globals, name)
		c.Emit(op.OpCode(len(c.globals) - 1))
		c.Emit(op.OpPop)
		return
	}
	// No need to call set local because the expression results is already on the top of the stack. HEHE
	// c.Emit(op.OpSetLocal)
	for _, symbol := range c.scope.Store {
		if symbol.Name == name && symbol.Depth == c.scopeDepth {
			panic(fmt.Errorf("Error: %v is already declared", name))
		}
	}

	symbol := Symbol{
		Name:  name,
		Scope: scope,
		Depth: c.scopeDepth,
	}
	c.scope.NumDefs++
	c.scope.Store = append(c.scope.Store, symbol)
}

func (c *Compiler) EnterScope() {
	c.scopeDepth++
}

func (c *Compiler) ExitScope() {
	if c.scopeDepth == 0 {
		panic("Error: Broken Compiler! Cannot Exit from global scope")
	}
	c.scopeDepth--
	i := len(c.scope.Store) - 1
	for ; i >= 0; i-- {
		symbol := c.scope.Store[i]
		if symbol.Depth == c.scopeDepth {
			break
		}
		c.Emit(op.OpPop)
	}
	c.scope.Store = c.scope.Store[:i+1]
}

func (c *Compiler) GetSymbol(name token.Token, scope *SymbolTable) (*Symbol, int, error) {
	if scope == nil || c.scopeDepth == 0 {
		return nil, 0, fmt.Errorf("Unable to resolve identifier: %v", name.Literal)
		// panic(fmt.Errorf("Error: Unable to resolve identifier: %v", name.Literal))
	}
	for i := len(scope.Store) - 1; i >= 0; i-- {
		symbol := scope.Store[i]
		if symbol.Name == name.Literal {
			return &symbol, i, nil
		}
	}
	return c.GetSymbol(name, scope.Outer)
}

func (c *Compiler) GetIdentifier(name token.Token) {
	scope := c.scope

	symbol, index, err := c.GetSymbol(name, scope)
	if err == nil {
		if symbol.Depth != 0 {
			c.Emit(op.OpLoadLocal)
		}
		c.Emit(op.OpCode(index))
		return
	}

	index = slices.Index(c.globals, name.Literal)
	if index != -1 {
		c.Emit(op.OpLoadGlobal)
		c.Emit(op.OpCode(index))
		return
	}

	panic(fmt.Errorf("Error: Undeclared Identifier: %v", name.Literal))
}

func (c *Compiler) SetGlobal(name token.Token) {
	scope := c.scope

	symbol, index, err := c.GetSymbol(name, scope)
	if err == nil {
		if symbol.Depth != 0 {
			c.Emit(op.OpSetLocal)
		}
		c.Emit(op.OpCode(index))
		return
	}

	index = slices.Index(c.globals, name.Literal)
	if index != -1 {
		c.Emit(op.OpSetGlobal)
		c.Emit(op.OpCode(index))
		return
	}

	panic(fmt.Errorf("Error: Undeclared Identifier: %v", name.Literal))
}

func (c *Compiler) Emit(op op.OpCode) {
	c.stream = append(c.stream, op)
}

func (c *Compiler) Compile(statements []parser.Statement) ([]op.OpCode, []object.Object) {
	for _, statement := range statements {
		statement.Emit(c)
	}

	return c.stream, c.constants
}

func (c *Compiler) SetOpCode(i int, op op.OpCode) {
	c.stream[i] = op
}

func (c *Compiler) GetBytecodeLength() int {
	return len(c.stream)
}
