package main

import "fmt"

func main() {
	i := 42

	fmt.Println(i)

	p := &i
	fmt.Println(p)
	// this is a pointer or a memory reference. When we print it out we see a memory location. Used for memory effieint way of passing information along even if it is very big.
	fmt.Println(*p)
	// * extracts the value
	*p = *p / 2
	fmt.Println(*p)
	// we can reference the object wherever it is being stored. And we can do transformation on all the information stored in the pointer.
}
