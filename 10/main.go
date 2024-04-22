package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Show all arguments to start
	fmt.Println(os.Args)

	// First you need to set all expected flags

	// First option is to receive a pointer to the result
	yFlag := flag.Bool("assume-yes", true, "Automatically select yes for any prompts which may need your input.")

	// Second is filling an already existing variable
	var listFlag bool
	flag.BoolVar(&listFlag, "list", false, "List all options")

	var qtyFlag int
	flag.IntVar(&qtyFlag, "qty", 123, "How many packages to show")

	// Parse arguments
	flag.Parse()

	// Print each flag result
	fmt.Println("yFlag", *yFlag)
	fmt.Println("listFlag", listFlag)
	fmt.Println("qtyFlag", qtyFlag)
	fmt.Println("flag.Args", flag.Args())

	// Notice that
	// 		go run . --list asdf
	// will return true to listFlag, but "asdf" will become a flag argument
	// this happens because --list doesn't expect nothing more, so if the following is not a flag, it will become a argument
	//		go run . asdf --list
	// will return false to listFlag because the parse meet something that is not a flag and start turning everything after into arguments

	// Notice that
	//		go run . --qty 10
	// is a flag that expect number after it, so it doesn't transform "10" into flag an argument
	//		go run . 10 --qty
	// same as above, "10" it's not a flag so will turn everything after in arguments
}
