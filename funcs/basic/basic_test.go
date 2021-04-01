package basic_test

import (
	"math"
	"strconv"
	"testing"

	dfuncs "github.com/overseven/go-math-expression-parser/funcs/basic"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestDefaultOperators(t *testing.T) {
	// UnarySum
	res, err := dfuncs.UnarySum(5.4)
	if err != nil {
		t.Error(err)
	}
	if res != 5.4 {
		t.Error("incorrect UnarySum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.UnarySum(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	res, err = dfuncs.UnarySum()
	if res != 0 || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	// UnarySub
	res, err = dfuncs.UnarySub(5.4)
	if err != nil {
		t.Error(err)
	}
	if res != -5.4 {
		t.Error("incorrect UnarySub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}
	res, err = dfuncs.UnarySub(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	res, err = dfuncs.UnarySub()
	if res != 0 || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	// Sqrt
	res, err = dfuncs.Sqrt(9.0)
	if err != nil {
		t.Error(err)
	}
	if res != 3.0 {
		t.Error("incorrect Sqrt result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}
	res, err = dfuncs.Sqrt(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = dfuncs.Sqrt()
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = dfuncs.Sqrt(-9)
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	// Abs
	res, err = dfuncs.Abs(-19.0)
	if err != nil {
		t.Error(err)
	}

	if res != 19.0 {
		t.Error("incorrect Abs result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Abs(19.0)
	if err != nil {
		t.Error(err)
	}

	if res != 19.0 {
		t.Error("incorrect Abs result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Abs(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect Abs error handling")
	}

	res, err = dfuncs.Abs()
	if res != 0 || err == nil {
		t.Error("incorrect Abs error handling")
	}

	// Mult
	res, err = dfuncs.Mult(11.2, 3)
	if err != nil {
		t.Error(err)
	}

	if !almostEqual(res, 33.6) {
		t.Error("incorrect Mult result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Mult(11.2, -3)
	if err != nil {
		t.Error(err)
	}

	if !almostEqual(res, -33.60) {
		t.Error("incorrect Mult result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Mult()
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = dfuncs.Mult(1)
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = dfuncs.Mult(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	// Div
	res, err = dfuncs.Div(15.0, 2)
	if err != nil {
		t.Error(err)
	}

	if res != 7.5 {
		t.Error("incorrect Div result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Div(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -4.0 {
		t.Error("incorrect Div result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Div()
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = dfuncs.Div(1)
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = dfuncs.Div(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	// Pow
	res, err = dfuncs.Pow(2, 5)
	if err != nil {
		t.Error(err)
	}

	if res != 32.0 {
		t.Error("incorrect Pow result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Pow(-3, 3)
	if err != nil {
		t.Error(err)
	}

	if res != -27.0 {
		t.Error("incorrect Pow result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Pow()
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = dfuncs.Pow(1)
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = dfuncs.Pow(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	// DivReminder
	res, err = dfuncs.DivReminder(17, 5)
	if err != nil {
		t.Error(err)
	}

	if res != 2.0 {
		t.Error("incorrect DivReminder result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.DivReminder(20, 0)
	if err == nil {
		t.Error(err)
	}

	if err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder()
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder(1)
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	// Sum
	res, err = dfuncs.Sum(15.0, 0.2)
	if err != nil {
		t.Error(err)
	}

	if res != 15.2 {
		t.Error("incorrect Sum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Sum(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -33.0 {
		t.Error("incorrect Sum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Sum()
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = dfuncs.Sum(1)
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = dfuncs.Sum(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	// Sub
	res, err = dfuncs.Sub(15.0, 0.2)
	if err != nil {
		t.Error(err)
	}

	if res != 14.8 {
		t.Error("incorrect Sub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Sub(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -55.0 {
		t.Error("incorrect Sub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = dfuncs.Sub()
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = dfuncs.Sub(1)
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = dfuncs.Sub(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}
}
