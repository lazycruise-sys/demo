// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Welcome to Golang"
	var nameTwo = "Another welcome to Golang"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameTwo = "Peaches"
	nameThree = "Titties"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yonda"

	fmt.Println(nameFour)

	// integers
	var ageFour = 40
	var ageFive int = 45
	ageSix := 60

	fmt.Println(ageFour, ageFive, ageSix)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -127
	var numThree uint8 = 255

	//floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 37891013910.76

	fmt.Println(scoreOne, scoreTwo)

	age := 35
	name := "Femi"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("New Line\n")
	fmt.Printf("My first, second and third numbers are %v, %v, and %v respectively\n", numOne, numTwo, numThree)

	// Println
	fmt.Println("Great Things are beginning")
	fmt.Println("My Name is", name, "and I'm", age, "years of age.")

	// Printf
	fmt.Printf("My name is %v and I'm %v years of age\n", name, age) // useful for variables
	fmt.Printf("My name is %q and I'm %q years of age\n", name, age) // useful for introducing quotations but not with integers
	fmt.Printf("My age variable holds a number of type %T\n", age)
	fmt.Printf("You scored %f points\n", 23.456)
	fmt.Printf("You scored %0.2f points\n", 23.456) // useful for limiting decimal points.

	// Sprintf
	var str = fmt.Sprintf("My name is %v and I'm %v years of age", name, age)
	fmt.Println(str)

}
