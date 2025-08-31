package compiler

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"Abbas-Askari/interpreter-v2/parser"
	"Abbas-Askari/interpreter-v2/token"
	"fmt"
	"log"
	"slices"
)

type Compiler struct {
	constants []object.Object
	globals   []string

	target *Target
}

type TargetType int

const (
	FUNCTION_TARGET TargetType = iota
	SCRIPT_TARGET
)

type Target struct {
	function   object.Function
	targetType TargetType

	scope      *SymbolTable
	scopeDepth int

	outer *Target
}

func NewTarget(outer *Target) *Target {
	t := SCRIPT_TARGET
	if outer != nil {
		t = FUNCTION_TARGET
	}

	store := []Symbol{}
	if outer != nil {
		// store = append(store, Symbol{Name: "this", Scope: LocalScope, Depth: 0})
	}

	return &Target{
		function: object.Function{
			Stream: []op.OpCode{},
		},
		targetType: t,
		scope: &SymbolTable{
			Store: store,
		},
		scopeDepth: 0,
		outer:      outer,
	}
}

func (c *Compiler) EnterTarget() {
	fmt.Println("Enter from ", c.target.function.Stream)
	t := NewTarget(c.target)
	// c.Declare("this")
	c.target = t
}

func (c *Compiler) ExitTarget(arity int) int {
	f := c.target.function
	f.Arity = arity
	fmt.Println("Exited Scope: ", c.target.scope)
	c.target = c.target.outer
	c.constants = append(c.constants, f)
	fmt.Println("Exited: ", f)
	fmt.Println("Back to ", c.target.function.Stream)
	return len(c.constants) - 1
}

func NewCompiler() *Compiler {
	return &Compiler{
		constants: []object.Object{},
		globals:   []string{},
		target:    NewTarget(nil),
	}
}

func (c *Compiler) AddConstant(o object.Object) int {
	c.constants = append(c.constants, o)
	return len(c.constants) - 1
}

func (c *Compiler) Declare(name string) {
	scope := LocalScope
	if c.target.scopeDepth == 0 && c.target.targetType == SCRIPT_TARGET {
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
	for _, symbol := range c.target.scope.Store {
		if symbol.Name == name && symbol.Depth == c.target.scopeDepth {
			panic(fmt.Errorf("Error: %v is already declared", name))
		}
	}

	symbol := Symbol{
		Name:  name,
		Scope: scope,
		Depth: c.target.scopeDepth,
	}
	c.target.scope.NumDefs++
	c.target.scope.Store = append(c.target.scope.Store, symbol)
}

func (c *Compiler) EnterScope() {
	c.target.scopeDepth++
}

func (c *Compiler) ExitScope() {
	if c.target.scopeDepth == 0 {
		panic("Error: Broken Compiler! Cannot Exit from global scope")
	}
	c.target.scopeDepth--
	i := len(c.target.scope.Store) - 1
	for ; i >= 0; i-- {
		symbol := c.target.scope.Store[i]
		if symbol.Depth <= c.target.scopeDepth {
			break
		}
		c.Emit(op.OpPop)
	}
	c.target.scope.Store = c.target.scope.Store[:i+1]
}

func (c *Compiler) GetSymbol(name token.Token, scope *SymbolTable) (*Symbol, int, error) {
	if scope == nil || c.target.scopeDepth == 0 {
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
	if true {
		fmt.Println(name, c.globals, c.target.scope, c.target.scopeDepth)
		fmt.Println(name, c.target.function.Stream)

		defer fmt.Println(c.target.function.Stream)
	}

	scope := c.target.scope

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
	log.Fatalf("Error: Undeclared Identifier: %v", name.Literal)
	panic(fmt.Errorf("Error: Undeclared Identifier: %v", name.Literal))
}

func (c *Compiler) SetGlobal(name token.Token) {
	scope := c.target.scope

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
	// fmt.Println("Emit", op)
	c.target.function.Stream = append(c.target.function.Stream, op)
}

func (c *Compiler) Compile(statements []parser.Declaration) (object.Function, []object.Object) {
	for _, statement := range statements {
		statement.Emit(c)
	}

	return c.target.function, c.constants
}

func (c *Compiler) SetOpCode(i int, op op.OpCode) {
	c.target.function.Stream[i] = op
}

func (c *Compiler) GetOpCode(i int) op.OpCode {
	return c.target.function.Stream[i]
}

func (c *Compiler) GetBytecodeLength() int {
	return len(c.target.function.Stream)
}
