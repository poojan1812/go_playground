// The default types not necessarily fulfill unique requirements and in such cases, new types can be declared
package main

import "fmt"

// Here a new type named custom_type is defined with underlying type of float64
type custom_type float64

// The variable custom_variable uses the new type custom_type which behind the scene uses float64
var custom_variable custom_type = 20

// New type here does not mean alias, so any operations on this type must require explicit converstion
// New and unique type better describe the use-case and enhance readability to avoid silly mistakes

// Unlike conventional OOP languages, GO does not have classes and objects

// Methods in GO are special functions that allows to define behaviour to custom type
// It reduces extra noise in the functions and enhances reusability and readability

// Method declaration
/*
        Receiver     Method      Result
       ==========   ========    =======
func   (k kelvin)   celcius()   celcius {
	return celcius(k - 273.15)
}
*/

// Any method have only one receiver and it acts like any other parameter
// Method call is the correct variable type followed by a dot and method name == k.celcius()

/*
Write a program with celsius, fahrenheit, and kelvin types and methods to convert from
any temperature type to any other temperature type.
*/

type celcius float64
type fahrenheit float64
type kelvin float64

// Celcius to Kelvin
func (c celcius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// Celcius to Fehrenheit
func (c celcius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// Fahrenheit to Celcius
func (f fahrenheit) celcius() celcius {
	return celcius((f - 32) * 5.0 / 9.0)
}

// Fahrenheit to Kelvin
func (f fahrenheit) kelvin() kelvin {
	return f.celcius().kelvin()
}

// Kelvin to Celcius
func (k kelvin) celcius() celcius {
	return celcius(k - 273.15)
}

// Kelvin to Fahrenheit
func (k kelvin) fahrenheit() fahrenheit {
	return k.celcius().fahrenheit()
}

func main() {
	var k kelvin = 294.0
	c := k.celcius()
	fmt.Print(k, "º K is ", c, "º C")
}

/*
Takeaway:
Use methods when you want to bind specific operations to a type, enhancing clarity and making the code easier to understand and maintain.
Use functions for broader, type-agnostic logic or operations involving multiple types.
*/
