package main

import "fmt"

type person struct {
	first string
	last  string
	favs  []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favs:  []string{"chocolate", "martini", "rum and coke"},
	}
	p2 := person{
		first: "Miss",
		last:  "Menypenny",
		favs:  []string{"strawberry", "vanilla", "cappuccino"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favs {
		fmt.Println("  ", i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favs {
		fmt.Println("  ", i, v)
	}
}
