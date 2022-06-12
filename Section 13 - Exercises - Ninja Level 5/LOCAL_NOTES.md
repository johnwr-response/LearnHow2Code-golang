[back](../LOCAL_NOTES.md)

## Section 13 - Exercises - Ninja Level 5
### Hands-on exercise `05-01`
```
md hands-on-exercise-05-01
cd hands-on-exercise-05-01
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/13/hands-on-exercise-05-01
go mod tidy
go run .
```
  1. Create your own type `person` which will have an underlying type of `struct` so that it can store the following data:
    - first name
    - last name
    - favorite ice creame flavours
  1. Create two values of the type `person`. Print out the values, ranging over the elements in the slice which stores the favorite flavors

