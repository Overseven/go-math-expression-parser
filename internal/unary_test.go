package internal_test

import (
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/internal"
	"github.com/overseven/go-math-expression-parser/parser"
)

func TestUnaryGetVarList(t *testing.T) {
	t1, t2, t3 := "", "1.55", "c"
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}

	unary1 := internal.Unary{Op: "+", Exp: &term1}
	unary2 := internal.Unary{Op: "-", Exp: &term2}
	unary3 := internal.Unary{Op: "sqrt", Exp: &term3}

	var vars = map[string]interface{}{}
	unary1.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	unary2.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	unary3.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}
}

func TestUnaryEvaluate(t *testing.T) {
	p := parser.NewParser()
	//p.AddFunction(foo, "foo")
	//p.AddFunction(average, "average")
	//exp1 := "foo(average(2, 4, 9), 100)"

	term1 := internal.Term{Val: ""}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "a"}
	term4 := internal.Term{Val: "var3000"}

	u1 := internal.Unary{Op: "+", Exp: &term1}
	u2 := internal.Unary{Op: "-", Exp: &term2}
	u3 := internal.Unary{Op: "+", Exp: &term3}
	u4 := internal.Unary{Op: "+", Exp: &term4}
	u5 := internal.Unary{Op: "~", Exp: &term2}

	var vars = map[string]float64{"a": 17.7}
	res, err := u1.Evaluate(vars, p)
	if res != 0.0 || err != nil {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = u2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, -4.0) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = u3.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 17.7) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = u4.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
	}

	res, err = u5.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
	}
}

func TestUnaryString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "A"}
	u1 := internal.Unary{Op: "+", Exp: &term1}
	u2 := internal.Unary{Op: "-", Exp: &term2}

	if u1.String() != "( + 2 )" {
		t.Error("incorrect string conversion = " + u1.String())
	}
	if u2.String() != "( - A )" {
		t.Error("incorrect string conversion = " + u2.String())
	}
}
