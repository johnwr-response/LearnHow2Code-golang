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
		last:  "Moneypenny",
		favs:  []string{"strawberry", "vanilla", "cappuccino"},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, p := range m {
		fmt.Println(p.first, p.last)
		for i2, f := range p.favs {
			fmt.Println("\t", i2, f)
		}
	}
}
