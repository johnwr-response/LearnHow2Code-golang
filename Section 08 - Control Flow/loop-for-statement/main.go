package main

import "fmt"

func main() {
	fmt.Println("Single Condition:")
	x := 1
	for x <= 4 {
		fmt.Printf("\t%v\n", x)
		x++
	}
	fmt.Println("Done")
	fmt.Println("For eternal loop:")
	x = 1
	for {
		if x > 4 {
			break
		}
		fmt.Printf("\t%v\n", x)
		x++
	}
	fmt.Println("Done")
}
