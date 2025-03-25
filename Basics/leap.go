package main

import (
	"fmt"
)

// || is OR operator ==> false if both the values are false, else true
// && is AND operator ==> True if both values are true, else false.

// Leap year is evenly divisible by 400
// OR
// Leap year is evenly divisible by 4 but not by 100
func main() {

	var year = 2016
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Leap year means one extra day.!")
	} else {
		fmt.Println("Not a leap year, so no extra day :(")
	}
}
