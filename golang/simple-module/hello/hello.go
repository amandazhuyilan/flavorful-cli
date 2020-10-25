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

	// Slice of names.
	names := []string{"Ale", "Eric", "Tony", "Brandon", "Justin", "Wayne"}
	// Request a greetings message.
	// message, err := greetings.Hello("Tyler")
	message, err := greetings.Hellos(names)
	// If an error occurred, print to console and exit.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}