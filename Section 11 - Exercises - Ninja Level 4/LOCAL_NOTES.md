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
### Hands-on exercise `04-02`
```
md hands-on-exercise-04-02
cd hands-on-exercise-04-02
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-02
go mod tidy
go run .
```
  1. Using a `COMPOSITE LITERAL`:
    1. create a `SLICE` of `TYPE` int
    1. assign 10 `VALUES`
  1. Range over the slice and print the values
  1. Using format printing:
    1. print the `TYPE` of the `SLICE`
### Hands-on exercise `04-03`
```
md hands-on-exercise-04-03
cd hands-on-exercise-04-03
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-03
go mod tidy
go run .
```
  1. Using the code from the previous exercise, use `SLICING` to create the following new slices and print them
    1. `[42,43,44,45,46]`
    1. `[47,48,49,50,51]`
    1. `[44,45,46,47,48]`
    1. `[43,44,45,46,47]`
### Hands-on exercise `04-04`
```
md hands-on-exercise-04-04
cd hands-on-exercise-04-04
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-04
go mod tidy
go run .
```
  1. Follow these steps:
    1. start with this slice:
      `x := []int{42,43,44,45,46,47,48,49,50,51}`
    1. appent the following value to the slice
      `52`
    1. print the slice
    1. in ***ONE STATEMENT*** append these to the slice
      - 53
      - 54
      - 55
    1. print the slice
    1. append this slice to the slice:
      `y := []int{56,57,58,59,60}`
    1. print the slice
