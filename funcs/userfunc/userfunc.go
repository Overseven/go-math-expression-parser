package userfunc

import (
	"github.com/overseven/go-math-expression-parser/interfaces"
)

// Func - the struct which contains a function and an argument
type Func struct {
	Op   string
	Args []interfaces.Expression
}

func (f *Func) GetVarList(vars map[string]interface{}) {
	for _, term := range f.Args {
		term.GetVarList(vars)
	}
}

// Evaluate function
func (f *Func) Evaluate(vars map[string]float64, p interfaces.ExpParser) (float64, error) {
	var args []float64
	for _, arg := range f.Args {
		res, err := arg.Evaluate(vars, p)
		if err != nil {
			return -1, err
		}
		args = append(args, res)
	}
	res, err := p.GetFunctions()[0][f.Op](args...)
	return res, err
}

// toString conversation
func (f *Func) String() string {
	str := ""
	for _, arg := range f.Args {
		str += arg.String() + ","
	}
	str = str[:len(str)-1]
	return "( " + string(f.Op) + " ( " + str + " ) )"
}

func (f *Func) SetOperation(op string) {
	f.Op = op
}
func (f *Func) GetOperation() string {
	return f.Op
}

func (f *Func) SetArgs(args []interfaces.Expression) {
	f.Args = args
}
func (f *Func) GetArgs() []interfaces.Expression {
	return f.Args
}
