package expp

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// the array of operations sorted by priority
	// priority[0] - highest priority (*, /, %, ^)
	// priority[1] - lowest priority (+, -)
	priority = [2]map[rune]func(a ...float64) float64{
		{
			'*': mult,
			'/': div,
			'%': divReminder,
			'^': pow,
		},
		{
			'+': sum,
			'-': sub,
		},
	}
)

// toString conversation
func (t Term) String() string {
	return t.Val
}

// toString conversation
func (n Node) String() string {
	return "( " + string(n.Op) + " " + n.L_exp.String() + " " + n.R_exp.String() + " )"
}

func (n Node) getVarList(vars map[string]interface{}) {
	n.L_exp.getVarList(vars)
	n.R_exp.getVarList(vars)
}
func (t Term) getVarList(vars map[string]interface{}) {
	if t.Val == "" {
		return
	}
	if _, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return
	}
	vars[t.Val] = struct{}{}

}

func containsInFuncMaps(op rune) (index int, ok bool) {
	for i, opMap := range priority {
		if _, ok := opMap[op]; ok {
			return i, true
		}
	}
	return -1, false
}

func parseStr(str []rune) (Exp, error) {

	level := 0
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
		if _, ok := priority[1][c]; ok && i != 0 {
			left := str[0:i]
			right := str[i+1:]
			resL, err := parseStr(left)
			if err != nil {
				return nil, err
			}
			resR, err := parseStr(right)
			if err != nil {
				return nil, err
			}
			return Node{c, resL, resR}, nil
		}
	}
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
		if _, ok := priority[0][c]; ok && i != 0 {
			left := str[0:i]
			right := str[i+1:]
			resL, err := parseStr(left)
			if err != nil {
				return nil, err
			}
			resR, err := parseStr(right)
			if err != nil {
				return nil, err
			}
			return Node{c, resL, resR}, nil
		}
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
					return parseStr(exp)
				}
				continue
			}
		}
	} else {
		return Term{string(str)}, nil
	}
	return nil, errors.New("unknow internal error")
}

func useRegExp(s string) {
	varOrNum := "(([a-z0-9]+)|([0-9]+(\\.[0-9]+)?))"
	binOps := "[\\+\\-\\*/%\\^]"
	unaryOps := "[\\-\\+]"
	funcs := "((sqrt)|(pow)|(bar))"
	binPattern := "^" + varOrNum + binOps + varOrNum + "\\z"
	unaryPattern := "^" + unaryOps + varOrNum + "\\z"
	funcPattern := "^" + funcs + "\\(" + varOrNum + "(," + varOrNum + ")*" + "\\)\\z"

	binMatch, _ := regexp.MatchString(binPattern, s)
	unMatch, _ := regexp.MatchString(unaryPattern, s)
	funcMatch, _ := regexp.MatchString(funcPattern, s)

	s = strings.ReplaceAll(s, "%", "%%")
	fmt.Printf("\"" + s + "\"\t")
	if binMatch {
		fmt.Println("binary")
	} else if unMatch {
		fmt.Println("unary")
	} else if funcMatch {
		fmt.Println("func")
	} else {
		fmt.Println("err")
	}
}
func ParseStrRegExp(str []rune) (Exp, error) {

	fmt.Println("Must be true:")
	s := "x+y"
	useRegExp(s)
	s = "x^y"
	useRegExp(s)
	s = "x%y"
	useRegExp(s)
	s = "1.3+2.9"
	useRegExp(s)
	s = "a5*165"
	useRegExp(s)
	s = "x1/y2"
	useRegExp(s)
	s = "-x1"
	useRegExp(s)
	s = "pow(x1)"
	useRegExp(s)
	s = "sqrt(x1)"
	useRegExp(s)
	s = "bar(x,y,14.5,x2)"
	useRegExp(s)

	fmt.Println("\nMust be false:")
	s = "x+y "
	useRegExp(s)
	s = " x+y"
	useRegExp(s)
	s = "x+5.6y"
	useRegExp(s)
	s = "1.4x+y"
	useRegExp(s)
	return nil, nil
}
