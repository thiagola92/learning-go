package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	// How Array declartion works?

	// "[]" holds the array size
	// "[3]string" or "[10]int"

	// You can fill the array on declaration using "{}"
	// "[]int{1,2,3}" or "[]string{"fa", "fi"}"

	// You can fill some and let others with default value
	// "[6]int{1,2}"

	// You can tell the function the size of the array
	// "func x(foobar [3]string)"
	// So it will refuse others sizes

	messages, err := great_everybody([]string{
		"foo",
		"bar",
		"foobar",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

func great_everybody(names []string) (map[string]string, error) {
	// How Map declaration works?
	// You use "make" when you need to allocate space in the heap
	// Map doesn't have a fix size, so you need to request space

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
