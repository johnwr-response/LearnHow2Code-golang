[back](../LOCAL_NOTES.md)

## Section 09 - Exercises - Ninja Level 3
### Hands-on exercise `03-01`
```
md hands-on-exercise-03-01
cd hands-on-exercise-03-01
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/09/hands-on-exercise-03-01
go mod tidy
go run .
```
  1. Print every number from 1 to 10000
