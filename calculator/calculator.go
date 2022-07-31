package calculator

import (
	"errors"
)

var (
	ErrDivideByZero = errors.New("b cannot be divided by zero")
)

func Add(a, b int) (int, error) {
	return a + b, nil
}

func Sub(a, b int) (int, error) {
	return a - b, nil
}

func Mul(a, b int) (int, error) {
	return a * b, nil
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
