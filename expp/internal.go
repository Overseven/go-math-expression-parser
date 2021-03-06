package expp

import (
	"errors"
	"strconv"
	"strings"
)

func prepareString(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)
	return str
}

// toString conversation
func (t Term) String() string {
	return t.Val
}

// toString conversation
func (n Node) String() string {
	return "( " + string(n.Op) + " " + n.LExp.String() + " " + n.RExp.String() + " )"
}

// toString conversation
func (u Unary) String() string {
	return "( " + string(u.Op) + " " + u.exp.String() + " )"
}

// toString conversation
func (f Func) String() string {
	str := ""
	for _, arg := range f.args {
		str += arg.String() + ","
	}
	str = str[:len(str)-1]
	return "( " + string(f.Op) + " ( " + str + " ) )"
}

func (n Node) getVarList(vars map[string]interface{}) {
	n.LExp.getVarList(vars)
	n.RExp.getVarList(vars)
}

func (u Unary) getVarList(vars map[string]interface{}) {
	u.exp.getVarList(vars)
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

func (f Func) getVarList(vars map[string]interface{}) {
	for _, term := range f.args {
		term.getVarList(vars)
	}
}

func unaryOperatorExist(op string) (index int, exist bool) {
	if _, ok := priority[0][op]; ok {
		return 0, true
	}
	return -1, false
}

func binaryOperatorExist(op string) (index int, exist bool) {
	for i:= 1; i<=2; i++ {
		if _, ok := priority[i][op]; ok {
			return i, true
		}
	}
	return -1, false
}

func parseFunc(str []rune) (f Func, isFunc bool, err error) {
	ind := strings.IndexRune(string(str), '(')
	var args [][]rune
	if ind <= 0 {
		return Func{}, false, nil
	}
	f.Op = string(str[:ind])
	if _, ok := priority[0][f.Op]; !ok {
		return Func{}, false, errors.New("function '" + f.Op + "' is not supported")
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
		arg, err := parseStr(elem)
		if err != nil {
			return f, true, err
		}
		f.args = append(f.args, arg)
	}

	return f, true, nil
}

func parseStr(str []rune) (Exp, error) {
	if len(str) == 0 {
		return Term{"0"}, nil
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
			if _, ok := priority[priorityLevel][string(c)]; ok {
				if i > 0 {
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
					return Node{string(c), resL, resR}, nil
				} else{
					right := str[i+1:]
					resR, err := parseStr(right)
					if err != nil {
						return nil, err
					}
					return Unary{string(c), resR}, nil
				}
			}
		}
	}

	// parse func
	if f, isFunc, err := parseFunc(str); err != nil {
		return nil, err
	} else if isFunc {
		return f, nil
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
