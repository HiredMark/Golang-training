package main

import (
	"fmt"
	"math/rand"
	"time"
)

const name = "ubuntu"

func countToTen() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

// only for loops exist in go. It uses C syntax for most of it.

func sumtoThousand() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

// this is an example of a while conditional

func ifTrue() {
	b := true
	if b {
		fmt.Println("B is true")
	}

	if t := true; t {
		fmt.Println("T is also true")
	}

	if i := 10; i == 0 {
		fmt.Println("i is zero")
	} else if i > 100 {
		fmt.Println("i is greater than 100")
	} else {
		fmt.Println("i is between 0 and 100")
	}
}

// example of If and Else boolean operators

func switchExample() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	switch x := r.Intn(5); x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is something else")
	}
	switch y := r.Intn(100); {
	case y < 10:
		fmt.Println("y is less than 10")
	case y < 50:
		fmt.Println("y is less than 50")
	default:
		fmt.Println("x is big")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")

	}

	fmt.Println(t)

}

// example of switches .

func main() {

	countToTen()
	sumtoThousand()
	ifTrue()
	switchExample()
}
