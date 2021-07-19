package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	// Structs or custom types. Related data stored in a shared memory point to make it easier to access.

	p := Point{0, 1}
	fmt.Println(p)

	p.X = 1
	fmt.Println(p)

	p1 := &p
	fmt.Println(p1.X)
	// We can reference pointer reference data by passing it around as a pointer and refering to a value within the pointer.
	p2 := Point{X: 12} // Y = 0
	p3 := Point{}      // {0,0}
	p4 := &Point{}     // start with a pointer

	fmt.Println(p2, p3, p4)
}
