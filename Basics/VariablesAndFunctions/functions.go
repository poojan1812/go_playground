// A function is block of code that performs specific task every time when called.

/* Function declaration is as follows
func    Intn(n int) int
It is a composition of func keyword, function name, parameter (name and type), and return type
*/

/*
Function call is as follows
rand.Intn(10)
Starts with the package name, function name and argument
*/

// Functions, variables which begin with uppercase letter are exported across packages (like Intn, Println, Contains)

// A function accepts parameters and is invoked with arguments

// Variadic functions: Functions which accepts multiple parameters. Usually identified by ... (ellipsis)
// Combination of ellipsis and empty interface ...interface{} means the function can accept any number of parameters of any type

// Functions within the package can be invoked without specifying package name

// Isolating business logic to small independent functions provides better isolation

// Pass by value: using parameter of the function as an argument of the function call keeping independent variables, function call and main function has no related impact

package main

import (
	"fmt"
)

func kelvinToCelcius(k float64) float64 {
	k = k - 273.15
	return k
}

func celciusToFehrenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

func kelvinToFehrenheit() {
	kelvin := 0.0
	celcius := kelvinToCelcius(kelvin)
	fehrenheit := celciusToFehrenheit(celcius)

	fmt.Print(kelvin, " Kelvin is converted to ", fehrenheit, " Fehrenheit\n")

}

func main() {
	kelvin := 294.0
	celcius := kelvinToCelcius(kelvin)
	fmt.Print(kelvin, " Kelvin is ", celcius, " Celcius\n")

	fehrenheit := celciusToFehrenheit(celcius)
	fmt.Print(celcius, " Celcius is ", fehrenheit, " Fehrenheit\n")

	kelvinToFehrenheit()
}
