package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is, ", x)
	y := sum()
	fmt.Println("The total is, ", y)
	z1 := sum2("James", 4, 5, 6)
	fmt.Println("The total is, ", z1)
	z2 := sum2("James")
	fmt.Println("The total is, ", z2)
}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now ", sum)
	}
	fmt.Println("The total is, ", sum)
	return sum
}
func sum2(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now ", sum)
	}
	fmt.Println("The total is, ", sum)
	return sum
}
