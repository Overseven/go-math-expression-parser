package expp

import (
	"errors"
	"sort"
	"strconv"
)

// Parser - context structure, which contains user-defined function
type Parser struct {
	operators  [levelsOfPriorities]map[string]FuncType
	expression Exp
}

// Exp - the base interface for Term and Node structures
type Exp interface {
	String() string
	Evaluate(vars map[string]float64, p *Parser) (float64, error)
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

// NewParser - create a Parser object with default set of operators and functions
func NewParser() *Parser {
	p := new(Parser)

	for i := range p.operators {
		p.operators[i] = make(map[string]FuncType)
		for key, f := range defaultOperators[i] {
			p.operators[i][key] = f
		}
	}

	return p
}

// AddFunction - add user's function and it string representation
func (p *Parser) AddFunction(f FuncType, s string) {
	p.operators[0][s] = f
}

// AddFunction - add user's function and it string representation
func (p *Parser) String() string {
	return p.expression.String()
}

// Parse - parsing a string format math expression, return Exp tree
func (p *Parser) Parse(str string) (Exp, error) {
	if indx, ok := ParenthesisIsCorrect(str); !ok {
		return nil, errors.New("incorrect parenthesis at " + strconv.Itoa(indx) + " position")
	}
	str = prepareString(str)
	//fmt.Println("Remove spaces: '" + str + "'")
	res, err := p.parseStr([]rune(str))
	if err != nil {
		return nil, err
	}
	p.expression = res
	return res, nil
}

// GetVarList - return list of variables which are used in the expression
func GetVarList(expr *Exp) []string {
	vars := make(map[string]interface{})
	(*expr).getVarList(vars)

	var sortedVars []string
	for v := range vars {
		sortedVars = append(sortedVars, v)
	}

	sort.Strings(sortedVars)
	return sortedVars
}

// Evaluate - execute expression and return result
func (p *Parser) Evaluate(vars map[string]float64) (float64, error) {
	result, err := p.expression.Evaluate(vars, p)
	return result, err
}

// Evaluate - execute expression tree
func (n *Node) Evaluate(vars map[string]float64, p *Parser) (float64, error) {
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
	result, err := p.operators[indx][n.Op](left, right)
	return result, err
}

// Evaluate - execute unary operator
func (u *Unary) Evaluate(vars map[string]float64, p *Parser) (float64, error) {
	right, err := u.exp.Evaluate(vars, p)
	if err != nil {
		return 0.0, err
	}
	indx, exist := unaryOperatorExist(u.Op, p)
	if !exist {
		return 0.0, errors.New("not supported unary operation: '" + u.Op + "'")
	}
	result, err := p.operators[indx][u.Op](right)
	return result, err
}

// Evaluate - return a value which contains in Term
func (t *Term) Evaluate(vars map[string]float64, p *Parser) (float64, error) {
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
func (f *Func) Evaluate(vars map[string]float64, p *Parser) (float64, error) {
	var args []float64
	for _, arg := range f.args {
		res, err := arg.Evaluate(vars, p)
		if err != nil {
			return -1, err
		}
		args = append(args, res)
	}
	res, err := p.operators[0][f.Op](args...)
	return res, err
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
	return "( " + string(u.Op) + " " + u.exp.String() + " )"
}

// toString conversation
func (f *Func) String() string {
	str := ""
	for _, arg := range f.args {
		str += arg.String() + ","
	}
	str = str[:len(str)-1]
	return "( " + string(f.Op) + " ( " + str + " ) )"
}
