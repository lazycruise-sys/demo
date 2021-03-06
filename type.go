package main

import (
	"fmt"
	"reflect"
)

// defined type declaration
type Liters float64
type Gallons float64
type Milliliters float64
type Title string
type Car struct {
	Name       string
	TopSpeed   float64
	EngineType string
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func main() {
	// initialization of fuel variable
	var fuel float64 = 10 // this doesn't show  or tell whether it is in liter or gallonsGallons
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

	// converting from liters to gallons
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("gallons: %.1f, liters: %.1f\n", carFuel, busFuel)

	// defined types and operators
	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(3.4))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(3.4))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	fmt.Println(Title("allen") == Title("allen"))
	fmt.Println(Title("allen") > Title("zodiac"))

	// defined types can be used in operations
	// together with literal values
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	checkCalcType := Gallons(5.6) + 2.4 // has a main.Gallons type
	fmt.Println(reflect.TypeOf(checkCalcType))
	fmt.Println(Gallons(1.2) == 1.2)
	// however, they can't be used in operations
	// with values of different types such as
	// Gallons(1.2) + Liter(23.4)
	// Gallons(2.3) > Liters(2.2)


	// converting between types using functions
	carFuel, busFuel= Gallons(1.2), Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("car fuel: %.2f gallons.\n", carFuel)
	fmt.Printf("bus fuel: %.2f liters.\n", busFuel)
}
