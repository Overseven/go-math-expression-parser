package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Overseven/go-math-expression-parser/expp"
)

// Foo - example of user define function
func Foo(a ...float64) float64 {
	fmt.Println("foo was called!")
	var sum float64
	for _, val := range a {
		sum += val
	}
	return sum
}

func main() {
	// add flag to print example
	exampleFlag := flag.Bool("example", false, "print example of usage")

	flag.Parse()

	if *exampleFlag {
		PrintExample()
		return
	}

	expp.AddFunction(Foo, "foo")
	//expp.ParseStrRegExp([]rune{})
	//
	fmt.Println("Input math expression:")

	// input expression
	reader := bufio.NewReader(os.Stdin)
	formula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// parsing expression
	exp, err := expp.ParseStr(formula)
	// if exp, isFunc, err := expp.ParseFunc([]rune("foo(2*5^2,bar(3,2),4)")); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// } else if isFunc {
	// 	fmt.Println(exp.String())
	// 	res, err := exp.Evaluate(map[string]float64{})
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	fmt.Println("Result =", res)
	// } else {
	// 	fmt.Println("This is not a function!")
	// }
	//exp, err := expp.ParseStr("sqrt(14+(x-10))")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	fmt.Println(exp)

	// get list of the variables used in the expression
	varsNeeded := expp.GetVarList(exp)

	// create [variable]value map to execute the expression
	vars := make(map[string]float64)

	// fill map
	for _, v := range varsNeeded {
		fmt.Print(v + " = ")
		var val float64
		_, err := fmt.Scan(&val)
		if err != nil {
			fmt.Println("Incorrect value!")
			return
		}
		vars[v] = val
	}

	// execute the expression using values of variables
	result, err := exp.Evaluate(vars)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// print result value
	fmt.Println("Result: ", result)
	//
}

// PrintExample prints instructions if flag -example is presented
func PrintExample() {
	fmt.Println("Instructions:")
	fmt.Println("1. Write math expression.")
	fmt.Printf("2. Define all used variables.\n\n")

	fmt.Println("You can use multiple vars in the expression:")
	fmt.Println("x ^ (y + 3) - z")
	fmt.Println("x = 2")
	fmt.Println("y = 1")
	fmt.Println("z = 4")
	fmt.Println("Result: 12")
}
