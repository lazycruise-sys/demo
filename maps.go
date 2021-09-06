package main

import "fmt"

func anoda(){
	fmt.Println("----anoda----")
}

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

	// full map literal declaration
	myMaps := map[string]float64{"a": 3.4, "b": 3.6, "c": 4.2}
	fmt.Println(myMaps)

	// empty map declaration
	emptyMap := map[string]float64{}
	fmt.Printf("%#v", emptyMap)

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters)
	fmt.Println(counters["b"])

	var nilMap map[string]int
	nilMap = make(map[string]int)
	fmt.Println(nilMap)
	nilMap["three"] = 3
	fmt.Println(nilMap)

	//handling default values in maps
	anoda()
	evidences := map[string]int{"ACD-13": 3, "AER-45": 12}
	acd13, ok := evidences["ACD-13"]
	fmt.Println(acd13, ok)
	dce01, ok := evidences["DCE-01"]
	fmt.Println(dce01, ok)

	evidences["DCE-01"] = 21
	dce01, ok = evidences["DCE-01"]
	fmt.Println(dce01, ok)
	fmt.Println(evidences)

	// using the delete() function to remove a key value and its data value pair
	anoda()
	delete(evidences, "DCE-01")
	dce01, ok = evidences["DCE-01"]
	fmt.Println(dce01, ok)
	fmt.Println(evidences)

}