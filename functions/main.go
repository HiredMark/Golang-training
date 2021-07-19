package main

import "fmt"

// consists of four major parts func + name + inputs in brackets + return type then {}
func hello() string {
	return "Hello"
}

func hola() string {
	return "Hola"
}

// passing in a g function that takes a function that is specified later, then a string it then outputs a string.
func greet(g func() string, name string) string {
	return g() + " " + name
}

// cascading functions. So a function can contain a function that returns a function. This is called a closure. This is esentially a function factory.
func greet2(g string) func(string) string {
	greeting := g
	return func(name string) string {
		return greeting + " " + name
	}
}

func main() {
	fmt.Println(hello(), ",World!")
	fmt.Println(greet(hola, "Suner"))

	h1, h2 := greet2("Hola"), greet2("Hello")

	fmt.Println(h1("Steve"))
	fmt.Println(h2("Steve"))
}
