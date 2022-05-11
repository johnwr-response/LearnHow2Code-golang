package main

import "fmt"

// DECLARE there is a VARIABLE with the identifier "y"
// ASSIGNS the VALUE 43
// declare and assign = Initiaization
var y = 43

// DECLARE there is a VARIABLE with the identifier "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, inyerfaces, slices, channels and maps.
var z int

func main() {
	// Short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42

	fmt.Println(x)
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("Again: ", y)
}
