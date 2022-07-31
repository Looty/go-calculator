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
	Test{2, 3, 5},
	Test{4, 8, 12},
	Test{6, 9, 15},
	Test{3, 10, 13},
}

var SubTests = []Test{
	Test{2, 3, -1},
	Test{4, 8, -4},
	Test{6, 9, -3},
	Test{3, 10, -7},
}

var MulTests = []Test{
	Test{2, 3, 6},
	Test{4, 8, 32},
	Test{6, 9, 54},
	Test{3, 10, 30},
}

var DivTests = []Test{
	Test{8, 0, 0},
	Test{16, 8, 2},
	Test{121, 11, 11},
	Test{33, 3, 11},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		res, err := Add(test.a, test.b)

		if err != nil {
			t.Errorf("Error %q not equal to expected %q", err, test.expected)
		}

		if res != test.expected {
			t.Errorf("Output %q not equal to expected %q", res, test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	for _, test := range SubTests {
		res, err := Sub(test.a, test.b)

		if err != nil {
			t.Errorf("Error %e not equal to expected %d", err, test.expected)
		}

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}

func TestMul(t *testing.T) {
	for _, test := range MulTests {
		res, err := Mul(test.a, test.b)

		if err != nil {
			t.Errorf("Error %e not equal to expected %d", err, test.expected)
		}

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}

func TestDiv(t *testing.T) {
	for _, test := range DivTests {
		res, err := Div(test.a, test.b)

		if err != nil {
			t.Errorf("Error %e not equal to expected %d", err, test.expected)
		}

		if res != test.expected {
			t.Errorf("Output %d not equal to expected %d", res, test.expected)
		}
	}
}
