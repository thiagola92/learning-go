package main

import (
	"fmt"

	// Notice that I'm not using the file name ("example.go").
	// Use the package name ("package my_package")
	my_package "github.com/thiagola92/learning-go/02"
	// Also notice that
	// "https://github.com/thiagola92/learning-go/tree/main/02"
	// Didn't work
)

func main() {
	message := my_package.SayHello("thiago")
	fmt.Println(message)
}
