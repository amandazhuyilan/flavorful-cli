package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
		// If no name is given, return an error message
		if name == "" {
			return "", errors.New("empty name")
		}

		// Return a greeting that embeds the name in a message.
		// ":=" is a shortcut to declare and initialize a variable in one line.
		// message := fmt.Sprintf("Hi, %v. Welcome!", name)
		
		// Create a message embedded with the name using random format.
		message := fmt.Sprintf(randomFormat(), name)
		return message, nil
}

func Hellos(names []string) (map[string]string, error)  {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop thorugh the recieved slices of names,
	// calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
	
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one greeting messages from a set.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hello, %v. Welcome!",
		"Great to see you, %v! I hope you are doing well!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of the formats.
	return formats[rand.Intn(len(formats))]
}
