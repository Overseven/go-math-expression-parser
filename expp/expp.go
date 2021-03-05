package expp

import (
	"errors"
	"sort"
	"strconv"
)

// Exp - the base interface for Term and Node structures
type Exp interface {
	String() string
	Evaluate(vars map[string]float64) (float64, error)
	getVarList(vars map[string]interface{})
}

// Term - the struct which contains a single value
type Term struct {
	Val string
}

// Node - the struct which contains two variables and a binary operation
type Node struct {
	Op   string
	LExp Exp
	RExp Exp
}

// Unary - the struct which contains a variable and a unary operation
type Unary struct {
	Op  string
	exp Exp
}

// Func - the struct which contains a function and an argument
type Func struct {
	Op   string
	args []Exp
}

// AddFunction - add user's function and it string representation
func AddFunction(f FuncType, s string) {
	priority[0][s] = f
}

// ParseStr - parsing a string format math expression, return Exp tree
func ParseStr(str string) (Exp, error) {
	if indx, ok := ParenthesisIsCorrect(str); !ok {
		return nil, errors.New("incorrect parenthesis at " + strconv.Itoa(indx) + " position")
	}
	str = prepareString(str)
	//fmt.Println("Remove spaces: '" + str + "'")
	res, err := parseStr([]rune(str))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetVarList - return list of variables which are used in the expression
func GetVarList(expr Exp) []string {
	vars := make(map[string]interface{})
	expr.getVarList(vars)

	var sortedVars []string
	for v := range vars {
		sortedVars = append(sortedVars, v)
	}

	sort.Strings(sortedVars)
	return sortedVars
}

// Evaluate - execute expression tree
func (n Node) Evaluate(vars map[string]float64) (float64, error) {
	left, err := n.LExp.Evaluate(vars)
	if err != nil {
		return 0.0, err
	}
	right, err := n.RExp.Evaluate(vars)
	if err != nil {
		return 0.0, err
	}
	indx, ok := containsInFuncMaps(n.Op)
	if !ok {
		return 0.0, errors.New("not supported operation: '" + string(n.Op) + "'")
	}
	result := priority[indx][n.Op](left, right)
	return result, nil
}

// Evaluate - return a value which contains in Term
func (t Term) Evaluate(vars map[string]float64) (float64, error) {
	if t.Val == "" {
		return 0.0, nil
	}
	if val, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return val, nil
	}
	val, ok := vars[t.Val]
	if !ok {
		return 0.0, errors.New("value '" + t.Val + " not found in map")
	}
	return val, nil
}

// Evaluate function
func (f Func) Evaluate(vars map[string]float64) (float64, error) {
	var args []float64
	for _, arg := range f.args {
		res, err := arg.Evaluate(vars)
		if err != nil {
			return -1, err
		}
		args = append(args, res)
	}
	res := priority[0][f.Op](args...)
	return res, nil
}

// ParenthesisIsCorrect - checks correct parenthesis pairs
func ParenthesisIsCorrect(str string) (index int, correct bool) {
	counter := 0
	for i, c := range str {
		switch c {
		case '(':
			counter++
		case ')':
			counter--
		default:
			continue
		}
		if counter < 0 {
			return i, false
		}
	}
	if counter > 0 {
		return len(str) - 1, false
	}
	return -1, true
}
