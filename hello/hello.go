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

	// Request a greeting message
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
