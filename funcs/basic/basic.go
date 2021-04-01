package basic

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/overseven/go-math-expression-parser/funcs"
)

var (
	// the array of operations sorted by operators
	// operators[0] - highest operators (unary, functions)
	// operators[1] - medium operators (*, /, %, ^)
	// operators[2] - lowest operators (+, -)
	DefaultOperators = [funcs.LevelsOfPriorities]map[string]funcs.FuncType{
		{
			"+":    UnarySum,
			"-":    UnarySub,
			"sqrt": Sqrt,
			"abs":  Abs,
		},
		{
			"*": Mult,
			"/": Div,
			"%": DivReminder,
			"^": Pow,
		},
		{
			"+": Sum,
			"-": Sub,
		},
	}
)

func UnarySum(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for unary sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0], nil
}
func UnarySub(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for unary subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return -args[0], nil
}

func Sqrt(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for 'sqrt' function. Need: 1, but get: " + strconv.Itoa(len(args)))
	}
	if args[0] < 0 {
		return 0, errors.New("'sqrt' function argument is negative: " + fmt.Sprintf("%f", args[0]))
	}
	return math.Sqrt(args[0]), nil
}

func Abs(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for 'abs' function. Need: 1, but get: " + strconv.Itoa(len(args)))
	}
	return math.Abs(args[0]), nil
}

func Mult(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for multiplication operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] * args[1], nil
}

func Div(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for division operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1] == 0.0 {
		return 0, errors.New("incorrect divisor for division operator")
	}

	return args[0] / args[1], nil
}

func Pow(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for power operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return math.Pow(args[0], args[1]), nil
}

func DivReminder(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for % operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1] == 0.0 {
		return 0, errors.New("incorrect divisor for % operator")
	}
	return float64(int(args[0]) % int(args[1])), nil
}

func Sum(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] + args[1], nil
}

func Sub(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] - args[1], nil
}
