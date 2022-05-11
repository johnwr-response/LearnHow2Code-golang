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