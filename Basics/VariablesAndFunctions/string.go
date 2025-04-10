package main

import (
	"fmt"
	"unicode/utf8"
)

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
	// Characters in a string can be fetched by index

	// Caeser cipher - way of manipulating string to deliver secret message
	// Rot13 - Adds 13 characters to cipher and decipher

	// UTF-8 is the dominating chatacter encoding method
	// UTF assigns a unique unicode to each character in any language and ensures the representation is consistent across systems.
	var question = "¿Cómo estás?"
	fmt.Println(len(question), "bytes")

	fmt.Println(utf8.RuneCountInString(question), "Runes")

	var c, size = utf8.DecodeRuneInString(question)
	fmt.Printf("Rune %c is %v bytes\n", c, size)

	caesar()

}

/*
Decipher the quote from Julius Caesar:
L fdph, L vdz, L frqtxhuhg.
      —Julius Caesar
Your program will need to shift uppercase and lowercase letters by –3. Remember that
'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters.
*/

func caesar() {
	var text = "L fdph, L vdz, L frqtxhuhg"

	// Iterate through each character in the string
	for i := 0; i < len(text); i++ {
		var c = text[i] // Current character

		// Check if the character is alphabetic (uppercase or lowercase)
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			c = c - 3 // Shift the character by -3

			// Handle wrap-around for uppercase letters
			// If the shifted character goes below 'A', wrap it back to 'Z'
			if (c < 'A' && text[i] >= 'A') || (c < 'a' && text[i] >= 'a') {
				c = c + 26 // Add 26 to wrap around within the alphabet
			}
		}

		// Print the shifted character
		fmt.Printf("%c", c)
	}
}
