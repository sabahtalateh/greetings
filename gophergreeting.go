package greetings

import (
	"fmt"
)

// MagicNumber is for magic
var MagicNumber int

// GopherGreetings function
func GopherGreetings() {
	fmt.Println("Hello from greetings.GopherGreetings")
	printGreetingUnexported()
}

func init() {
	MagicNumber = 108
}
