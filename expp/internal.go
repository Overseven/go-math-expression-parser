package expp

import (
	"errors"
	"strconv"
)

var (
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

func (t Term) String() string {
	return t.Val
}

func (n Node) String() string {
	return string("( " + string(n.Op) + " " + n.L_exp.String() + " " + n.R_exp.String() + " )")
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
