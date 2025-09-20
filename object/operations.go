package object

import "fmt"

func Equal(left, right Object) Object {
	if right.Type() != left.Type() {
		return Boolean{Value: false}
	}

	if right.Type() == STRING {
		return Boolean{Value: right.(String).Value == left.(String).Value}
	}

	if right.Type() == NUMBER {
		return Boolean{Value: right.(Number).Value == left.(Number).Value}
	}

	if right.Type() == CLOSURE {
		refA := right.(Closure).Function
		refB := left.(Closure).Function
		return Boolean{Value: &refA == &refB}
	}

	return Boolean{Value: right == left}
}

func Mod(left, right Object) Object {
	if right.Type() != NUMBER || left.Type() != NUMBER {
		panic(fmt.Errorf("Cannot mod types %v and %v", right.Type(), left.Type()))
	}
	return Number{Value: float64(int(left.(Number).Value) % int(right.(Number).Value))}
}

func And(left, right Object) Object {
	return Boolean{Value: right.GetTruthy().Value && left.GetTruthy().Value}
}

func Or(left, right Object) Object {
	return Boolean{Value: right.GetTruthy().Value || left.GetTruthy().Value}
}

func NotEqual(left, right Object) Object {
	return Boolean{Value: !Equal(right, left).(Boolean).Value}
}

func Greater(left, right Object) Object {
	if right.Type() != left.Type() {
		panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
	}

	if right.Type() == NUMBER {
		return Boolean{Value: right.(Number).Value > left.(Number).Value}
	}

	if right.Type() == STRING {
		return Boolean{Value: right.(String).Value > left.(String).Value}
	}

	panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
}

func Less(left, right Object) Object {
	if right.Type() != left.Type() {
		panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
	}

	if right.Type() == NUMBER {
		return Boolean{Value: right.(Number).Value < left.(Number).Value}
	}

	if right.Type() == STRING {
		return Boolean{Value: right.(String).Value < left.(String).Value}
	}

	panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
}

func GreaterOrEqual(left, right Object) Object {
	if right.Type() != left.Type() {
		panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
	}

	if right.Type() == NUMBER {
		return Boolean{Value: right.(Number).Value >= left.(Number).Value}
	}

	if right.Type() == STRING {
		return Boolean{Value: right.(String).Value >= left.(String).Value}
	}

	panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
}

func LessOrEqual(left, right Object) Object {
	if right.Type() != left.Type() {
		panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
	}

	if right.Type() == NUMBER {
		return Boolean{Value: right.(Number).Value <= left.(Number).Value}
	}

	if right.Type() == STRING {
		return Boolean{Value: right.(String).Value <= left.(String).Value}
	}

	panic(fmt.Errorf("Cannot compare types %v and %v", right.Type(), left.Type()))
}

func Not(o Object) Object {
	return Boolean{
		Value: !o.GetTruthy().Value,
	}
}

func Neg(o Object) Object {
	if o.Type() != NUMBER {
		panic(fmt.Errorf("Cannot negate %v, Only numbers can have unary minus", o))
	}

	return Number{
		Value: -o.(Number).Value,
	}
}
