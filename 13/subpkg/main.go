// Notice that we don't need "go.mod" here
package subpkg

import "fmt"

func PrintPkg() {
	fmt.Println("Print from subpackage")
}
