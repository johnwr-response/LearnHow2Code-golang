package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1973
	for {
		if i > time.Now().Year() {
			break
		}
		fmt.Println(i)
		i++
	}
}
