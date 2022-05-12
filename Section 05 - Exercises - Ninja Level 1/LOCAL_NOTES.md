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