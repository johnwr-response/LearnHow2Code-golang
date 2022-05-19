[back](../LOCAL_NOTES.md)

## Section 06 - Programming fundamentals    
### Bool type
```
md bool-type
cd bool-type
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/bool-type
go mod tidy
go run .
```
### How computers work
 - [Slide](../Section%2002%20-%20Course%20Overview/RESOURCES/003%2Bhow%2Bcomputers%2Bwork.pdf)
### Numeric types
```
md numeric-types
cd numeric-types
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/numeric-types
go mod tidy
go run .
```
- All numeric types are distinct except:
    - byte is an alias for uint8
    - rune is an alias for int32 (for utf-8)
- int and uint will automatically optimize for 32 or 64 bits depending on your environment
- For most cases, only `int` and `byte64` will be used. 
- Additional info from package runtime
  - Constant GOOS : Gets the operating system of your runtime
  - Constant GOARCH : Gets the processor architecture of your runtime
### String type
```
md string-type
cd string-type
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/string-type
go mod tidy
go run .
```
### Numeral systems
```
md numeral-systems
cd numeral-systems
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/numeral-systems
go mod tidy
go run .
```
### Constants
```
md constants
cd constants
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/constants
go mod tidy
go run .
```
### Iota
```
md iota
cd iota
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/06/iota
go mod tidy
go run .
```