package parser

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/overseven/go-math-expression-parser/funcs"
	dfuncs "github.com/overseven/go-math-expression-parser/funcs/basic"
	"github.com/overseven/go-math-expression-parser/funcs/userfunc"
	"github.com/overseven/go-math-expression-parser/interfaces"
	"github.com/overseven/go-math-expression-parser/internal"
)

// Parser - context structure, which contains user-defined function
type Parser struct {
	Operators  [funcs.LevelsOfPriorities]map[string]funcs.FuncType
	Expression interfaces.Expression
}

// NewParser - create a Parser object with default set of operators and functions
func NewParser() *Parser {
	p := new(Parser)

	for i := range p.Operators {
		p.Operators[i] = make(map[string]funcs.FuncType)
		for key, f := range dfuncs.DefaultOperators[i] {
			p.Operators[i][key] = f
		}
	}

	return p
}

// AddFunction - add user's function and it string representation
func (p *Parser) AddFunction(f funcs.FuncType, s string) {
	p.Operators[0][s] = f
}

func (p *Parser) GetFunctions() [funcs.LevelsOfPriorities]map[string]funcs.FuncType {
	return p.Operators
}

// String - string representation of expression
func (p *Parser) String() string {
	return p.Expression.String()
}

// Parse - parsing a string format math expression, return Exp tree
func (p *Parser) Parse(str string) (interfaces.Expression, error) {
	if indx, ok := internal.ParenthesisIsCorrect(str); !ok {
		return nil, errors.New("incorrect parenthesis at " + strconv.Itoa(indx) + " position")
	}
	str = internal.PrepareString(str)
	//fmt.Println("Remove spaces: '" + str + "'")
	res, err := p.parseStr([]rune(str))
	if err != nil {
		return nil, err
	}
	p.Expression = res
	return res, nil
}

// Evaluate - execute expression and return result
func (p *Parser) Evaluate(vars map[string]float64) (float64, error) {
	result, err := p.Expression.Evaluate(vars, p)
	return result, err
}

func (p *Parser) parseFunc(str []rune) (f interfaces.Function, isFunc bool, err error) {
	ind := strings.IndexRune(string(str), '(')
	var args [][]rune
	if ind <= 0 {
		return &userfunc.Func{}, false, nil
	}
	f.SetOperation(string(str[:ind]))
	if _, ok := p.Operators[0][f.GetOperation()]; !ok {
		return &userfunc.Func{}, false, errors.New("function '" + f.GetOperation() + "' is not supported")
	}

	level := 0

	start := ind + 1

	for i := start; i <= len(str)-1; i++ {
		c := str[i]
		switch c {
		case '(':
			level++
			//end--
			continue

		case ')':
			level--
			if i != len(str)-1 {
				continue
			}
			fallthrough

		default:
			if level > 0 {
				continue

			} else if c == ',' || i == len(str)-1 {
				//fmt.Println("start:", start, "i:", i)
				args = append(args, str[start:i])
				start = i + 1

			}
		}
	}

	// fmt.Println("Func " + f.Op + " args:")
	// for i, elem := range args {
	// 	fmt.Println(strconv.Itoa(i) + ".   '" + string(elem) + "'")
	// }
	// fmt.Println("End func " + f.Op + " args.")

	for _, elem := range args {
		arg, err := p.parseStr(elem)
		if err != nil {
			return f, true, err
		}
		f.SetArgs(append(f.GetArgs(), arg))
	}

	return f, true, nil
}

func (p *Parser) parseStr(str []rune) (interfaces.Expression, error) {
	if len(str) == 0 {
		return &internal.Term{Val: "0"}, nil
	}
	level := 0

	for priorityLevel := 2; priorityLevel >= 1; priorityLevel-- {
		for i := len(str) - 1; i >= 0; i-- {
			c := str[i]
			if c == ')' {
				level++
				continue
			}
			if c == '(' {
				level--
				continue
			}
			if level > 0 {
				continue
			}
			if _, ok := p.GetFunctions()[priorityLevel][string(c)]; ok {
				if i > 0 {
					left := str[0:i]
					right := str[i+1:]
					resL, err := p.parseStr(left)
					if err != nil {
						return nil, err
					}
					resR, err := p.parseStr(right)
					if err != nil {
						return nil, err
					}
					return &internal.Node{Op: string(c), LExp: resL, RExp: resR}, nil
				}
				right := str[i+1:]
				resR, err := p.parseStr(right)
				if err != nil {
					return nil, err
				}
				return &internal.Unary{Op: string(c), Exp: resR}, nil
			}
		}
	}

	// parse func
	if f, isFunc, err := p.parseFunc(str); err != nil {
		return nil, err
	} else if isFunc {
		return f.(interfaces.Expression), nil
	}

	if str[0] == '(' {
		for i, c := range str {
			if c == '(' {
				level++
				continue
			}
			if c == ')' {
				level--
				if level == 0 {
					exp := str[1:i]
					return p.parseStr(exp)
				}
				continue
			}
		}
	} else {
		return &internal.Term{Val: string(str)}, nil
	}
	return nil, errors.New("unknow internal error")
}

// GetVarList - return list of variables which are used in the expression
func GetVarList(expr interfaces.Expression) []string {
	vars := make(map[string]interface{})
	expr.GetVarList(vars)

	var sortedVars []string
	for v := range vars {
		sortedVars = append(sortedVars, v)
	}

	sort.Strings(sortedVars)
	return sortedVars
}
