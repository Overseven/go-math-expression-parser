package interfaces

import (
	"github.com/overseven/go-math-expression-parser/funcs"
)

type ExpParser interface {
	AddFunction(f funcs.FuncType, s string)
	GetFunctions() [funcs.LevelsOfPriorities]map[string]funcs.FuncType
	String() string
	Parse(str string) (Expression, error)
	Evaluate(vars map[string]float64) (float64, error)
}

// Exp - the base interface for Term and Node structures
type Expression interface {
	String() string
	Evaluate(vars map[string]float64, p ExpParser) (float64, error)
	GetVarList(vars map[string]interface{})
}

// Function - the struct which contains a function and an argument
type Function interface {
	Expression
	SetOperation(string)
	GetOperation() string
	SetArgs([]Expression)
	GetArgs() []Expression
}
