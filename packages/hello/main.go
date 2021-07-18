package hello

import "fmt"

func formatGreeting(name string) string {
	return "Hello " + name
}

func SayHello(name string) {
	fmt.Println(formatGreeting(name))
}

// Capital letters are exported. Lower case variables are package private.
