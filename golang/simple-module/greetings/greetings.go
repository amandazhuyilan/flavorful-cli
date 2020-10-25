package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
		// Return a greeting that embeds the name in a message.
		// ":=" is a shortcut to declare and initialize a variable in one line.
		message := fmt.Sprintf("Hi, %v. Welcome!", name)
		return message
}