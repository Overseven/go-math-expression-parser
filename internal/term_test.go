package internal_test

import (
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/funcs/userfunc"
	"github.com/overseven/go-math-expression-parser/interfaces"
	"github.com/overseven/go-math-expression-parser/internal"
	"github.com/overseven/go-math-expression-parser/parser"
)

<<<<<<< HEAD
func TestTermGetVarList(t *testing.T) {
	t1, t2, t3 := "", "1.55", "c"
=======
func TestGetVarList(t *testing.T) {
	t1, t2, t3 := "", "b", "145"
>>>>>>> 3879c25a16df5963ce0b5084b09778485a4136b3
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}

	var vars = map[string]interface{}{}
	term1.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	term2.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	term3.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}
}

func TestTermEvaluate(t *testing.T) {
	p := parser.NewParser()

	term1 := internal.Term{Val: ""}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "a"}
	term4 := internal.Term{Val: "var3000"}
	term5 := internal.Term{Val: "R"}

	var vars = map[string]float64{"a": 17.7}
	res, err := term1.Evaluate(vars, p)
	if res != 0.0 || err != nil {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = term2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 4.0) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = term3.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
<<<<<<< HEAD
	if !fuzzyEqual(res, 17.7) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = term4.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
=======
	if  !fuzzyEqual(res, 17.7) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = term4.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if  res != 30012.0 {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	_, err = term5.Evaluate(vars, p)
	if err == nil {
		t.Error("incorrect error handling")
>>>>>>> 3879c25a16df5963ce0b5084b09778485a4136b3
	}
}

func TestTermString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}
	f1 := userfunc.Func{Op: "average", Args: []interfaces.Expression{&term1, &term2, &term3}}
	f2 := userfunc.Func{Op: "foo", Args: []interfaces.Expression{&f1, &term4}}

	if f1.String() != "( average ( 2,4,9 ) )" {
		t.Error("incorrect string conversion = " + f1.String())
	}
	if f2.String() != "( foo ( ( average ( 2,4,9 ) ),100 ) )" {
		t.Error("incorrect string conversion = " + f2.String())
	}
}
