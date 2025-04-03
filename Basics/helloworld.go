// Starting with hello world program
// the whole program is combined in a package named main
package main

// importing the functions of a module named fmt (format)
import (
	"fmt"
)

// function contains the business logic
// function named main will be first evaluated by the compiler
// curly braces strict rules to avoid adding semicolons
func main() {

	// Here P is in upper case to imply that this function can be used outside this package as well
	// This convention enhances the readability and gives idea which functions are within the package and which are not.
	fmt.Println("Hello, World!")
}
