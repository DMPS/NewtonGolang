package newton

import (
	"testing"
)

var calcTests = []struct {
	operation, expression, expected string
}{
	{"factor", "x^2-4", "(x - 2) (x + 2)"},
	{"factor", "x^4", "x^4"},
	{"factor", "x^4+x^3-x^2", "x^2 (x^2 + x - 1)"},
	{"factor", "4x-3x+x^3", "x (x^2 + 1)"},
	{"derive", "x^2", "2 x"},
	{"derive", "x^4+x^3-x^2", "4 x^3 + 3 x^2 - 2 x"},
	{"derive", "log(x+4x^2)", "1 / (4 x^2 + x) + 8 x / (4 x^2 + x)"},
	{"derive", "sin(5x)+acos(7x)", "5 cos(5 x) + d(acos(7 x),x)"},
}

func TestCalc(t *testing.T) {
	for _, tt := range calcTests {
		actual := Calc(Data{Operation: tt.operation, Expression: tt.expression})
		if actual.Result != tt.expected {
			t.Errorf("Calc(%s, %s): expected %s, actual %s", tt.operation, tt.expression, tt.expected, actual.Result)
		}
	}
}
