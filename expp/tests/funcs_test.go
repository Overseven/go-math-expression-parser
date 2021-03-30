package tests

import (
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/expp"
)

func TestDefaultOperators(t *testing.T) {
	// UnarySum
	res, err := expp.UnarySum(5.4)
	if err != nil {
		t.Error(err)
	}
	if res != 5.4 {
		t.Error("incorrect UnarySum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.UnarySum(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	res, err = expp.UnarySum()
	if res != 0 || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	// UnarySub
	res, err = expp.UnarySub(5.4)
	if err != nil {
		t.Error(err)
	}
	if res != -5.4 {
		t.Error("incorrect UnarySub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}
	res, err = expp.UnarySub(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	res, err = expp.UnarySub()
	if res != 0 || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	// Sqrt
	res, err = expp.Sqrt(9.0)
	if err != nil {
		t.Error(err)
	}
	if res != 3.0 {
		t.Error("incorrect Sqrt result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}
	res, err = expp.Sqrt(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = expp.Sqrt()
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = expp.Sqrt(-9)
	if res != 0 || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	// Abs
	res, err = expp.Abs(-19.0)
	if err != nil {
		t.Error(err)
	}

	if res != 19.0 {
		t.Error("incorrect Abs result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Abs(19.0)
	if err != nil {
		t.Error(err)
	}

	if res != 19.0 {
		t.Error("incorrect Abs result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Abs(5.4, 4.3)
	if res != 0 || err == nil {
		t.Error("incorrect Abs error handling")
	}

	res, err = expp.Abs()
	if res != 0 || err == nil {
		t.Error("incorrect Abs error handling")
	}

	// Mult
	res, err = expp.Mult(11.2, 3)
	if err != nil {
		t.Error(err)
	}

	if res != 33.6 {
		t.Error("incorrect Mult result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Mult(11.2, -3)
	if err != nil {
		t.Error(err)
	}

	if res != -33.6 {
		t.Error("incorrect Mult result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Mult()
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = expp.Mult(1)
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = expp.Mult(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Mult error handling")
	}

	// Div
	res, err = expp.Div(15.0, 2)
	if err != nil {
		t.Error(err)
	}

	if res != 7.5 {
		t.Error("incorrect Div result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Div(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -4.0 {
		t.Error("incorrect Div result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Div()
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = expp.Div(1)
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = expp.Div(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Div error handling")
	}

	// Pow
	res, err = expp.Pow(2, 5)
	if err != nil {
		t.Error(err)
	}

	if res != 32.0 {
		t.Error("incorrect Pow result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Pow(-3, 3)
	if err != nil {
		t.Error(err)
	}

	if res != -27.0 {
		t.Error("incorrect Pow result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Pow()
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = expp.Pow(1)
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = expp.Pow(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Pow error handling")
	}

	// DivReminder
	res, err = expp.DivReminder(17, 5)
	if err != nil {
		t.Error(err)
	}

	if res != 2.0 {
		t.Error("incorrect DivReminder result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.DivReminder(20, 0)
	if err != nil {
		t.Error(err)
	}

	if err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = expp.DivReminder()
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = expp.DivReminder(1)
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = expp.DivReminder(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	// Sum
	res, err = expp.Sum(15.0, 0.2)
	if err != nil {
		t.Error(err)
	}

	if res != 15.2 {
		t.Error("incorrect Sum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Sum(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -33.0 {
		t.Error("incorrect Sum result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Sum()
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = expp.Sum(1)
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = expp.Sum(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Sum error handling")
	}

	// Sub
	res, err = expp.Sub(15.0, 0.2)
	if err != nil {
		t.Error(err)
	}

	if res != 15.2 {
		t.Error("incorrect Sub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Sub(-44, 11)
	if err != nil {
		t.Error(err)
	}

	if res != -33.0 {
		t.Error("incorrect Sub result: " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = expp.Sub()
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = expp.Sub(1)
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = expp.Sub(1, 2, 3)
	if res != 0 || err == nil {
		t.Error("incorrect Sub error handling")
	}
}
