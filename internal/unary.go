package internal

import (
	"errors"

	"github.com/overseven/go-math-expression-parser/interfaces"
)

// Unary - the struct which contains a variable and a unary operation
type Unary struct {
	Op  string
	Exp interfaces.Expression
}

func (u *Unary) GetVarList(vars map[string]interface{}) {
	u.Exp.GetVarList(vars)
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

// toString conversation
func (u *Unary) String() string {
	return "( " + string(u.Op) + " " + u.Exp.String() + " )"
}
