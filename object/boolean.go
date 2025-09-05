package object

type Boolean struct {
	Value bool
}

func (b Boolean) Add(o Object) Object {
	panic("Cannot add booleans")
}

func (b Boolean) Sub(o Object) Object {
	panic("Cannot add booleans")
}

func (b Boolean) Mul(o Object) Object {
	panic("Cannot add booleans")
}

func (b Boolean) Div(o Object) Object {
	panic("Cannot add booleans")
}

func (b Boolean) GetTruthy() Boolean {
	return b
}

func (b Boolean) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (b Boolean) Type() ObjectType {
	return BOOLEAN
}

func (b Boolean) GetPrototype() *Map {
	return nil
}
