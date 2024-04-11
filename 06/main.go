package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// Split multiple values into variables
	message, err := get_hi_msg("thiago") // Change here to see error

	// Check for error
	if err != nil {
		log.Fatal("Error!")
	}

	fmt.Println(message)
}

// Go functions can return multiple values
// In this case the second value is whenever an error happend
func get_hi_msg(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("HI %v", name), nil
}
