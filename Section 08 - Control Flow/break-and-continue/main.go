package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 4 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Printf("\t%v\n", x)
	}
	fmt.Println("Done")
}
