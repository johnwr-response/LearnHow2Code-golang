package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "Q" {
		fmt.Println(x)
	} else {
		fmt.Println("Neither")
	}
}
