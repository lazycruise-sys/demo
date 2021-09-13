package main

import "fmt"

// define a type named part
type part struct {
	description string
	count       int
}

// define a type named car
type car struct {
	name     string
	topSpeed float64
}

// function declaration using part type
func showInfo(p part) {
	fmt.Println("description: ", p.description)
	fmt.Println("count: ", p.count)
}

// function declaration creating a part type
func minimumOrder(description string) part {
	var p part
	if description == "hex bolts" {
		p.count = 150
		p.description = description
		return p
	} else {
		p.count = 100
		p.description = description
		return p
	}
}

func main() {
	// declare a variable of type "car"
	var porsche car
	porsche.name = "porsche 911 r"
	porsche.topSpeed = 323
	fmt.Println("name: ", porsche.name)
	fmt.Println("top speed", porsche.topSpeed)

	var bolt part
	bolt.description = "hex bolts"
	bolt.count = 24
	// fmt.Println("part description: ", bolt.description)
	// fmt.Println("count: ", bolt.count)
	showInfo(bolt)

	// determining the minimum count of hex bolts and helical screw
	hexBolts := minimumOrder("hex bolts")
	helicalScrews := minimumOrder("helical screw")
	fmt.Println(hexBolts.count)
	fmt.Println(helicalScrews.count)
}
