package main

import "fmt"

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
}
