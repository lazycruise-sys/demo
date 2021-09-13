package main

import (
	"fmt"
)

// declaring subscriber type
type subscriber struct {
	name string
	rate float64
	active bool
}

// function to create inital subscriber type
func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

// function to print out subscriber data
func showSubscriberData(s subscriber) {
	fmt.Println("subscriber name: ", s.name)
	fmt.Println("subscriber monthly rate: ", s.rate)
	fmt.Println("subscriber status: ", s.active)
}

// main function
func main() {
	// subscriber 1 data
	subscriber1 := defaultSubscriber("amit sangh")
	subscriber1.rate = 4.99
	subscriber2 := defaultSubscriber("bruce banner")
	showSubscriberData(subscriber1)
	showSubscriberData(subscriber2)
}

