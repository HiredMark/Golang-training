package main

import "fmt"

// can infer that a var or const is a string
const name = "MXLinux19.3"

func main() {
	// fmt.Println("Hello, World!")

	var hello string = "Hello,World!"
	// can infer variable inside a function
	hello2 := "Hello 2"

	fmt.Println(hello)
	fmt.Println(hello2)
	// name = "MXLinux19.4" // Constants don't allow for changes
	fmt.Println(name)
}

/*

Types of Variables

Bool - Boolean true or false

String

int - integer

int int8 int16 int32 int64

uint - unsigned integer

uint8 - uint64 uintptr

byte - alias for uint8

rune - alias for int32

float32 float64 < default value is float64

complex64 complex128



*/
