// https://gobyexample.com/command-line-subcommands
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	// NewFlagSet acts like a subcommand object, to hold all flags of that subcommand
	installCmd := flag.NewFlagSet("install", flag.ContinueOnError)

	// Now we have to use that NewFlagSet to create the variables
	yFlag := installCmd.Bool("assume-yes", true, "Automatically select yes for any prompts which may need your input.")

	// We tell NewFlagSet the string to parse (part of the os.Args, not including the subcommand)
	fmt.Println(os.Args[2:])
	installCmd.Parse(os.Args[2:])

	// go run . install
	// go run . install --assume-yes
	// go run . install --assume-yes=False
	fmt.Println("yFlag", *yFlag)
}
