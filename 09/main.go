package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args contains all arguments passed through cli as string
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
