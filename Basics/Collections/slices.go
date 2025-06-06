// Slice does not alter the value, rather it creates a window/view for some part of the big value.

// Slices operate on `half-open interval`. i.e banks[0:4] slices the value from 0 to 3 (n-1) for the bank variable.
package main

import (
	"fmt"
)

func main() {

	banks := [...]string{
		"RBC",
		"TD",
		"ScotiaBank",
		"BMO",
		"CIBC",
		"National Bank",
		"CapitalOne",
		"Neo Financial",
	}
	fmt.Println(banks)

	// These slices can also be itereted through indexes, and further slicing of these slices is also possible.
	popularBanks := banks[0:3]
	// Omitting last index implies slice till the last element.
	smallBanks := banks[7:]

	fmt.Println(popularBanks)
	fmt.Println(smallBanks)

	// !!! Since slice is just a window through the array, any change in the slices is also reflected in the actual underlying array !!!

	// Similarly strings can also be slices, but altering the slice does not change underlying string value, and vise versa.

	// Slice indexes represent bytes and not rune - so the slicing range will represent the range of total bytes in the string.

	// COMPOSITE LITERAL SYNTAX - slices can be declared directly - and Go will create underlying array behing the scene.
	// No fixed length or ellipses like array
	otherBanks := []string{"SBI", "HDFC", "ICICI", "IDBI"}
	fmt.Printf("%T\n", otherBanks)

	// Slices over arrays are generally preferred due to flexibility of slices. - Dynamic sizing, link to underlying array, Standard Library support, flexible sub-slicing without creating new arrays.

	// Methods can also be attached to slices. - type stringSlice []string

}

func biggerSlice() {
	// Slices combined with built-in append function provides variable sized arrays that grows as needed.

	AmericanBanks := []string{"Amex", "Goldman Sachs", "Wells Fargo", "JP Morgan"}
	AmericanBanks = append(AmericanBanks, "citibank") // Append is a variadic function

	// With append, the slice can grow dynamically without defining the range

	// Length - Number of elements visible in a slice
	// Capacity - Total size of the underlying array - cap built-in function to check capacity

	// If the underlying array does not have enough capacity and we still perform append operation - it allocates another array of twice the size.
	// i.e - slice with 5 length 5 capacity --> append 1 --> new array created with 6 length 10 capacity

	// Create slice using make - preallocate capacity to avoid unnecessary array expansion
	// Here, a slice (and underlying array) is created with capacity 10 and length 0.
	paymentBanks := make([]string, 0, 5)
	paymentBanks = append(paymentBanks, "PayTM", "GooglePay", "RazorPay")

	// Use ellipses ... with the last parameter to declare a variadic function that accepts variable number of arguments
	/*
	 func terraform (prefix string, worlds ...string) []string {
	 }
	*/

	// A slice can be passed to a function using ellipses - newPlanets := terraform("New", planets...)

}
