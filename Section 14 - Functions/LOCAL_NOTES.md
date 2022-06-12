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
### Variadic parameter
```
md variadic-parameter
cd variadic-parameter
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/14/variadic-parameter
go mod tidy
go run .
```
- `...`
### Unfurling a slice
```
md unfurling-a-slice
cd unfurling-a-slice
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/14/unfurling-a-slice
go mod tidy
go run .
```
- If `f` is ***variadic*** with a final paramenter `p` of type `...T`, then within `f` the type of `p` is equivalent to type `[]T`.
- If `f` is invoked with no actual parameters for `p`, the value passed to `p` is `nil`.
- Otherwise, the value passed is a new slice of type `[]T` with a new underlying array whose successive elements are the actual arguments, which all must be assignable to `T`.
- The length and capacity of the slice is therefore the number of arguments bound to `p` and may differ for each call site.
