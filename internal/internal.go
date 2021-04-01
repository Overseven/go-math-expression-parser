package internal

import (
	"errors"
	"strconv"
	"strings"

	"github.com/overseven/go-math-expression-parser/interfaces"
)

// Term - the struct which contains a single value
type Term struct {
	Val string
}

// Node - the struct which contains two variables and a binary operation
type Node struct {
	Op   string
	LExp interfaces.Expression
	RExp interfaces.Expression
}

// Unary - the struct which contains a variable and a unary operation
type Unary struct {
	Op  string
	Exp interfaces.Expression
}

func PrepareString(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)
	return str
}

func (n *Node) GetVarList(vars map[string]interface{}) {
	n.LExp.GetVarList(vars)
	n.RExp.GetVarList(vars)
}

func (u *Unary) GetVarList(vars map[string]interface{}) {
	u.Exp.GetVarList(vars)
}

func (t *Term) GetVarList(vars map[string]interface{}) {
	if t.Val == "" {
		return
	}
	if _, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return
	}
	vars[t.Val] = struct{}{}

}

func unaryOperatorExist(op string, p interfaces.ExpParser) (index int, exist bool) {
	if _, ok := p.GetFunctions()[0][op]; ok {
		return 0, true
	}
	return -1, false
}

func binaryOperatorExist(op string, p interfaces.ExpParser) (index int, exist bool) {
	for i := 1; i <= 2; i++ {
		if _, ok := p.GetFunctions()[i][op]; ok {
			return i, true
		}
	}
	return -1, false
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

// Evaluate - execute expression tree
func (n *Node) Evaluate(vars map[string]float64, p interfaces.ExpParser) (float64, error) {
	left, err := n.LExp.Evaluate(vars, p)
	if err != nil {
		return 0.0, err
	}
	right, err := n.RExp.Evaluate(vars, p)
	if err != nil {
		return 0.0, err
	}
	indx, exist := binaryOperatorExist(n.Op, p)
	if !exist {
		return 0.0, errors.New("not supported binary operation: '" + string(n.Op) + "'")
	}
	result, err := p.GetFunctions()[indx][n.Op](left, right)
	return result, err
}

// Evaluate - execute unary operator
func (u *Unary) Evaluate(vars map[string]float64, p interfaces.ExpParser) (float64, error) {
	right, err := u.Exp.Evaluate(vars, p)
	if err != nil {
		return 0.0, err
	}
	indx, exist := unaryOperatorExist(u.Op, p)
	if !exist {
		return 0.0, errors.New("not supported unary operation: '" + u.Op + "'")
	}
	result, err := p.GetFunctions()[indx][u.Op](right)
	return result, err
}

// Evaluate - return a value which contains in Term
func (t *Term) Evaluate(vars map[string]float64, p interfaces.ExpParser) (float64, error) {
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

// toString conversation
func (t *Term) String() string {
	return t.Val
}

// toString conversation
func (n *Node) String() string {
	return "( " + string(n.Op) + " " + n.LExp.String() + " " + n.RExp.String() + " )"
}

// toString conversation
func (u *Unary) String() string {
	return "( " + string(u.Op) + " " + u.Exp.String() + " )"
}
