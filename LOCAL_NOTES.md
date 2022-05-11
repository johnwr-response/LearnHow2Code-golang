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
  go mod init golang-course/04/introduction-to-packages
  go mod tidy
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