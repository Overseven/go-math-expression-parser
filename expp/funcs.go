package expp

import "math"

// FuncType - internal type of functions
type FuncType func(args ...float64) float64

var (
	// the array of operations sorted by priority
	// priority[0] - highest priority (unary, functions)
	// priority[1] - medium priority (*, /, %, ^)
	// priority[2] - lowest priority (+, -)
	priority = [3]map[string]FuncType{
		{
			// "+":    unarySum,
			// "-":    unarySub,
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

// TODO: add errors in return values

func unarySum(vars ...float64) float64 { return vars[0] }
func unarySub(vars ...float64) float64 { return -vars[0] }

func sqrt(vars ...float64) float64 { return math.Sqrt(vars[0]) }

func mult(vars ...float64) float64        { return vars[0] * vars[1] }
func div(vars ...float64) float64         { return vars[0] / vars[1] }
func pow(vars ...float64) float64         { return math.Pow(vars[0], vars[1]) }
func divReminder(vars ...float64) float64 { return float64(int(vars[0]) % int(vars[1])) }

func sum(vars ...float64) float64 { return vars[0] + vars[1] }
func sub(vars ...float64) float64 { return vars[0] - vars[1] }
