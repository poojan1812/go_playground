package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Number with decimal value is float64 (double-precision)
	// float64 is default floating-point type. consumes 8 bytes of memory.

	var date float64 = 31.23445
	fmt.Printf("%T\n", date)
	// zero value or initialization of variable without a value
	var price float64 = 1.0 / 3

	// Precision specifies how many digits to use after decimal point.
	// Use precision and width combination to reduce rounding off errors.
	fmt.Printf("%v\n", price)
	fmt.Printf("%.3v\n", price)

	/*
	   Save some money to buy a gift for your friend. Write a program that randomly places
	   nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it contains
	   at least $20.00. Display the running balance of the piggy bank after each deposit,
	   formatting it with an appropriate width and precision.
	*/

	var piggy float64

	for piggy <= 20.00 {
		var money float64
		var random = rand.Intn(3)

		if random == 0 {
			money = 0.05
		} else if random == 1 {
			money = 0.10
		} else {
			money = 0.25
		}
		piggy = piggy + money
		fmt.Printf("Money saved for my friend %5.2f\n", piggy) // 5.2f means 5 digits in total and 2 digits after decimal point and f means float
	}

}
