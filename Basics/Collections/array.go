// Arrays are ordered collections of elements with fixed length and type.
package main

import (
	"fmt"
)

func main() {

	var planets [8]string
	// Each element can be accessed through Array Index.
	planets[0] = "Earth"

	// Recommended not to go outside the bound. It throws compile time error or runtime panic when referenced with invalid index.

	// Composite literal syntax - declaring and assigning array in a single go
	// Instead of specifying the index length, 'ellipsis' [...] will count the elements while compiling
	banks := [5]string{"RBC", "CIBC", "TD", "ScotiaBank", "BMO"}
	fmt.Println(len(banks))

	// Iteration of Arrays
	for i := 0; i <= 5; i++ {
		bank := banks[i]
		fmt.Println(i, bank)
	}

	// Assigning Arrays to a variable or a function creates a complete copy of the array and is treated as a separate variable onwards.
	// Array length and type is a combined value, and operations possible only on matching types. i.e: [8]strings and [5]strings are different
	// For above reason, arrays are not often used as parametes for functions

	// 2D array - var chess [8][8]string -- Each index position holds a value
}
