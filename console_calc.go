package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	expp "github.com/Overseven/go-math-expression-parser/parser"
)

// Foo - example of user-defined function
func Foo(a ...float64) (float64, error) {
	fmt.Println("foo was called!")
	var sum float64
	for _, val := range a {
		sum += val
	}
	return sum, nil
}

func main() {
	// add flag to print example
	exampleFlag := flag.Bool("example", false, "print example of usage")
	treeFlag := flag.Bool("tree", false, "print parsed tree of execution")
	flag.Parse()

	if *exampleFlag {
		PrintExample()
		return
	}

	parser := expp.NewParser()
	// add user function for parsing
	parser.AddFunction(Foo, "foo")

	fmt.Println("Input a math expression:")

	// input expression
	reader := bufio.NewReader(os.Stdin)
	formula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// parsing expression
	exp, err := parser.Parse(formula)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print parsed tree if flag -tree is presented
	if *treeFlag {
		fmt.Println("\nParsed execution tree:", exp)
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
	result, err := parser.Evaluate(vars)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// print result value
	fmt.Println("Result: ", result)
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
