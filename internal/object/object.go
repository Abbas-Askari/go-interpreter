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
	BUFFER   = "BUFFER"
)

type Object interface {
	Add(Object) (Object, error)
	Sub(Object) (Object, error)
	Mul(Object) (Object, error)
	Div(Object) (Object, error)
	GetTruthy() Boolean
	String() string
	Type() ObjectType
	GetPrototype() *Map
}
