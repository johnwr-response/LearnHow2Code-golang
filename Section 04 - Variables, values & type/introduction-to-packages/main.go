package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, function return", 42, true)
	fmt.Println(n)
	fmt.Println(err)
	m, _ := fmt.Println("Hello, ignora a function return", 42, true)
	fmt.Println(m)
}
