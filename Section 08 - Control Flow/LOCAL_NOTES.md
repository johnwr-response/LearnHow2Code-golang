[back](../LOCAL_NOTES.md)

## Section 08 - Control flow    
### Understanding control flow
The ***Control flow*** is the order in which individual `statements`, `instructions` or `function calls` of a program are executed or evaluated.
### Loop - init, condition, post
```
md loop-init-condition-post
cd loop-init-condition-post
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/08/loop-init-condition-post
go mod tidy
go run .
```
### Loop - nesting loops
```
md loop-nesting-loops
cd loop-nesting-loops
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/08/loop-nesting-loops
go mod tidy
go run .
```
### Loop - for statement
```
md loop-for-statement
cd loop-for-statement
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/08/loop-for-statement
go mod tidy
go run .
```
- A `for` statement runs a block of code based on a `condition`, a `for-clause` ***or*** a `range-clause`.
- It's similar to, but not the same as, `C`. It unifies ***for*** and ***while*** and there is no ***do-while***. 
  - Like a ***C*** `for`:
    ```for init; condition; post {}```
  - Like a ***C*** `while`:
    ```for condition {} ```
  - Like a ***C*** `for(;;)`
    ```for {}```
