package object

import (
	"Abbas-Askari/interpreter-v2/colors"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
)

type Function struct {
	Stream       []op.OpCode
	Arity        int
	Constants    []Object
	UpValueCount int
	LineInfo     []int
	ColumnInfo   []int
	Name         string
	ScriptName   string
	__proto__    *Map
}

var PrototypeFunction *Map = &Map{
	Map: map[string]Object{
		"prototype": Number{Value: -234}, // Just to test that prototype is being used
	},
}

func NewFunction(arity int, name string, scriptName string, stream []op.OpCode, lineInfo []int, columnInfo []int, constants []Object) Function {
	__proto__ := &Map{
		Map: map[string]Object{
			"length":    Number{Value: 0},
			"prototype": Map{Map: map[string]Object{}},
			"__proto__": PrototypeFunction,
		},
	}
	return Function{
		Stream:     stream,
		Arity:      arity,
		Constants:  constants,
		Name:       name,
		ScriptName: scriptName,
		LineInfo:   lineInfo,
		ColumnInfo: columnInfo,
		__proto__:  __proto__,
	}
}

func (b Function) String() string {
	return colors.Colorize(fmt.Sprintf("FUNC<%v>", b.Name), colors.BLUE)
}

func (b Function) Type() ObjectType {
	return FUNCTION
}

func (b Function) Add(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Sub(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Mul(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) Div(o Object) Object {
	panic("Cannot add Functions")
}

func (b Function) GetTruthy() Boolean {
	return Boolean{true}
}

func (b Function) GetPrototype() *Map {
	return b.__proto__
}
