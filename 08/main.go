package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {
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
