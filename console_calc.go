package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Overseven/go-math-expression-parser/expp"
)

func main() {
	// add flag to print example
	exampleFlag := flag.Bool("example", false, "print example of usage")

	flag.Parse()

	if *exampleFlag {
		PrintExample()
		return
	}

	fmt.Println("Input math expression with '*', '/', '^', '%', '+', '-':")

	// input expression
	reader := bufio.NewReader(os.Stdin)
	formula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// parsing expression
	exp, err := expp.ParseStr(formula)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

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
}

// if flag -example is presented, print example of usage
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
