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
### Hands-on exercise `05-02`
```
md hands-on-exercise-05-02
cd hands-on-exercise-05-02
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/13/hands-on-exercise-05-02
go mod tidy
go run .
```
  1. Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
  1. Access each value in the map. Print out the values, ranging over the slice.
### Hands-on exercise `05-03`
```
md hands-on-exercise-05-03
cd hands-on-exercise-05-03
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/13/hands-on-exercise-05-03
go mod tidy
go run .
```
  1. Create a new type: `vehicle`
    - The underlying type is struct
    - The fields:
      - `doors`
      - `colors`
  1. Create two new types `truck` & `sedan`
    - The underlying type of each of these new types is a struct
    - Embed the `vehicle` type in both `truck` & `sedan`
    - Give truck the field `fourWheel` which will be set to bool
    - Give sedan the field `luxury` which will be set to bool
  1. Using the `vehicle`, `truck` and `sedan` structs:
    - using a composite literal, create a value of type `truck` and assign values to the fields
    - using a composite literal, create a value of type `sedan` and assign values to the fields
  1. print out each of these values
  1. print out a single field from each of these values
### Hands-on exercise `05-04`
```
md hands-on-exercise-05-04
cd hands-on-exercise-05-04
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/13/hands-on-exercise-05-04
go mod tidy
go run .
```
  1. Create and use an anonymous struct