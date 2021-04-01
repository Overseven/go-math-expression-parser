package internal_test

import (
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