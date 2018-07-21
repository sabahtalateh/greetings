package greetings

import (
	"fmt"
)

// PrintGreeting prints the greeting.
func PrintGreeting() {
	fmt.Println("I'm printing this message from the greetings.PrintGreeting")
}

func printGreetingUnexported() {
	fmt.Println("I'm printing this message from the greetings.pritGreetingUnexported")
}
