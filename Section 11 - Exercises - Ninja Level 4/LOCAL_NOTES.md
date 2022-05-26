[back](../LOCAL_NOTES.md)

## Section 11 - Exercises - Ninja Level 4
### Hands-on exercise `04-01`
```
md hands-on-exercise-04-01
cd hands-on-exercise-04-01
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-01
go mod tidy
go run .
```
  1. Using a `COMPOSITE LITERAL`:
    1. create an `ARRAY` which holds 5 `VALUES` of `TYPE` int
    1. assign `VALUES` to each index position
  1. Range over the array and print the values
  1. Using format printing:
    1. print the `TYPE` of the `ARRAY`