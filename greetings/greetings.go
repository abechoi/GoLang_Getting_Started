package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) map[string]string {

	// Make a map of strings
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			messages[name] = err.Error()
		} else {
			messages[name] = message
		}

	}
	return messages
}

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
}

// functions with lowercased initials can only be used locally.
// Cannot be exported.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
