[back](../LOCAL_NOTES.md)

## Section 07 - Exercises - Ninja Level 2
### Hands-on exercise `02-01`
```
md hands-on-exercise-02-01
cd hands-on-exercise-02-01
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/07/hands-on-exercise-02-01
go mod tidy
go run .
```
  1. Write a program that prints a number in decimal, binary and hex