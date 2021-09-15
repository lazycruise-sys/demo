package main

import (
	"fmt"
)

// declaring subscriber type
type subscriber struct {
	name   string
	rate   float64
	active bool
}

// creating an example struct
type myStruct struct {
	myField int
}

// function to create inital subscriber type
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

// function to print out subscriber data
func showSubscriberData(s *subscriber) {
	fmt.Println("subscriber name: ", s.name)
	fmt.Println("subscriber monthly rate: ", s.rate)
	fmt.Println("subscriber status: ", s.active)
}

// accessing the pointer value of the subscriber struct to make changes to it.
func applyDiscount(s *subscriber) {
	// (*s).rate = 3.99 // this is a tedious process
	s.rate = 3.99
}

// multplying a variable value at the pointer level
func double(number *int) {
	*number *= 2
}

// main function
func main() {
	// subscriber 1 data
	subscriber1 := defaultSubscriber("amit sangh") // this is not a struct, but rather a struct pointer
	subscriber1.rate = 4.99
	// applyDiscount(subscriber1)
	subscriber2 := defaultSubscriber("bruce banner")
	showSubscriberData(subscriber1)
	showSubscriberData(subscriber2)

	// subscriber dat update
	applyDiscount(subscriber1) // since already a struct pointer
	fmt.Println(subscriber1.rate)
	// testing pointer knowledge
	amount := 6
	double(&amount)
	fmt.Println(amount)
	pointToAmount := &amount    // pointer variable declaration
	fmt.Println(pointToAmount)  // prints the pointer (memory address)
	fmt.Println(*pointToAmount) // prints the value-at the pointer

	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println(pointer.myField)
	pointer.myField = 9
	fmt.Println(pointer.myField)
}
