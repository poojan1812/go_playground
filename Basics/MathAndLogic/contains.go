package main

import (
	"fmt"
	"strings"
)

// Boolean values can be true or false
func main() {
	fmt.Println("You find yourself in a dimly lit cave")

	// The value of out variable is true of false if the command variable contains a specific string.
	var command = "Walk outside"
	var out = strings.Contains(command, "outside")

	fmt.Println("Leave the cave?", out)

	// Comparision operator
	fmt.Println("You are young")

	var age = 45
	var compare = age < 25

	fmt.Printf("Am I young at %v? %v\n", age, compare)
}
