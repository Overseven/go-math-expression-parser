# EXPP - tiny math expression parser

## Contents
- [Supported operations](#supported-operations)
- [Example of usage](#example)
- [User-defined function](#user-defined-functions)
- [Todo](#todo)

## Supported operations
This parser supports some elements of math expressions:
- unary operators `+, -`
- binary operators `+, -, *, /, ^, %`
- any variables without spaces and operator symbols
- parenthesis `10*(x+4)`
- functions `sqrt(x), abs(x)`
- user defined functions with a comma-separated list of arguments
 
## Example
This part contains the example of parsing and evaluating expression:
```go
s := "(price - purchasePrice) * numOfGoods * 0.87"
```

Create `expp.Parser` object:
```go
parser := expp.NewParser()
```


To parse expression call `parser.Parse()` function. `expp.Exp` string conversation returns string with [prefix style operation notation](http://www.cs.man.ac.uk/~pjj/cs212/fix.html) 
```go
exp, _ := parser.Parse(s3)
fmt.Println("Parsed execution tree: ", exp3)
// Parsed execution tree: ( * ( * ( - price purchasePrice ) numOfGoods ) 0.87 )
```

To get sorted list of all variables used in the expression call ``expp.GetVarList()`` function:
```go
vars3 := expp.GetVarList(exp3)
fmt.Println("Variables: ", vars3)
// Variables: [numOfGoods price purchasePrice]
```
All variables must be defined to calculate an expression result:
```go
values := make(map[string]float64)
values["numOfGoods"] = 20
values["price"] = 15.4
values["purchasePrice"] = 10.3
``` 
Getting the result of evaluation:
```go
result, _ := parser.Evaluate(values)
fmt.Println("Result: ", result)
// Result: 88.74
```
Additional example contains in the `console_calc.go` [file](https://github.com/Overseven/go-math-expression-parser/blob/main/console_calc.go)

## User-defined functions
You can add to the parser your own function and set the expression string presentation name.
To do this, you need to create `expp.Parser` object with using `expp.NewParser` function
```go
package main

import (
	"fmt"

	"github.com/Overseven/go-math-expression-parser/expp"
)

// Foo - example of user define function
func Foo(a ...float64) (float64, error) {
	fmt.Println("Foo was called!")
	var sum float64
	for _, val := range a {
		sum += val
	}
	return sum, nil
}

func main() {
    s := "10 * bar(60, 6, 0.6)"
    
    // create parser object
    parser := expp.NewParser()
    
    // add function to parsing
    parser.AddFunction(Foo, "bar")
    
    // parsing
    exp, err := parser.Parse(s)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    
    fmt.Println("\nParsed execution tree:", exp)
    // output: 'Parsed execution tree: ( * 10 ( bar ( 60,6,0.6 ) ) )'
    
    // execution of the expression
    result, err := parser.Evaluate(map[string]float64{})
    if err != nil {
        fmt.Println("Error: ", err)
    }
    
    fmt.Println("Result: ", result)
    // output: 'Result: 666' 
}
```
## TODO
- [x] binary operators 
- [x] unary operators
- [x] simple predefined functions (like `sqrt(x)` and `abs(x)`)
- [x] comma-separated list of arguments
- [x] [user-defined functions](#user-defined-functions)
- [x] tests
- [x] create struct `expp.Parser`, which contains parser context with included user-defined functions  