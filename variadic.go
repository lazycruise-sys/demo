package main

import (
	"fmt"
	"math"
)

func stringInts(data string, numbers ...int) {
	fmt.Printf("%s, %d\n", data, numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func myAverage(nums ...int) string {
	var sum float64
	length := float64(len(nums))
	for _, num := range nums {
		sum += float64(num)
	}
	result := fmt.Sprintf("%.3f", sum/length)
	return result
}

func myMax(nums ...float64) float64 {
	// var max int
	max := math.Inf(-1)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var rnge []float64
	for _, num := range numbers {
		if num > min && num < max {
			rnge = append(rnge, num)
		}
	}
	return rnge
}

func main() {
	stringInts("man", 2)
	stringInts("man", 2, 3, 40)
	stringInts("man")
	mix(2, true, "ad", "donda", "clb", "new")
	fmt.Println(myAverage(2,2,2,2,2,2,4,23,23))
	fmt.Println(myMax(-2,3,1,67,23))
	fmt.Println(myMax(-2,-3,-67,-23))
	fmt.Println(inRange(0,10,4,3,32,2,1,1,2,3,764,2))
}