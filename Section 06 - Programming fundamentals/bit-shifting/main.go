package main

import "fmt"

const (
	_   = iota
	ikb = 1 << (iota * 10)
	imb = 1 << (iota * 10)
	igb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t\t%b\n", ikb, ikb)
	fmt.Printf("%d\t\t\t%b\n", imb, imb)
	fmt.Printf("%d\t\t%b\n", igb, igb)
}
