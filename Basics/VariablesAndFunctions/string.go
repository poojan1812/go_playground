package main

import "fmt"

func main() {
	// String is declared within ""
	// var word string = "peace"

	// Raw strings (without escape sequence) are declared within ``
	fmt.Println(`This is \n raw string.`)

	// Rune (an alias of int32) is used for unicode representation
	// Rune and int32 are used interchangably

	var alpha rune = 960
	var omega rune = 969
	// %c represents characters of the corresponding numeric values.
	fmt.Printf("%c%c\n", alpha, omega)

	// Strings in GO are immutable

}
