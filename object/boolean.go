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

func (b Boolean) GetTruthy(o Object) Boolean {
	return b
}
