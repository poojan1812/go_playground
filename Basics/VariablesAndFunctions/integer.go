package main

import (
	"fmt"
	"math/big"
)

func main() {

	// Signed integer represenents both positive and negative values
	var first int
	// unsigned integer represents only positive values
	var second uint

	// Print the type of the variable with %T
	fmt.Printf("%T\n%T\n", first, second)

	// 8 bit unsigned integer ==> CSS (RGB color representation) (255,255,255)

	// Integers are accurate but have limited range.
	// Example: for uint8 an increament after 255 wraps it to 0 again.

	// Use integer types large enough to avoid wrapping.

	// Float can store big numbers but lacks precision, and integers are accurate but lack big ranges.

	// math/big library supports really really big numbers (like quintillion)
	var lightspeed = big.NewInt(240000000000000)
	fmt.Println(lightspeed)
}
