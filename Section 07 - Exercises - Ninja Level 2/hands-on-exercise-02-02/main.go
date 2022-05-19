package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Printf("42 == 42\t:\t%v\n", a)
	fmt.Printf("42 <= 43\t:\t%v\n", b)
	fmt.Printf("42 >= 43\t:\t%v\n", c)
	fmt.Printf("42 != 43\t:\t%v\n", d)
	fmt.Printf("42 < 43\t\t:\t%v\n", e)
	fmt.Printf("42 > 43\t\t:\t%v\n", f)
}
