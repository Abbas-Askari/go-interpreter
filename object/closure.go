package object

type Closure struct {
	Function Function
}

func (b Closure) String() string {
	return CLOSURE
}

func (b Closure) Type() ObjectType {
	return CLOSURE
}

func (b Closure) Add(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Sub(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Mul(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) Div(o Object) Object {
	panic("Cannot add Closures")
}

func (b Closure) GetTruthy() Boolean {
	return Boolean{true}
}
