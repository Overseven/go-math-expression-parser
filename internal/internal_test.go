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
func TestBase(t *testing.T) {
	type TestData struct {
		input  string
		output float64
	}

	data := []TestData{
		{"", 0.0},
		{"10+50+5", 65},
		{"2*2+2", 6},
		{"2*(2+2)", 8},
		{"100+sqrt(3^2+(2*2+3))", 104},
	}
	parser := expp.NewParser()
	for _, d := range data {
		_, err := parser.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := parser.Evaluate(map[string]float64{})
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, d.output) {
			t.Error("incorrect result, need: " + fmt.Sprintf("%f", d.output) + ", but get: " + fmt.Sprintf("%f", res))
		}
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

func TestUserFunction(t *testing.T) {
	func1 := func(args ...float64) (float64, error) {
		return args[0]*args[1] - args[2], nil
	}
	func2 := func(args ...float64) (float64, error) {
		return args[0] + args[1] + args[2], nil
	}
	pars := expp.NewParser()
	pars.AddFunction(func1, "func1")
	pars.AddFunction(func2, "func2")

	type TestVars map[string]float64
	type TestData struct {
		input  string
		vars   TestVars
		output float64
	}

	data := []TestData{
		{"func1(2, 3, 1)", TestVars{}, 5.0},
		{"func1(2^x, y, (x+y))", TestVars{"x": 3, "y": 5.1}, 32.7},
		{"func2(600, 60, 6)", TestVars{}, 666},
		{"func2(func2(700,70,7), 222, -8)", TestVars{}, 991},
	}

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
