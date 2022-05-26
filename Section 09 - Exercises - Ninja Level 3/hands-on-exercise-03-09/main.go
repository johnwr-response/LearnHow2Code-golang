package main

import "fmt"

func main() {
	favSport := "Skiing"
	switch favSport {
	case "Skiing":
		fmt.Println("Go to the mountains!")
	case "Swimming":
		fmt.Println("Go to the pool!")
	case "Surfing":
		fmt.Println("Go to the ocean!")
	case "Wingsuit flying":
		fmt.Println("What would you like me to say at your funeral?")
	}
}
