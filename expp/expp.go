package expp

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// base interface for Term and Node structures
type Exp interface {
	String() string
	Evaluate(vars map[string]float64) (float64, error)
	getVarList(vars map[string]interface{})
}

// struct which contains single value
type Term struct {
	Val string
}

// struct which contains two variables and binary operation
type Node struct {
	Op    rune
	L_exp Exp
	R_exp Exp
}

// parsing math expression in string, return Node tree
func ParseStr(str string) (Exp, error) {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)
	//fmt.Println("Remove spaces:", str)
	res, err := parseStr([]rune(str))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// get list of variables which used in the expression
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

// execute expression tree
func (n Node) Evaluate(vars map[string]float64) (float64, error) {
	left, err := n.L_exp.Evaluate(vars)
	if err != nil {
		return 0.0, err
	}
	right, err := n.R_exp.Evaluate(vars)
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

// returns value which contains in Term
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
