// Functions are first-class in Go - they work at all places where integers, strings and other types work.
// These functions can be trated like a regular value - assign them to variables, pass them as arguments to other functions, or return them from functions.

// Example of weather sensors returning temperature reading from 150 to 300 Celcius.
// Two functions of sensors are declared
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}

// Sensor variable is assigned to functions here.
// It is assigned and NOT called (because function calls always have perenthesis () )
// So the reference of these functions are stored in the variable
// and when we use/call the variable, then it invokes the functions it's assigned to
// The perenthesis with sensor() actually indicates invokation of the assigned function

// This way of using functions allows us to determine which functions get invoked in runtime
func main() {
	sensor := fakeSensor
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())

	measureTemperature(3, fakeSensor)

}

// =========================================================
// PASSING FUNCTIONS TO ANOTHER FUNCTIONS

// Here, this function accepts another function as a parameter
// When sensor() is called, it invokes the function passed in as an argument (See line 36)
// So when running this code, sensor() will invoke fakeSensor function

// With this, we can define in runtime which function to pass as an argument which eventually invokes a different function and this is how we control the execution.
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}

// DECLARING FUNCTION TYPES

// type sensor func() kelvin
// This now becomes a custom type sensor functions and can be used as measureTemperature(samples int, s sensor)

// =====================================================
// CLOSURES AND ANONYMOUS FUNCTIONS

// Anonymous function (or function literal) is a function without name.
// They keep reference of the variables within the scope.

/*
var f = func(){
fmt.Println("hello")
}
func main(){
f()
}
*/

// Anonymous functions are useful when we want to create functions on the fly
// Like returning a function form another function - declaring and returning a named function is also possible, but anonymous functions becomes handy
//
