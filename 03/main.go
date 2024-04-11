package main

import (
	"fmt"

	my_package "github.com/thiagola92/learning-go/02"
)

func main() {
	message := my_package.SayHello("thiago")
	fmt.Println(message)
}
