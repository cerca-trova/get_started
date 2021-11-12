package greetings

import (
	"fmt"
)

// Function Hello returns a greetings for the names person
func Hello(name string) string {
	message := fmt.Sprintf("Hi,%v. Welcome!", name)
	return message
}
