package main

import "fmt"

func main() {

	// key value store. A lot tike a dictionary. int is key and string is value
	var m map[int]string

	m = make(map[int]string)
	m[0] = "Hello"
	m[1] = " World!"
	fmt.Println(m)
	// literal notation.
	var s = map[string]string{
		"hello": "world",
	}
	fmt.Println(s)

	// insert new things into a map
	s["world"] = "hello"
	fmt.Println(s)

	// fetch nearly instantenous look up.
	w := s["hello"]
	fmt.Println(w)

	// if we don't know a key exists we can use a double assignment.

	h, ok := s["goo"]
	fmt.Println(ok, h)

	// deleting from a map.
	delete(s, "hello")
	fmt.Println(s)

}
