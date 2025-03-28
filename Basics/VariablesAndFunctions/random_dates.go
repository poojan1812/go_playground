/* Generate random dates from a random year
For February, assign daysInMonth to 29 for leap years and 28 for other years.
Use a for loop to generate and display 10 random dates.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var era = "AD"

	year := rand.Intn(2100) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	for random := 0; random < 10; random++ {
		day := rand.Intn(daysInMonth) + 1

		// Prints the date in "AD YYYY-MM-DD" format:
		// %s for the string (era), %d for integers (year, month, day), and %02d ensures zero-padded 2-digit month/day.
		fmt.Printf("%s %d-%02d-%02d\n", era, year, month, day)
	}
}
