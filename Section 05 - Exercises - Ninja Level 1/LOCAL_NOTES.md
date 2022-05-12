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

  1. Using the short declaration operator, ***ASSIGN*** these ***VALUES*** to ***VARIABLES*** with the ***IDENTIFIERS*** `x` and `y` and `z`
      1. `42`
      1. `"James Bond"`
      1. `true`
  1. Print the valuesstored in these variables using
      1. a single print statement
      1. multiple print statements 
