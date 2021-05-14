package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Println(err)
	}

}
