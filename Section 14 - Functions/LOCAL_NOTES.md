[back](../LOCAL_NOTES.md)

## Section 14 - Functions
### Syntax
```
md syntax
cd syntax
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/14/syntax
go mod tidy
go run .
```
- func (r receiver) identifier(parameters) (return(s)) { ... }
- know the difference between parameters and arguments
  - we define our func with parameters (if any)
  - we call our func and pass in arguments (if any)
- ***Everything*** in Go is `PASS BY VALUE`
