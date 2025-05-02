// Maps are ideal for collection of unstructured data. Map keys can be of any type.

// Dictionary in Python is an alternative to this.

package main

import (
	"fmt"
)

func main() {
	groceryStore := map[string]int{
		"Walmart": 10,
		"FreshCo": 5,
	}
	distance := groceryStore["Walmart"]
	fmt.Println(distance)

	// Unlike arrays, Maps aren't copied when assigned to new variables or passed in a functions.
	// Maps share the same underlying data - changes to one reflects on the other one.

	shoppingStore := groceryStore
	shoppingStore["Walmart"] = 25
	fmt.Println(shoppingStore)
	fmt.Println(groceryStore)

	// Maps need to be allocated with make function like slice unless declared using composite literal.
	// modeOfPayments := make(map[string]int, 5)

}
