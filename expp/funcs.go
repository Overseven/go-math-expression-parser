package expp

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// FuncType - internal type of functions
type FuncType func(args ...float64) (float64, error)

var (
	// the array of operations sorted by priority
	// priority[0] - highest priority (unary, functions)
	// priority[1] - medium priority (*, /, %, ^)
	// priority[2] - lowest priority (+, -)
	priority = [3]map[string]FuncType{
		{
			"+":    unarySum,
			"-":    unarySub,
			"sqrt": sqrt,
		},
		{
			"*": mult,
			"/": div,
			"%": divReminder,
			"^": pow,
		},
		{
			"+": sum,
			"-": sub,
		},
	}
)

func unarySum(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for unary sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0], nil
}
func unarySub(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for unary subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return -args[0], nil
}

func sqrt(args ...float64) (float64, error) {
	if len(args) != 1 {
		return 0, errors.New("incorrect count of args for 'sqrt' function. Need: 1, but get: " + strconv.Itoa(len(args)))
	}
	if args[0] < 0 {
		return 0, errors.New("'sqrt' function argument is negative: " + fmt.Sprintf("%f", args[0]))
	}
	return math.Sqrt(args[0]), nil
}

func mult(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for multiplication operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] * args[1] , nil
}

func div(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for division operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1] == 0.0 {
		return 0, errors.New("incorrect divisor for division operator")
	}

	return args[0] / args[1], nil
}

func pow(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for power operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return math.Pow(args[0], args[1]), nil
}

func divReminder(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for % operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	if args[1] == 0.0 {
		return 0, errors.New("incorrect divisor for % operator")
	}
	return float64(int(args[0]) % int(args[1])), nil
}

func sum(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for sum operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] + args[1], nil
}

func sub(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("incorrect count of args for subtract operator. Need: 2, but get: " + strconv.Itoa(len(args)))
	}
	return args[0] - args[1], nil
}
