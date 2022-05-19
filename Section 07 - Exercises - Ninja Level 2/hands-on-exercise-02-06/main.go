package main

import "fmt"

const (
	a = 2022 - iota
	b = 2022 - iota
	c = 2022 - iota
	d = 2022 - iota
)

func main() {
	fmt.Printf("This year is\t%v\n", a)
	fmt.Printf("This year -1 is\t%v\n", b)
	fmt.Printf("This year -2 is\t%v\n", c)
	fmt.Printf("This year -3 is\t%v\n", d)
}
