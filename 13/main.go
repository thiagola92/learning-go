package main

import (
	// Use this module name + subpackage name
	"13/subpkg"
	"fmt"
)

func main() {
	fmt.Println("Print line")
	other_f()
	subpkg.PrintPkg()
}
