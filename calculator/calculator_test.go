package calculator

import (
	"testing"
)

type Test struct {
	a        int
	b        int
	expected int
}

var addTests = []Test{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

var SubTests = []Test{
	{2, 3, -1},
	{4, 8, -4},
	{6, 9, -3},
	{3, 10, -7},
}

var MulTests = []Test{
	{2, 3, 6},
	{4, 8, 32},
	{6, 9, 54},
	{3, 10, 30},
}

var DivTests = []Test{
	{8, 8, 1},
	{16, 8, 2},
	{121, 11, 11},
	{33, 3, 11},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		res, _ := Add(test.a, test.b)

		if res != test.expected {
			t.Errorf("Output %q not equal to expected %q", res, test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	for _, test := range SubTests {
		res, _ := Sub(test.a, test.b)

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}

func TestMul(t *testing.T) {
	for _, test := range MulTests {
		res, _ := Mul(test.a, test.b)

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}

func TestDiv(t *testing.T) {
	for _, test := range DivTests {
		res, _ := Div(test.a, test.b)

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}
