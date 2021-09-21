package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {

	container := [2]int{} // holds the two highest values

	// temp variable to hold new highest
	measure1, measure2 := 0, 0

	// loop system
	for _, y := range ages {
		if y > measure1 {
			measure2, measure1 = measure1, y
		} else if y > measure2 {
			measure2 = y
		}
	}

	// assigning the highest into the array
	container[0], container[1] = measure2, measure1

	return container
}

// import "sort"

// func TwoOldestAges(ages []int) [2]int {
//   sort.Sort(sort.Reverse(sort.IntSlice(ages)))
//   return [2]int{ages[1],ages[0]}
// }

func main() {
	testArray := []int{98, 5, 87, 45, 8, 108}
	fmt.Println(TwoOldestAges(testArray))
}
