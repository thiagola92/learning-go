// Applications must have a "main" package
package main

import (
	"fmt"

	"rsc.io/quote/v4" // Downloaded after using "go mod tidy"
)

func main() {
	fmt.Println("Hello code!")
	fmt.Println(quote.Go())
}
