package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	installCmd := flag.NewFlagSet("install", flag.PanicOnError) // Change the error
	assumeYesFlag := installCmd.Bool("assume-yes", true, "Automatically select yes for any prompts which may need your input.")
	yFlag := installCmd.Bool("y", true, "Automatically select yes for any prompts which may need your input.")
	installCmd.Parse(os.Args[2:])

	// When running "go run . install -s"
	//		If ContinueOnError
	//			- It will show how to use "install" command
	//			- It will continue the code
	//		If ExitOnError
	//			- It will show how to use "install" command
	//			- It will exit
	//		If PanicOnError
	//			- It will show how to use "install" command
	//			- It will show stacktrace
	fmt.Println("assumeYesFlag", *assumeYesFlag)
	fmt.Println("yFlag", *yFlag)
}
