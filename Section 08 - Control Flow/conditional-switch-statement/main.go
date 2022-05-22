package main

import "fmt"

func main() {
	fmt.Println("Switch on boolean value:")
	switch {
	case false:
		fmt.Println("  False - this should not print #1")
	case 2 == 4:
		fmt.Println("  False - this should not print #2")
	case 3 == 3:
		fmt.Println("  True - this should print #1")
	case 4 == 4:
		fmt.Println("  True - but this should not print #3")
	}
	fmt.Println("No default fallthrough:")
	switch {
	case false:
		fmt.Println("  False - this should not print #1")
	case 2 == 4:
		fmt.Println("  False - this should not print #2")
	case 3 == 3:
		fmt.Println("  True - this should print #1")
		fallthrough
	case 4 == 4:
		fmt.Println("  True - this should print #2")
	}
	fmt.Println("Funky fallthrough (generally just don't use it...):")
	switch {
	case false:
		fmt.Println("  False - this should not print #1")
	case 2 == 4:
		fmt.Println("  False - this should not print #2")
	case 3 == 3:
		fmt.Println("  True - this should print #1")
		fallthrough
	case 4 == 4:
		fmt.Println("  True - this should print #2")
		fallthrough
	case 7 == 9:
		fmt.Println("  False - this should not print but does with fallthrough #1")
		fallthrough
	case 11 == 14:
		fmt.Println("  False - this should not print but does with fallthrough #2")
		fallthrough
	case 15 == 15:
		fmt.Println("  True - this should print #3")
		fallthrough
	default:
		fmt.Println("  this is default and will only print if no other case is true, OR if the last case has fallthrough")
	}
	fmt.Println("Default case:")
	switch {
	case false:
		fmt.Println("  False - this should not print #1")
	case 2 == 4:
		fmt.Println("  False - this should not print #2")
	default:
		fmt.Println("  this is default and will only print if no other case is true")
	}
	fmt.Println("Switch on a value:")
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("  Miss Moneypenny")
	case "Bond":
		fmt.Println("  James Bond")
	case "Q":
		fmt.Println("  Q")
	default:
		fmt.Println("  Not found")
	}
	fmt.Println("Switch on multiple matches for a case:")
	n = "Bond"
	switch n {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("  Miss Moneypenny or James Bond or Dr No")
	case "M":
		fmt.Println("  M")
	case "Q":
		fmt.Println("  Q")
	default:
		fmt.Println("  Not found")
	}
}
