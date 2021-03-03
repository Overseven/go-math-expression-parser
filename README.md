# go-math-expression-parser
Simple math expression parser

This parser can work with `+, -, *, /, ^, %` operators, parenthesis and not sensitive to spaces
 
```go
s1 := "x * (y%3)"
s2 := "x1^(-1)"
s3 := "(price - purchasePrice) * numOfGoods * 0.87"
```

To parse expression call `expp.ParseStr` function. `expp.Exp` string conversation gives prefix [style operation representation](http://www.cs.man.ac.uk/~pjj/cs212/fix.html) 
```go
exp1, _ := expp.ParseStr(s1)
fmt.Println("Parsed expression: ", exp1)
// Parsed expression: ( * x ( % y 3 ) )

exp2, _ := expp.ParseStr(s2)
fmt.Println("Parsed expression: ", exp2)
// Parsed expression: ( ^ x1 -1 )

exp3, _ := expp.ParseStr(s3)
fmt.Println("Parsed expression: ", exp3)
// Parsed expression: ( * ( * ( - price purchasePrice ) numOfGoods ) 0.87 )
```

To get sorted list of all used in the expression variables call ``expp.GetVarList`` function:
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
```
To calculate expression values for all variables must be presented:
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
Get result of evaluation:
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
```