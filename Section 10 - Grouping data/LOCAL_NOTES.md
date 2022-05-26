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
### Slice - for range
```
md slice-for-range
cd slice-for-range
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/slice-for-range
go mod tidy
go run .
```
### Slice - slicing a slice
```
md slice-slicing-a-slice
cd slice-slicing-a-slice
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/slice-slicing-a-slice
go mod tidy
go run .
```
### Slice - append to a slice
```
md slice-append-to-a-slice
cd slice-append-to-a-slice
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/slice-append-to-a-slice
go mod tidy
go run .
```
### Slice - deleting from a slice
```
md slice-deleting-from-a-slice
cd slice-deleting-from-a-slice
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/10/slice-deleting-from-a-slice
go mod tidy
go run .
```
