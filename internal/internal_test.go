package internal_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/internal"
	expp "github.com/overseven/go-math-expression-parser/parser"
)

const float64EqualityThreshold = 1e-9

func fuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestUnaryOperatorExist(t *testing.T) {
	parser := expp.NewParser()

	ind, exist := internal.UnaryOperatorExist("~", parser)
	if ind != -1 || exist {
		t.Error("incorrect error handling")
	}

	ind, exist = internal.UnaryOperatorExist("-", parser)
	if ind != 0 || !exist {
		t.Error("incorrect result")
	}

	ind, exist = internal.UnaryOperatorExist("abs", parser)
	if ind != 0 || !exist {
		t.Error("incorrect result")
	}
}

func TestBinaryOperatorExist(t *testing.T) {
	parser := expp.NewParser()

	ind, exist := internal.BinaryOperatorExist("~", parser)
	if ind != -1 || exist {
		t.Error("incorrect error handling")
	}

	ind, exist = internal.BinaryOperatorExist("^", parser)
	if ind != 1 || !exist  {
		t.Error("incorrect result")
	}

	ind, exist = internal.BinaryOperatorExist("-", parser)
	if ind != 2 || !exist {
		t.Error("incorrect result")
	}
}

func TestEvalWithVars(t *testing.T) {
	type TestVars map[string]float64
	type TestData struct {
		input  string
		vars   TestVars
		output float64
	}

	data := []TestData{
		{"x+y", TestVars{"x": 7.7, "y": 1.2}, 8.9},
		{"x+(-y)", TestVars{"x": 100.0, "y": 12.0}, 88},
		{"x1*(x2^2)", TestVars{"x1": -100.0, "x2": 7.0}, -4900},
		{"(доход-расход)*налог", TestVars{"доход": 1520, "расход": 840, "налог": 0.87}, 591.6},
	}

	pars := expp.NewParser()

	for _, d := range data {
		_, err := pars.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := pars.Evaluate(d.vars)
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, d.output) {
			t.Error("incorrect result, need: " + fmt.Sprintf("%f", d.output) + ", but get: " + fmt.Sprintf("%f", res))
		}
	}
}

func TestParenthesisIsCorrect(t *testing.T) {
	type TestData struct {
		s       string
		correct bool
	}

	data := []TestData{
		{"", true},
		{"()", true},
		{")(", false},
		{"func2(600, 60, 6)", true},
		{"(func2(600, 60, 6))", true},
		{"func2(func2(700,70,7), 222, -8)", true},
		{"func2(func2(700,70,7), 222, -8", false},
		{"func2(func2(700,70,7), 222, -8))", false},
		{"(func2(func2(700,70,7), 222, -8)", false},
	}
	for i, d := range data {
		if _, cor := internal.ParenthesisIsCorrect(d.s); cor != data[i].correct {
			t.Error("incorrect result for " + strconv.Itoa(i) + " case: '" + d.s +
				"'. Need: " + strconv.FormatBool(data[i].correct) +
				", but get: " + strconv.FormatBool(cor))
		}
	}

}

func Foo1(args ...float64) (float64, error) {
	return 0.1, nil
}
func TestUnaryOperatorExist(t *testing.T) {
	p := expp.NewParser()
	p.AddFunction(Foo1, "foo1")
	_, exist := internal.UnaryOperatorExist("foo1", p)
	if !exist {
		t.Error("func not found")
	}

	_, exist = internal.UnaryOperatorExist("bar", p)
	if exist {
		t.Error("func false found")
	}
}

func TestBinaryOperatorExist(t *testing.T) {
	p := expp.NewParser()
	p.AddFunction(Foo1, "foo1")
	_, exist := internal.BinaryOperatorExist("^", p)
	if !exist {
		t.Error("operator not found")
	}

	_, exist = internal.BinaryOperatorExist("~", p)
	if exist {
		t.Error("operator false found")
	}
}
