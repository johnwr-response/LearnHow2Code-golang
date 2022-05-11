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