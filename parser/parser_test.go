package parser

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"

)

const float64EqualityThreshold = 1e-9

func fuzzyEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestGetVarList(t *testing.T) {
	allElemsIsUnique := func(arr []string) error {
		for i, v := range arr {
			for j := i + 1; j < len(arr); j++ {
				if v == arr[j] {
					return errors.New("non unique var: " + v)
				}
			}
		}
		return nil
	}
	isEqual := func(arr1, arr2 []string) error {
		if len(arr1) != len(arr2) {
			return errors.New("different arrays size. " +
				"len(arr1): " + strconv.Itoa(len(arr1)) +
				", len(arr2): " + strconv.Itoa(len(arr2)))
		}
		sort.Strings(arr1)
		sort.Strings(arr2)
		for i, v := range arr1 {
			if v != arr2[i] {
				return errors.New("different arrays elements: " + v + ", " + arr2[i])
			}
		}
		return nil
	}

	type TestData struct {
		input  string
		output []string
	}

	data := []TestData{
		{"", []string{}},
		{"x", []string{"x"}},
		{"x*(sqrt(y)+1)", []string{"x", "y"}},
		{"(доход-расход)*налог", []string{"доход", "расход", "налог"}},
	}

	parser := NewParser()

	for _, d := range data {
		exp, err := parser.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res := GetVarList(exp)
		err = allElemsIsUnique(res)
		if err != nil {
			t.Error(err)
		}
		err = isEqual(res, d.output)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewParser(t *testing.T) {
	func1 := func(args ...float64) (float64, error) {
		return args[0] + 100, nil
	}
	func2 := func(args ...float64) (float64, error) {
		return args[0] + 200, nil
	}
	parser1 := NewParser()
	parser2 := NewParser()
	parser1.AddFunction(func1, "f1")
	parser1.AddFunction(func2, "f2")

	parser2.AddFunction(func1, "f2")
	parser2.AddFunction(func2, "f1")

	data := []string{"f1(1)", "f2(2)"}

	parser1.Parse(data[0])
	res, err := parser1.Evaluate(map[string]float64{})
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 101) {
		t.Error("incorrect parser1 result, need: 101.0, but get: " + fmt.Sprintf("%f", res))
	}

	parser1.Parse(data[1])
	res, err = parser1.Evaluate(map[string]float64{})
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 202) {
		t.Error("incorrect parser1 result, need: 202.0, but get: " + fmt.Sprintf("%f", res))
	}

	parser2.Parse(data[0])
	res, err = parser2.Evaluate(map[string]float64{})
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 201) {
		t.Error("incorrect parser2 result, need: 201.0, but get: " + fmt.Sprintf("%f", res))
	}

	parser2.Parse(data[1])
	res, err = parser2.Evaluate(map[string]float64{})
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 102) {
		t.Error("incorrect parser2 result, need: 102.0, but get: " + fmt.Sprintf("%f", res))
	}
}

func TestNewParser2(t *testing.T) {
	parser := NewParser()
	type TestData struct {
		input  string
		output float64
	}
	data := []TestData{
		{"15+20", 35},
		{"2^3-10", -2},
		{"sqrt(14+(4^(0.5)))", 4},
	}

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


func TestParse(t *testing.T) {
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
	parser := NewParser()
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

func TestParserString(t *testing.T){
	p := NewParser()
	p.Parse("")
	if p.String() != "0" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("1 + a")
	if p.String() != "( + 1 a )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("2^(sqrt(14+2))")
	if p.String() != "( ^ 2 ( sqrt ( ( + 14 2 ) ) ) )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("abs(- 2)")
	if p.String() != "( abs ( ( - 2 ) ) )" {
		t.Error("incorrect string conversion = " + p.String())
	}
	p.Parse("- 4")
	if p.String() != "( - 4 )" {
		t.Error("incorrect string conversion = " + p.String())
	}
}

func TestParse2(t *testing.T) {
	p := NewParser()
	exp, err := p.Parse("2 * (x+y")
	if exp != nil || err == nil {
		t.Error("incorrect error handling")
	}
	exp, err = p.Parse("Foo(x+y)")
	if exp != nil || err == nil {
		t.Error("incorrect error handling")
	}
}

func TestParseStr(t *testing.T){
	//p := Parser{}
	//i, err := p.parseStr([]rune("asf"))
	// TODO: finish
}

func TestParseFunc(t *testing.T){
	//p := Parser{}
	//_, isFunc, err := p.parseFunc([]rune("foo(a+b)"))
	// TODO: finish

}