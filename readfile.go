package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the data file for reading
	file, err := os.Open("data.txt") // creates os.File value and error type, PathError
	if err != nil {                  // if there was an error openining the file, report it and exit
		log.Fatal(err)
	}

	// create a new Scanner for the file
	scanner := bufio.NewScanner(file) // creates bufio.Scanner value
	for scanner.Scan() {              // read a line from the file by looping until the end of the file is reached and scanner.Scan returns false
		fmt.Println(scanner.Text()) // prints the read line
	}

	// close the file to free resources
	err = file.Close()
	if err != nil { // if there was an error closing the file, report it and exit
		log.Fatal(err)
	}

	// if there was an error scanning the file, report it and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
