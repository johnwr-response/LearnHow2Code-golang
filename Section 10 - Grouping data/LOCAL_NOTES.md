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