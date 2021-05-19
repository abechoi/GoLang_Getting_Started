package main

import (
	"fmt"
	// uses quote package from pkg.go.dev, after running "go mod tidy" to download
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
