package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// comments serve as the pseudocode breakdown

	// 01. generate a random number from 1 to 100 and store it as a target number
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(101)

	// 02. prompt the player to guess what the target number is, and store their response.
	fmt.Println("we have selected a number between 0 and 100")
	fmt.Println("can you guess the number?")
	fmt.Println(target) // for debugging purpose

	reader := bufio.NewReader(os.Stdin) // lets us read keyboard input (creates a bufio.Reader value)
	fmt.Println("Make your guess?")
	input, err := reader.ReadString('\n') // reads what the user types up until they press Enter
	if err != nil {
		log.Fatal(err) // closes program if an erro exists
	}

	input = strings.TrimSpace(input)  // removes newline(\n)
	guess, err := strconv.Atoi(input) // converts string to integer
	if err != nil {
		log.Fatal(err)
	}

	// 03. if player's guess is greater or less than target indicate to the user.

	if guess < target {
		fmt.Println("oops! your guess was LOW")
	} else if guess > target {
		fmt.Println("oops! your guess was HIGH")
	}

	// 04. allow the player to guess up to 10 times and indicate the number of guesses left

	// 05. if player's guess is correct, indicate to them and stop the loop

	// 06. if player runs out of turns with no correct guesses, indicate to the user and show the correct number

}
