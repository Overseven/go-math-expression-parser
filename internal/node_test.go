package internal_test

import (
<<<<<<< HEAD
	"strconv"
	"testing"

	"github.com/overseven/go-math-expression-parser/internal"
	"github.com/overseven/go-math-expression-parser/parser"
)

func TestNodeGetVarList(t *testing.T) {
	t1, t2, t3, t4, t5 := "", "1.55", "c", "d", "e"
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}
	term4 := internal.Term{Val: t4}
	term5 := internal.Term{Val: t5}

	n0 := internal.Node{Op: "-", LExp: &term4, RExp: &term5}

	n1 := internal.Node{Op: "+", LExp: &term1, RExp: &term2}
	n2 := internal.Node{Op: "-", LExp: &term2, RExp: &term3}
	n3 := internal.Node{Op: "+", LExp: &term3, RExp: &term4}
	n4 := internal.Node{Op: "+", LExp: &term3, RExp: &n0}

	var vars = map[string]interface{}{}
	n1.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	n2.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}

	vars = map[string]interface{}{}
	n3.GetVarList(vars)

	if len(vars) != 2 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}
	if _, ok := vars[t4]; !ok {
		t.Error("not found t3")
	}

	vars = map[string]interface{}{}
	n4.GetVarList(vars)

	if len(vars) != 3 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}
	if _, ok := vars[t4]; !ok {
		t.Error("not found t4")
	}
	if _, ok := vars[t5]; !ok {
		t.Error("not found t5")
	}
}

func TestNodeEvaluate(t *testing.T) {
	p := parser.NewParser()
	//p.AddFunction(foo, "foo")
	//p.AddFunction(average, "average")
	//exp1 := "foo(average(2, 4, 9), 100)"

	term1 := internal.Term{Val: "1"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "a"}
	term4 := internal.Term{Val: "var3000"}

	n1 := internal.Node{Op: "-", LExp: &term1, RExp: &term2}
	n2 := internal.Node{Op: "+", LExp: &term2, RExp: &term3}
	n3 := internal.Node{Op: "+", LExp: &term3, RExp: &term4}
	n3_2 := internal.Node{Op: "+", LExp: &term4, RExp: &term3}
	n4 := internal.Node{Op: "+", LExp: &term1, RExp: &n2}
	n5 := internal.Node{Op: "~", LExp: &term1, RExp: &n2}

	var vars = map[string]float64{"a": 17.7}
	res, err := n1.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, -3.0) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = n2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, 21.7) {
		t.Error("incorrect result = " + strconv.FormatFloat(res, 'e', 4, 64))
	}

	res, err = n3.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
	}

	res, err = n3_2.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
	}

	res, err = n4.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if res != 22.7 {
		t.Error("incorrect error handling!")
	}

	res, err = n5.Evaluate(vars, p)
	if res != 0.0 || err == nil {
		t.Error("incorrect error handling!")
	}
}

func TestNodeString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "A"}
	term3 := internal.Term{Val: "b"}
	term4 := internal.Term{Val: "vVv"}

	n0 := internal.Node{Op: "+", LExp: &term1, RExp: &term2}
	n1 := internal.Node{Op: "-", LExp: &term1, RExp: &n0}
	n2 := internal.Node{Op: "pow", LExp: &term3, RExp: &term4}

	if n0.String() != "( + 2 A )" {
		t.Error("incorrect string conversion = " + n0.String())
	}
	if n1.String() != "( - 2 ( + 2 A ) )" {
		t.Error("incorrect string conversion = " + n1.String())
	}
	if n2.String() != "( pow b vVv )" {
		t.Error("incorrect string conversion = " + n2.String())
	}
}
=======
    "github.com/overseven/go-math-expression-parser/internal"
    "strconv"
    "testing"
)

func TestNode_GetVarList(t *testing.T) {
    t11, t12 := "", ""
    t21, t22 := "a", "b"
    t31, t32 := "x", "14"
    term11 := internal.Term{Val: t11}
    term12 := internal.Term{Val: t12}
    term21 := internal.Term{Val: t21}
    term22 := internal.Term{Val: t22}
    term31 := internal.Term{Val: t31}
    term32 := internal.Term{Val: t32}

    node1 := internal.Node{Op: "+", LExp: &term11, RExp: &term12}
    node2 := internal.Node{Op: "+", LExp: &term21, RExp: &term22}
    node3 := internal.Node{Op: "+", LExp: &term31, RExp: &term32}
    node4 := internal.Node{Op: "+", LExp: &node3, RExp: &term21}
    var vars = map[string]interface{}{}
    node1.GetVarList(vars)

    if len(vars) != 0 {
        t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
    }

    vars = map[string]interface{}{}
    node2.GetVarList(vars)

    if len(vars) != 2 {
        t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
    }

    if _, ok := vars[t21]; !ok {
        t.Error("not found t21")
    }
    if _, ok := vars[t22]; !ok {
        t.Error("not found t22")
    }

    vars = map[string]interface{}{}
    node3.GetVarList(vars)

    if len(vars) != 1 {
        t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
    }
    if _, ok := vars[t31]; !ok {
        t.Error("not found t31")
    }


    vars = map[string]interface{}{}
    node4.GetVarList(vars)

    if len(vars) != 2 {
        t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
    }
    if _, ok := vars[t21]; !ok {
        t.Error("not found t21")
    }
    if _, ok := vars[t31]; !ok {
        t.Error("not found t31")
    }
}

func TestNode_String(t *testing.T) {
    t11, t12 := "", ""
    t21, t22 := "a", "b"
    t31, t32 := "x", "14"
    term11 := internal.Term{Val: t11}
    term12 := internal.Term{Val: t12}
    term21 := internal.Term{Val: t21}
    term22 := internal.Term{Val: t22}
    term31 := internal.Term{Val: t31}
    term32 := internal.Term{Val: t32}

    node1 := internal.Node{Op: "+", LExp: &term11, RExp: &term12}
    node2 := internal.Node{Op: "+", LExp: &term21, RExp: &term22}
    node3 := internal.Node{Op: "+", LExp: &term31, RExp: &term32}
    node4 := internal.Node{Op: "+", LExp: &node3, RExp: &term21}

    if node1.String() != "( +   )" {
        t.Error("incorrect string conversion = " + node1.String())
    }
    if node2.String() != "( + a b )" {
        t.Error("incorrect string conversion = " + node2.String())
    }
    if node3.String() != "( + x 14 )" {
        t.Error("incorrect string conversion = " + node3.String())
    }
    if node4.String() != "( + ( + x 14 ) a )" {
        t.Error("incorrect string conversion = " + node4.String())
    }
}

func TestNode_Evaluate(t *testing.T) {
    // TODO: finish!
}
>>>>>>> 3879c25a16df5963ce0b5084b09778485a4136b3
