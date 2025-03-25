package main

import (
	"fmt"
	"math/rand"
)

// const is for data that remains constant
// var is for data that keeps changing
// decalre multiple variables at once with var()
// declare multiple variables in one line with var1, var2 = value1, value2
func main() {
	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	// increment or other operations in shorthand
	var short = 1
	short += 1000
	fmt.Println(short)

	// importing math/rand module
	var number = rand.Intn(25)
	fmt.Println(number)
}
