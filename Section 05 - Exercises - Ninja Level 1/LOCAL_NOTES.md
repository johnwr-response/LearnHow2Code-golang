[back](../LOCAL_NOTES.md)

## Section 05 - Exercises - Ninja Level 1
### Hands-on exercise `01-01`
```
md hands-on-exercise-01-01
cd hands-on-exercise-01-01
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/05/hands-on-exercise-01-01
go mod tidy
go run .
```
  1. Using the short declaration operator, ***ASSIGN*** these ***VALUES*** to ***VARIABLES*** with the ***IDENTIFIERS*** `x` and `y` and `z`:
      1. `42`
      1. `"James Bond"`
      1. `true`
  1. Print the valuesstored in these variables using
      1. a single print statement
      1. multiple print statements 
### Hands-on exercise `01-02`
```
md hands-on-exercise-01-02
cd hands-on-exercise-01-02
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/05/hands-on-exercise-01-02
go mod tidy
go run .
```
  1. Use var to ***DECLARE*** three variables. The variables should have package level scope. Do not assign ***VALUES*** to the variables. Use the following ***IDENTIFIERS*** for the variables and make sure that the variables are or the following ***TYPE*** (meaning they can store ***VALUES*** of that ***TYPE***):
      1. identifier `x` type `int`
      1. identifier `y` type `string`
      1. identifier `z` type `bool`
  1. in func main
      1. print out the values for each identifier
      1. the compiler assigned values to the variables. What are these called
### Hands-on exercise `01-03`
```
md hands-on-exercise-01-03
cd hands-on-exercise-01-03
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/05/hands-on-exercise-01-03
go mod tidy
go run .
```
- Using the code from the previous exercise. 
  1. At the packet level scope, assign the following values to the three variables
      1. for x assign `42`
      1. for y assign `James Bond`
      1. for z assign `true`
  1. In func main
      1. use fmt.Sprintf to print all of the ***VALUES*** to one single string. ***ASSIGN*** the returned value of ***TYPE*** string using the short declaration operator to a ***VARIABLE*** with the ***IDENTIFIER*** `s`
      1. print out the value stored in variable `s`
### Hands-on exercise `01-04`
```
md hands-on-exercise-01-04
cd hands-on-exercise-01-04
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/05/hands-on-exercise-01-04
go mod tidy
go run .
```
  1. Create your own type. Have the underlying `TYPE` be `int`
  1. Create a ***VARIABLE*** of your new ***TYPE*** with the ***IDENTIFIER*** `x` using the ***VAR*** keyword
  1. In func main
      1. Print out the value of variable `x`
      1. Print out the type of variable `x`
      1. Assign `42` to the ***VARIABLE*** `x` using the `=` ***OPERATOR***
      1. Print the value of the variable `x`
### Hands-on exercise `01-05`
```
md hands-on-exercise-01-05
cd hands-on-exercise-01-05
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/05/hands-on-exercise-01-05
go mod tidy
go run .
```
- Using the code from the previous exercise. 
  1. At the package level scope, using the `var` keyword, create a ***VARIABLE*** with the ***IDENTIFIER*** `y`. The variable should be of the ***UNDERLYING TYPE*** of your custom type `x` 
  1. In func main:
      1. This should already be done:
          1. Print out the value of variable `x`
          1. Print out the type of variable `x`
          1. Assign `42` to the ***VARIABLE*** `x` using the `=` ***OPERATOR***
          1. Print the value of the variable `x`
      1. Now do this:
          1. now use ***CONVERSION*** to convert the ***TYPE*** of the ***VALUE*** stored in `x` to the ***UNDERLYING TYPE***
          1. then use the short declaration operator to ***ASSIGN*** that value to `y`
          1. print out the value of variable `y`
          1. Print out the type of variable `y`
