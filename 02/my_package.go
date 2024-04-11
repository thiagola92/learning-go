// Notice that this is not an application,
// so we are not using "package main".
package my_package

import "fmt"

// Exported functions that with capital letter.
func SayHello(name string) string {
	message := fmt.Sprintf("Hi %v, welcome!", name)
	return message
}
