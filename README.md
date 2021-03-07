# EXPP - tiny math expression parser

## Contents
- [Supported operations](#supported-operations)
- [Examples](#examples)
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
 
## Examples
This part contains the example of parsing and evaluating for four expressions:
```go
// examples of expressions to parse
s1 := "x * (y%3)"
s2 := "x1^(-1)"
s3 := "(price - purchasePrice) * numOfGoods * 0.87"
s4 := "sqrt(abs(-1*(2^4)))"
```

To parse expression call `expp.ParseStr()` function. `expp.Exp` string conversation returns string with [prefix style operation notation](http://www.cs.man.ac.uk/~pjj/cs212/fix.html) 
```go
exp1, _ := expp.ParseStr(s1)
fmt.Println("Parsed execution tree: ", exp1)
// Parsed execution tree: ( * x ( % y 3 ) )

exp2, _ := expp.ParseStr(s2)
fmt.Println("Parsed execution tree: ", exp2)
// Parsed execution tree: ( ^ x1 -1 )

exp3, _ := expp.ParseStr(s3)
fmt.Println("Parsed execution tree: ", exp3)
// Parsed execution tree: ( * ( * ( - price purchasePrice ) numOfGoods ) 0.87 )

exp4, _ := expp.ParseStr(s4)
fmt.Println("Parsed execution tree: ", exp4)
// Parsed execution tree: ( sqrt ( ( abs ( ( - ( * 1 ( ^ 2 4 ) ) ) ) ) ) )
```

To get sorted list of all variables used in the expression call ``expp.GetVarList()`` function:
```go
vars1 := expp.GetVarList(exp1)
fmt.Println("Variables: ", vars1)
// Variables: [x y]

vars2 := expp.GetVarList(exp2)
fmt.Println("Variables: ", vars2)
// Variables: [x1]

vars3 := expp.GetVarList(exp3)
fmt.Println("Variables: ", vars3)
// Variables: [numOfGoods price purchasePrice]

vars4 := expp.GetVarList(exp4)
fmt.Println("Variables: ", vars4)
// Variables: []
```
All variables must be defined to calculate an expression result:
```go
values1 := make(map[string]float64)
values1["x"] = 10
values1["y"] = 2

values2 := make(map[string]float64)
values2["x1"] = 50

values3 := make(map[string]float64)
values3["numOfGoods"] = 20
values3["price"] = 15.4
values3["purchasePrice"] = 10.3
``` 
Getting the result of evaluation:
```go
result1, _ := exp.Evaluate(values1)
fmt.Println("Result: ", result1)
// Result: 20

result2, _ := exp.Evaluate(values2)
fmt.Println("Result: ", result2)
// Result: 0.02
 
result3, _ := exp.Evaluate(values3)
fmt.Println("Result: ", result3)
// Result: 88.74

result4, _ := exp.Evaluate(map[string]float64{})
fmt.Println("Result: ", result4)
// Result: 4.0
```

## User-defined functions
You can add to parser your own function and set expression string presentation name
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
    expp.AddFunction(Foo, "bar")
    
    // parsing
    exp, err := expp.ParseStr(s)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    
    fmt.Println("\nParsed execution tree:", exp)
    // output: 'Parsed execution tree: ( * 10 ( bar ( 60,6,0.6 ) ) )'
    
    // execution of the expression
    result, err := exp.Evaluate(map[string]float64{})
    // output: 'Foo was called!'
    
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
- [ ] create struct `expp.Parser`, which contains parser context with included user-defined functions  
