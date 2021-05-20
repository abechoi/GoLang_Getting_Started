package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	/*
		:= operator declares and initialized in one line as opposed to:
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}
