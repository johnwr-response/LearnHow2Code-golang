package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v (%c):\n", i, i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	i := 198
	fmt.Printf("%v (%c):\n", i, i)
	for j := 0; j < 3; j++ {
		fmt.Printf("\t%#U\n", i)
	}
	i = 216
	fmt.Printf("%v (%c):\n", i, i)
	for j := 0; j < 3; j++ {
		fmt.Printf("\t%#U\n", i)
	}
	i = 197
	fmt.Printf("%v (%c):\n", i, i)
	for j := 0; j < 3; j++ {
		fmt.Printf("\t%#U\n", i)
	}
}
