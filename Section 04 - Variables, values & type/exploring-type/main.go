package main

import "fmt"

// Go is a STATIC programming language, not a DYNAMIC programming language
// A VARIABLE is DECLARED to hold a VALUE of a certain TYPE

// DECLARED the VARIABLE with the IDENTIFIER "y" is of TYPE int
var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z" is of TYPE string
var z = "Shaken, not stirred"
var a string = `James said:
            "Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
