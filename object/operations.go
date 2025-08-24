package object

import "fmt"

func Equal(left, right Object) Object {
	if right.Type() != left.Type() {
		return Boolean{Value: false}
	}

	return Boolean{Value: right == left}
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
