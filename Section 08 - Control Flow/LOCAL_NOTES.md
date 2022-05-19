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
