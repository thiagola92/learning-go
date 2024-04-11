package main

import (
	"fmt"

	// External package
	my_package "github.com/thiagola92/learning-go/02"

	// go mod edit -replace github.com/thiagola92/learning-go/04=../04
	// This redirect the URL to the directory 04 in this repository
	my_second_package "github.com/thiagola92/learning-go/04"
)

func main() {
	// Using external package
	message := my_package.SayHello("thiago")
	fmt.Println(message)

	// Using local package
	my_second_package.SayBye()
}
