package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for i, xs := range xxs {
		fmt.Println(i, xs)
		for j, val := range xs {
			fmt.Printf("\tindex position: %v\t value:\t%v\n", j, val)
		}
	}
}
