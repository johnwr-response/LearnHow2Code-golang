package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle:   vehicle{doors: 2, color: "white"},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{doors: 4, color: "black"},
		luxury:  true,
	}
	fmt.Printf("%T\n", t)
	fmt.Println("  Doors:", t.doors)
	fmt.Println("  Color:", t.color)
	fmt.Println("  Four-wheel drive: :", t.fourWheel)
	fmt.Printf("%T\n", s)
	fmt.Println("  Doors:", s.doors)
	fmt.Println("  Color:", s.color)
	fmt.Println("  Luxury: :", s.luxury)
}
