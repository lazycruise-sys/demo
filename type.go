package main

import (
	"fmt"
	"reflect"
)

// defined type declaration
type Liters float64
type Gallons float64
type Car struct {
	Name       string
	TopSpeed   float64
	EngineType string
}

func main() {
	// initialization of fuel variable
	var fuel float64 = 10 // this doesn't show  or tell whether it is in liter or gallons
	fmt.Println(fuel)

	// using defined types, liter and gallon
	var carFuel Gallons                                                                // defining a variable with the type Gallons
	var busFuel Liters                                                                 // defining a variable with the type Liters
	var FisayoCar Car                                                                  // defining a variable with the type Car
	carFuel = Gallons(10.0)                                                            // converting a float64 to Gallon
	busFuel = Liters(240.0)                                                            // converting a float64 to Liters
	FisayoCar = Car{Name: "Mercedes Maybach", TopSpeed: 345.67, EngineType: "4-wheel"} // declaring the fields of FisayoCar
	fmt.Println(carFuel, busFuel, FisayoCar)                                           // printing variables value

	trainFuel := Liters(4560.0)            // short variable declarations together with type conversion
	fmt.Println(trainFuel)                 // printing variable trainFuel
	fmt.Println(reflect.TypeOf(trainFuel)) // checking type

	// golang only considers the value of the underlying type of float64
	// therefore, there is no difference between Gallons(Liters(240.0))
	// and Gallons(240.0)
	testFuel := Gallons(13.45)
	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(testFuel)

	fmt.Println(carFuel, busFuel)
	fmt.Println(Gallons(63.0), reflect.TypeOf(Gallons(63.0)))

}
