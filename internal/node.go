package internal

import (
	"errors"

	"github.com/overseven/go-math-expression-parser/interfaces"
)

// Node - the struct which contains two variables and a binary operation
type Node struct {
	Op   string
	LExp interfaces.Expression
	RExp interfaces.Expression
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
	indx, exist := BinaryOperatorExist(n.Op, p)
	if !exist {
		return 0.0, errors.New("not supported binary operation: '" + string(n.Op) + "'")
	}
	result, err := p.GetFunctions()[indx][n.Op](left, right)
	return result, err
}

func (n *Node) GetVarList(vars map[string]interface{}) {
	n.LExp.GetVarList(vars)
	n.RExp.GetVarList(vars)
}

// toString conversation
func (n *Node) String() string {
	return "( " + string(n.Op) + " " + n.LExp.String() + " " + n.RExp.String() + " )"
}
