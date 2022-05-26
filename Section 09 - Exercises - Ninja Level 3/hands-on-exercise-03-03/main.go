package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1973
	for i <= time.Now().Year() {
		fmt.Println(i)
		i++
	}
}
