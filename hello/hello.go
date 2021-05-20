package main

import (
	"fmt"
	"log"

	// uses quote package from pkg.go.dev, after running "go mod tidy" to download

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {

	// Use Go() from "rsc.io/quote"
	fmt.Println(quote.Go())

	// Prepends "greetings: " in front of errors
	log.SetPrefix("greetings: ")
	// 0 disables timestamps for errors
	log.SetFlags(0)

	// A slice of names
	names := []string{"Abe", "Choi", "Jong-Hae"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Abe")
	// If there is an error, print error message and exit
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Else, print message
	fmt.Println(messages)

}
