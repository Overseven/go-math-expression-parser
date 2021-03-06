package expp

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
	return math.Abs(a - b) <= float64EqualityThreshold
}
func TestBase(t *testing.T) {
	const size = 5
	// TODO: combine input and output to struct
	input := [size]string{"", "10+50+5", "2*2+2", "2*(2+2)", "100+sqrt(3^2+(2*2+3))"}
	output := [size]float64{0.0, 65, 6, 8, 104}

	for i:= 0; i<size; i++ {
		exp, err := ParseStr(input[i])
		if err != nil {
			t.Error(err)
		}
		res, err := exp.Evaluate(map[string]float64{})
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, output[i]) {
			t.Error("incorrect result, need: " + fmt.Sprintf("%f", output[i]) + ", but get: " + fmt.Sprintf("%f", res))
		}
	}
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
	isEqual := func(arr1, arr2 []string) error{
		if len(arr1) != len(arr2){
			return errors.New("different arrays size. "+
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
	const size = 4
	input := [size]string{"", "x", "x*(sqrt(y)+1)", "(доход-расход)*налог"}

	output := [size][]string{
		{},
		{"x"},
		{"x", "y"},
		{"доход", "расход", "налог"},
	}

	for i := 0; i < size; i++ {
		exp, err := ParseStr(input[i])
		if err != nil {
			t.Error(err)
		}
		res := GetVarList(exp)
		err = allElemsIsUnique(res)
		if err != nil {
			t.Error(err)
		}
		err = isEqual(res, output[i])
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEvalWithVars(t *testing.T) {
	const size = 4
	input := [size]string{"x+y", "x+(-y)", "x1*(x2^2)", "(доход-расход)*налог"}
	vars := [size]map[string]float64{
		{"x" :7.7, 	  "y" : 1.2},
		{"x" :100.0,  "y" : 12.0},
		{"x1":-100.0, "x2": 7.0},
		{"доход":1520, "расход": 840, "налог": 0.87},
	}
	output := [size]float64{8.9, 88.0, -4900, 591.6}
	for i:= 0; i < size; i++ {
		exp, err := ParseStr(input[i])
		if err != nil {
			t.Error(err)
		}
		res, err := exp.Evaluate(vars[i])
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, output[i]) {
			t.Error("incorrect result, need: " + fmt.Sprintf("%f", output[i]) + ", but get: " + fmt.Sprintf("%f", res))
		}
	}
}

func TestUserFunction(t *testing.T) {
	func1 := func(args...float64) (float64, error) {
		return args[0] * args[1] - args[2], nil
	}
	func2 := func(args...float64)(float64, error) {
		return args[0] + args[1] + args[2], nil
	}
	AddFunction(func1, "func1")
	AddFunction(func2, "func2")

	const size = 4
	input := [size]string{
		"func1(2, 3, 1)",
		"func1(2^x, y, (x+y))",
		"func2(600, 60, 6)",
		"func2(func2(700,70,7), 222, -8)",
	}
	vars  := [size]map[string]float64{
		{},
		{"x" :3,  "y" : 5.1},
		{},
		{},
	}
	output := [size]float64{
		5.0,
		32.7,
		666,
		991,
	}

	for i:= 0; i < size; i++ {
		exp, err := ParseStr(input[i])
		if err != nil {
			t.Error(err)
		}
		res, err := exp.Evaluate(vars[i])
		if err != nil {
			t.Error(err)
		}
		if !fuzzyEqual(res, output[i]) {
			t.Error("incorrect result, need: " + fmt.Sprintf("%f", output[i]) + ", but get: " + fmt.Sprintf("%f", res))
		}
	}
}

func TestParenthesisIsCorrect(t *testing.T) {
	type ParTest struct{
		s string
		correct bool
	}

	const size = 9

	data := [size]ParTest{
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
		if _, cor := ParenthesisIsCorrect(d.s); cor != data[i].correct{
			t.Error("incorrect result for " + strconv.Itoa(i) + " case: '" + d.s +
				"'. Need: " + strconv.FormatBool(data[i].correct) +
				", but get: " + strconv.FormatBool(cor))
		}
	}

}