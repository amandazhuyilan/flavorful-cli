package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger
	// Set the log entry prefix and a flag to disable printing
	// time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greetings message.
	message, err := greetings.Hello("")
	// If an error occurred, print to console and exit.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}