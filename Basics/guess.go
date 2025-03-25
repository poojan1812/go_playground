/*
Write a guess-the-number program. Make the computer pick random numbers between
1–100 until it guesses your number, which you declare at the top of the program. Display
each guess and whether it was too big or too small.
*/
package main

import (
	"fmt"
	"math/rand"
)

// break keyword is used to break the execution of the loop once the given condition is met.
func main() {
	const mynumber = 27

	for {
		var random = rand.Intn(100) + 1

		if random < mynumber {
			fmt.Printf("Too small number %v\n", random)
		} else if random > mynumber {
			fmt.Printf("Too big number %v\n", random)
		} else {
			fmt.Printf("Number matches with %v\n", random)
			break
		}

	}
}
