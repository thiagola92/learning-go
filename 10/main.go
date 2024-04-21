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
	nFlag := flag.Bool("assume-no", false, "Automatically select no for any prompts which may need your input.")

	// Second is filling an already existing variable
	var listFlag bool
	flag.BoolVar(&listFlag, "list", false, "List all options")

	// Parse arguments
	flag.Parse()

	// Print each flag result
	fmt.Println(*yFlag)
	fmt.Println(*nFlag)
	fmt.Println(listFlag)
}
