package internal_test

import (
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/funcs/userfunc"
	"github.com/overseven/go-math-expression-parser/interfaces"
	"github.com/overseven/go-math-expression-parser/internal"
	"github.com/overseven/go-math-expression-parser/parser"
)

func TestGetVarList(t *testing.T) {
	t1, t2, t3 := "a", "b", "c"
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}

	var vars = map[string]interface{}{}
	term1.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t1]; !ok {
		t.Error("not found t1")
	}

	vars = map[string]interface{}{}
	term2.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t2]; !ok {
		t.Error("not found t2")
	}

	vars = map[string]interface{}{}
	term3.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}

}

func TestEvaluate(t *testing.T) {
	p := parser.NewParser()
	//p.AddFunction(foo, "foo")
	//p.AddFunction(average, "average")
	//exp1 := "foo(average(2, 4, 9), 100)"

	term1 := internal.Term{Val: ""}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "a"}
	term4 := internal.Term{Val: "var3000"}

	var vars = map[string]float64{"a": 17.7, "var3000": 30012}
	res, err := term1.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if res != 0.0 {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = term2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !almostEqual(res, 105.0) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = f3.Evaluate(vars, p)
	if err == nil || res != -16 {
		t.Error("foo error was not handled!")
	}

	res, err = f4.Evaluate(vars, p)
	if err == nil || res != -1 {
		t.Error("foo error was not handled!")
	}
}

func TestString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}
	f1 := userfunc.Func{"average", []interfaces.Expression{&term1, &term2, &term3}}
	f2 := userfunc.Func{"foo", []interfaces.Expression{&f1, &term4}}

	if f1.String() != "( average ( 2,4,9 ) )" {
		t.Error("incorrect string conversion = " + f1.String())
	}
	if f2.String() != "( foo ( ( average ( 2,4,9 ) ),100 ) )" {
		t.Error("incorrect string conversion = " + f2.String())
	}
}
