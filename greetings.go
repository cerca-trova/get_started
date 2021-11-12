package greetings

import (
	"errors"
	"fmt"
)

// Function Hello returns a greetings for the names person
func Hello(name string) (string, error) {
	// If name was not given, raise an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi,%v. Welcome!", name)
	return message, nil
}
