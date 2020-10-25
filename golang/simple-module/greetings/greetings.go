package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
		// If no name is given, return an error message
		if name == "" {
			return "", errors.New("empty name")
		}

		// Return a greeting that embeds the name in a message.
		// ":=" is a shortcut to declare and initialize a variable in one line.
		message := fmt.Sprintf("Hi, %v. Welcome!", name)
		return message, nil
}