package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 3.45
	return &myFloat
}

func printPointer(boolPointer *bool) {
	fmt.Println(*boolPointer)
}

func main() {
	var myInt int = 3 // initialize a variable of integer value 3
	var myIntPointer *int // initialize a variable of integer pointer
	myIntPointer = &myInt // assign the address value of the myInt value (pointer)
	fmt.Println(myIntPointer) // print the pointer
	fmt.Println(*myIntPointer) // print the value at the pointer
	*myIntPointer = 7 // change the value of the myInt variable at the pointer

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	myBool = true
	printPointer(&myBool)
}