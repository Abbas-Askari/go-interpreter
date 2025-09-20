package object

type ObjectType string

const (
	NUMBER   = "NUMBER"
	BOOLEAN  = "BOOLEAN"
	STRING   = "STRING"
	NIL      = "NIL"
	FUNCTION = "FUNCTION"
	CLOSURE  = "CLOSURE"
	UPVALUE  = "UPVALUE"
	MAP      = "MAP"
	ARRAY    = "ARRAY"
)

type Object interface {
	Add(Object) Object
	Sub(Object) Object
	Mul(Object) Object
	Div(Object) Object
	GetTruthy() Boolean
	String() string
	Type() ObjectType
	GetPrototype() *Map
}
