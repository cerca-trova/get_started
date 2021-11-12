package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Function Hello returns a greetings for the names person
func Hello(name string) (string, error) {
	// If name was not given, raise an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func randomFormat() string {
	formats := []string{
		"Hi,%v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
