package main

import (
	"fmt"
)

// creating subscriber as a defined type
type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var myStruct struct {
		number float64
		word   string
		toogle bool
	}

	fmt.Printf("%#v\n", myStruct)
	fmt.Println(myStruct.number)
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toogle = true
	fmt.Println(myStruct)

	// subscriber/customer data
	var customer struct {
		name   string
		rate   float64
		active bool
	}

	// assigning customer data
	customer.name = "Wayne Rooney"
	customer.rate = 4.99
	customer.active = true

	// printing customer data
	fmt.Println("Name:", customer.name)
	fmt.Println("Monthly Rate:", customer.rate)
	fmt.Println("Active?", customer.active)

	// defined types and structs
	// type definitions allows one to create (defined) types of their own, which are based on
	// an underlying type
	type myType struct {
		name   string
		rate   float64
		active bool
	}

	// declaring a variable based on the defined subscriber type
	// this helps with repititive declaration of struct literal and multiple use of similar data structure
	var subscriber1 subscriber
	subscriber1.name = "olorunfemi"
	fmt.Println(subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "jacob"
	fmt.Println(subscriber2.name)
}
