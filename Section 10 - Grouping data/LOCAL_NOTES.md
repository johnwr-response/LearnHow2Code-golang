[back](../LOCAL_NOTES.md)

## Section 10 - Grouping data    
### Array
```
md array
cd array
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/array
go mod tidy
go run .
```
- Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation but primarily they are a building block of slices
- Generally USE SLICES INSTEAD though
### Slice - composite literal
```
md slice-composite-literal
cd slice-composite-literal
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/slice-composite-literal
go mod tidy
go run .
```
- *Slices are super cool!*
- A `SLICE` allows to group together `VALUES` of the same `TYPE`
- `COMPOSITE LITERALS` constructs values for structs, arrays and maps
- They consists of the type of the literal followed by a brace-bound list of elements
- A `COMPOSITE LITERAL` is an expression that creates a new instance each time it is evaluated
