package compiler

type SymbolScope int

const (
	GlobalScope SymbolScope = iota
	LocalScope
	FreeScope // for closures
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Depth int // position in globals array OR stack slot
}

type SymbolTable struct {
	Outer   *SymbolTable // enclosing scope
	Store   []Symbol     // name → symbol
	NumDefs int          // how many locals defined in this scope
}
