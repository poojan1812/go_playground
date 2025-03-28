package main

import (
	"fmt"
)

// Golang variable scope follows simple "Curly Braces" convention

func main() {
	var global = 5

	for global > 0 {
		var inloop = 10
		fmt.Println(inloop)
		global--
	}
	// inloop variable is out of scope for the below print statement
	// fmt.Println(inloop)
}
