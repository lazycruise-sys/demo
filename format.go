package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%f liters needed\n", 1.8199999999999999999)

	// the value 12 is the minimum width of the string and 2 is the minimum width of the integer (and also serves as padding for the value)
	fmt.Println()
	fmt.Println("----------------NEW CODE---------------")
	fmt.Println()
	fmt.Printf("%12s | %s\n", "products", "cost in cents")
	fmt.Println("-------------------------------")
	fmt.Printf("%12s | %2d\n", "stamps", 50)
	fmt.Printf("%12s | %2d\n", "paper clips", 5)
	fmt.Printf("%12s | %2d\n", "tape", 99)

	// the number 7 serves as the minimum width of the entire number (often the padding) and the 1,2,3 serves as width after the decimal point
	fmt.Println()
	fmt.Println("----------------NEW CODE---------------")
	fmt.Println()
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456) // rounded to three places (the definition of the width) , with 7 digit padding
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456) // rounded to two places, with 7 digit padding
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456) // rounded to one place, with 7 digit padding
	fmt.Printf("%%.1f: %.1f\n", 12.3456)   // rounded to one place, no padding
	fmt.Printf("%%.2f: %.2f\n", 12.3456)   // rounded to two places, no padding (2 decimal points)
	fmt.Println()
	fmt.Println("----------------NEW CODE---------------")
	fmt.Println()
	fmt.Printf("%.2f\n", 1.2600000000000000000)
}
