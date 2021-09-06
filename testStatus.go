package main

import (
	"fmt"
)

func status(name string) {
	results := map[string]float64{"andy": 34, "carlos": 86}
	grade, ok := results[name]
	if !ok {
		fmt.Printf("%s results have not been added.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing.\n", name)
	}
}

func main() {
	status("andy")
	status("Andy")
	status("eden")
}
