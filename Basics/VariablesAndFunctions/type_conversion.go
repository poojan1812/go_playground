package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Types doesn't mix. All type of variables must be converted into one before performing any operations.

	age := 50
	// converted := float64(age)

	// Age is converted to string through Itoa function and hence join operation works fine
	str := "Here I converted age to string " + strconv.Itoa(age)
	fmt.Println(str)

	secstr := fmt.Sprintf("Converted %v using sprintf", age)
	fmt.Println(secstr)

	// Once a variable is declared, its type cannot be changes
	// Go follows static typing

	input()

}

/*
Write a program that converts strings to Booleans:
 The strings “true”, “yes”, or “1” are true.
 The strings “false”, “no”, or “0” are false.
 Display an error message for any other values.
*/

func input() {
	str := "true"
	boolean := true

	switch str {
	case "true", "yes", "1":
		boolean = true
	case "false", "no", "0":
		boolean = false
	default:
		fmt.Println(str, "is invalid input")
	}
	fmt.Println(boolean)
}
