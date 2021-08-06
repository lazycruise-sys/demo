package main

import "fmt"

func main() {

	// array
	// var ages [3]int = [3]int {3, 45, 6}

	var ages = [3]int {3, 45, 6}

	names := [4]string {"Yonda", "Ye", "The Life of Pablo", "Black Skinhead"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(ages))

	names [1] = "Burna Boy"

	fmt.Println(names)

	// slices
	scores := []int{23, 45, 6, 78, 89, 100}
	scores[2] = 56

	scores = append(scores, 23)

	fmt.Println(scores)

	// ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := scores[:4]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}