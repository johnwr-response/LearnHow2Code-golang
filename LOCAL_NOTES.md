# Learn How To Code: Google's Go (golang) Programming Language
The Ultimate Comprehensive Course - Perfect for Both Beginners and Experienced Developers

## Section 01 : Introduction
## Section 02 : Course Overview
## Section 03 : Your Development Environment
### Creating a go module
  ```
  go mod init example.com/username/repo
  go mod tidy
  go test
  ```
### Adding a dependency
  ```
  go get rsc.io/quote
  go test
  ```
### Upgrading Dependencies
  ```
  go list -m all
  go get golang.org/x/text
  go test
  go list -m all
  go get rsc.io/sampler
  go list -m all
  go test
  go list -m -versions rsc.io/sampler
  go get rsc.io/sampler@v1.3.1
  go list -m all
  go test 
  ```
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