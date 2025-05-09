// Unlike other programming languages, GO does not have classes and objects.
// Methods can be attached to other types the same way as float64 or struct.
// GO does not have special constuctor declaration like Python, Java and other OOP
// But it follows a convention to allow to define constuctor like functions.

package main

import "fmt"

type location struct {
	lat, long float64
}

type coordinates struct {
	d, m, s float64
}

func (c coordinates) decimal() float64 {
	return c.d + c.m/60 + c.s/3600
}

// Functions in newType or NewType convention are usually constructor of a value for the given type. Such convention is better for readability and function calls.
func newLocation(lat, long coordinates) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	lat := coordinates{d: 37, m: 47, s: 0}
	long := coordinates{d: 122, m: 25, s: 0}
	loc := newLocation(lat, long)
	fmt.Println(loc)
}
