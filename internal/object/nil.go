package object

import "fmt"

type Nil struct{}

func (n Nil) Type() ObjectType {
	return NIL
}

func (n Nil) Add(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on nil")
}

func (n Nil) Sub(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on nil")
}

func (n Nil) Mul(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on nil")
}

func (n Nil) Div(o Object) (Object, error) {
	return Nil{}, fmt.Errorf("cannot perform arithmetic operations on nil")
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
