// The default types not necessarily fulfill unique requirements and in such cases, new types can be declared
package main

// Here a new type named celcius is defined with underlying type of float64
type celcius float64

// The variable temperature uses the new type celcius which behind the scene uses float64
var temprature celcius = 20

// New type here does not mean alias, so any operations on this type must require explicit converstion
// New and unique type better describe the use-case and enhance readability to avoid silly mistakes

// Unlike conventional OOP languages, Go does not have classes and objects
