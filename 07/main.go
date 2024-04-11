package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Slices are not Arrays in C, because they can grown in size.
	// Slices are not Lists in Python, because they have a type.
	options := []string{
		"Hi %v",
		"Hello %v",
		"Great to see you, %v",
	}

	message := options[rand.Intn(len(options))]
	fmt.Printf(message, "thiago")
}
