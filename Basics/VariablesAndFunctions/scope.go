package main

import (
	"fmt"
	"math/rand"
)

// Golang variable scope follows simple "Curly Braces" convention
func main() {
	var global = 1
	for global > 0 {
		var inloop = 10
		fmt.Println(inloop)
		global--
	}
	// inloop variable is out of scope for the below print statement
	// fmt.Println(inloop)

	// Main reason for shorthand variable declaration in loops is tighening the scope of variables
	// With this, only specific loop can access the inline variable and reduces the chances of accidential updates/leaks of the variable outside the scope

	//Too tightly declared variables may cause code duplication and less readability.

	// Shorthand variable declaration for if_else loop
	if short := rand.Intn(5); short == 0 {
		fmt.Println(short)
	} else if short == 1 {
		fmt.Println(short)
	} else {
		fmt.Printf("Short value in IF_ELSE loop is %v\n", short)
	}

	// shorthand variable declaration in FOR loop
	for shortcut := 3; shortcut > 0; shortcut-- {
		fmt.Printf("Shortcut value in FOR loop is %v\n", shortcut)
	}
}
