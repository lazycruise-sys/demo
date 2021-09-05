package main

import "fmt"

func main() {
	var ranks map[string]float64
	ranks = make(map[string]float64) // or ranks := make(map[string]float64) is enough
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["gold"])
	fmt.Println(ranks)

	elements := make(map[string]string)
	elements["H"] = "hydrogen"
	elements["He"] = "helium"
	elements["O"] = "oxygen"
	elements["Cl"] = "chlorine"
	fmt.Println(elements)

	shows := make(map[bool][]string)
	shows[true] = []string{"the blacklist", "bosch"}
	shows[false] = []string{"game of thrones", "prison break"}
	fmt.Println(shows[true][1])
}