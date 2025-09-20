package object

type Nil struct{}

func (n Nil) Type() ObjectType {
	return NIL
}

func (n Nil) Add(o Object) Object {
	return Nil{}
}

func (n Nil) Sub(o Object) Object {
	return Nil{}
}

func (n Nil) Mul(o Object) Object {
	return Nil{}
}

func (n Nil) Div(o Object) Object {
	return Nil{}
}

func (n Nil) String() string {
	return "nil"
}

func (n Nil) GetTruthy() Boolean {
	return Boolean{Value: false}
}

func (n Nil) GetPrototype() *Map {
	return nil
}
