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
    1. append the following value to the slice
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
### Hands-on exercise `04-05`
```
md hands-on-exercise-04-05
cd hands-on-exercise-04-05
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-05
go mod tidy
go run .
```
  1. To `DELETE` from a slice, we `APPEND` aling with `SLICING`. Follow these steps:
    1. start with this slice:
      `x := []int{42,43,44,45,46,47,48,49,50,51}`
    1. Use APPEND and SLICING to get these values printed
      - `[42,43,44,48,49,40,51]`
### Hands-on exercise `04-06`
```
md hands-on-exercise-04-06
cd hands-on-exercise-04-06
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-06
go mod tidy
go run .
```
  1. Create a slice to store all the names of the states in the Unites States of America. Use **make** and **append** for this. The goal is to not have the array that underlies the slice created more than once. Whats the length of your slice? Whats the capacity? Print out all the values, along with their index position, using the range clause.
    ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`
### Hands-on exercise `04-07`
```
md hands-on-exercise-04-07
cd hands-on-exercise-04-07
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-07
go mod tidy
go run .
```
  1. Create a slice of a slice of string. Store the following data in the multi-dimentional slice:
    - `James`, `Bond`, `Shaken, not stirred`
    - `Miss`, `Moneypenny`, `Helloooooo, James.`
  1. Range over the records, then range over the data in each record.
### Hands-on exercise `04-08`
```
md hands-on-exercise-04-08
cd hands-on-exercise-04-08
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/11/hands-on-exercise-04-08
go mod tidy
go run .
```
  1. Create a map with a key of `TYPE` string which is a person's "last_first" name, and a value of `TYPE` []string which stores their favourite things. Store three records in your map. Print out all of the values, along with their index position in the slice. 
  `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
  `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
  `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
