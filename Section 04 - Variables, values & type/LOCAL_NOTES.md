[back](../LOCAL_NOTES.md)

## Section 04 : Variables, values & type
### Playground
- [The Go Playground](https://go.dev/play/)
- [Better Go Playground with syntax highlight support](https://goplay.tools/)
### Hello world
  ```
  go mod init example.com/username/008-hello-world
  go mod tidy
  go fmt
  go run .\anything.go
  ```
### Introduction to packages
  ```
  md introduction-to-packages
  cd introduction-to-packages
  ni main.go
  ```
  - Variadic parameters
    - The `...<some type>` is how we signify a variadic parameter
    - The type `interface{}` is the empty interface
      - Every value is also of type `interface{}`
    - So the parameter `...interface{}` means ***give me as many arguments of any type as you'd like***
  - Throwing away returns
    - Use the `_` underscore character to throw away returns
  - You can't have unused variables in your code
    - This is code pollution
    - The compiler doesn't allow it
  - We use this notation in Go
    - `package.identifier`
      - ex: fmt.Println()
        - We would read that as: **From package `fmt` I am using `Println` func** 
    - An identifier is the name of the variable, constant, func
  - Packages
    - Code that is already written and of which you can use
    - Imports
### Short declaration operator
  ```
  md short-declaration-operator
  cd short-declaration-operator
  ni main.go
  go mod init golang-course/04/short-declaration-operator
  go mod tidy
  go run .
  ```
  - Terminology
    - `keywords`
      - These are words that are reserved for use by the Go programming language
        - Sometimes called ***reserved words***
        - You can't use a keywordfor anything other than it's purpose
    - `operator`
      - In `'2 + 2'` the `'+'` is the ***OPERATOR***
      - An operator is a character that represents an action, i.e. `'+'` is an arithmetic ***OPERATOR*** that represents addition
    - `operand`
      - In `'2 + 2'` the `'2'` are the ***OPERANDS***
    - `statement`
      - In programming a statement is the smallest standalone element of a program that expresses some action to be carried out. It is an instruction that commands the computer to perform a spesific action. A program is formed by a sequence of one or more statements.
    - `expression`
      - In programming an expression is a combination of one or more explicit values, constants, variables, operators and functions that the programming language interprets and computes to produce another value. For example `'2+3'` is an expression which evaluates to `'5'`.
  - Golang mascot
    -  
### The var keyword
  ```
  md the-var-keyword
  cd the-var-keyword
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/the-var-keyword
  go mod tidy
  go run .
  ```
  - Parens
    `'()'`
  - Curly Braces
    `'{}'`
  - Where var can be used
    - Any place within the package
  - Scope
    - Where a variable exists and is accessible
    - Best practice: ***Keep scope as `narrow` as possible***
### Exploring type
  ```
  md exploring-type
  cd exploring-type
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/exploring-type
  go mod tidy
  go run .
  ```
  - DECLARE a VARIABLE of a certain type and it can hold VALUES of that TYPE
  - `'Go suffers no fools'`
    - like `'dead men tells no tales'`
  - `var z int = 21`
    - ***package scope***
  - Primitive data types
    - In computer science, primitive data type is either of the following:
      - **A basic type** is a data type provided by a programming language as a basic bilding block. Most languages allow more complicated *composite types* to be constructedstarting from basic types.
      - **A built-in type** is a data type for which the programming language provides built-in support.
    - In most programming languages, all basic data types are built-in. In addition, many languages also provide a set of composite data types. Opinions vary as to whether a built-in type that is not basic should be considered 'primitive'.
    - [Primitive data type](https://en.wikipedia.org/wiki/Primitive_data_type)
  - Composite data types
    - In computer science, a **composite data type** or **compound data type** is any data type which can be constructed in a program using the programming language's primitive data types and other composite data types. It is sometimes called a **structure** or **aggregate data type**, although the latter term may also refer to arrays, lists, etc. <u>The act of constructing a composite type is known as `composition`</u>. 
### Zero value
  ```
  md zero-value
  cd zero-value
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/zero-value
  go mod tidy
  go run .
  ```
  - Understanding `zero value`
    - `false` for ***booleans***
    - `0` for ***integers***
    - `0.0` for ***floats***
    - `""` for ***strings***
    - `nil` for
      - ***pointers***
      - ***functions***
      - ***interfaces***
      - ***slices***
      - ***channels***
      - ***maps***
  - Use `short declaration operator` as much as possible
  - Use `var` for
    - ***zero value***
    - ***package scope***
### The fmt package
  ```
  md the-fmt-package
  cd the-fmt-package
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/the-fmt-package
  go mod tidy
  go run .
  ```
  - [package fmt](https://godocs.io/fmt)
  - basic code setup
  - %v
  - escaped characterslike \n or \t
  - format printing
    - group#1 - General printing to standard out
      - func Print()
      - func Printf()
      - func Println()
    - group#2 - Printing to a string which you can assign to a variable
      - func Sprint()
      - func Sprintf()
      - func Sprintln()
    - group#3 - Printing to a file or a web server's response
      - func Fprint()
      - func Fprintf()
      - func Fprintln()
    - Single character escapes represents special values:
      - `\a` : U+0007 alert or bell
      - `\b` : U+0008 backspace
      - `\f` : U+000C form feed
      - `\n` : U+000A line feed (newline)
      - `\r` : U+000D carriage return
      - `\t` : U+0009 horizontal tab
      - `\v` : U+000b vertical tab
      - `\\` : U+005C backslash
      - `\'` : U+0027 single quote (valid escape only within rune literals)
      - `\"` : U+0022 double quote (valid escape only within string literals)
### Creating your own type
  ```
  md creating-your-own-type
  cd creating-your-own-type
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/creating-your-own-type
  go mod tidy
  go run .
  ```
### Conversion, not casting
  ```
  md conversion-not-casting
  cd conversion-not-casting
  ni main.go
  Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
  go mod init golang-course/04/conversion-not-casting
  go mod tidy
  go run .
  ```
