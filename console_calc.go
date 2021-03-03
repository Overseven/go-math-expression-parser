package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Overseven/go-math-expression-parser/expp"
)

func main() {
	exampleFlag := flag.Bool("example", false, "print example of usage")

	flag.Parse()

	if *exampleFlag {
		PrintExample()
		return
	}

	fmt.Println("Input math expression with '*', '/', '^', '%', '+', '-':")
	reader := bufio.NewReader(os.Stdin)
	formula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	exp, err := expp.ParseStr(formula)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	varsNeeded := expp.GetVarList(exp)
	vars := make(map[string]float64)

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

	result, err := exp.Evaluate(vars)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", result)
}

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
