package main

import (
	"fmt"
	"math"

	"syuleyman.com/packages/hello"
)

// ^^^ look at the packages structure above. That is how imports work.
const name = "ubuntu"

func main() {

	x, y := 30.0, 40.0
	fmt.Println(math.Ceil(x / y))
	fmt.Println(math.Floor(y / x))

	hello.SayHello("Suner")
}
