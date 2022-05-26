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
### Hands-on exercise `03-02`
```
md hands-on-exercise-03-02
cd hands-on-exercise-03-02
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/09/hands-on-exercise-03-02
go mod tidy
go run .
```
  1. Print every rune code point of the uppercase apphabet three times. Your output should look like this:
    1. 1
      1. U+0041 'A' 
      1. U+0041 'A' 
      1. U+0041 'A' 
    1. 2
      1. U+0042 'B' 
      1. U+0042 'B' 
      1. U+0042 'B' 
    And through the rest of the alphabet characters
### Hands-on exercise `03-03`
```
md hands-on-exercise-03-03
cd hands-on-exercise-03-03
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/09/hands-on-exercise-03-03
go mod tidy
go run .
```
  1. Create a for loop using this syntax:
    - `for condition {}`
    . Have it print out the years you have been alive. 
