package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	index := len(ages)
	container := [2]int{ages[index-2], ages[index-1]}
	return container
}

func main() {
	testArray := []int{98, 5, 87, 45, 8, 108}
	fmt.Println(TwoOldestAges(testArray))
}
