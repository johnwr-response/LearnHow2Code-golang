package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Flemming")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Flemming"])
	fmt.Println(m["James"])
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println(v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
