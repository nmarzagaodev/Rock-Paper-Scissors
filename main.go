package main

import (
	"fmt"
	"math/rand"
)

// ask for player option
// get computer random option
// keep track of score

const (
	ROCK     = iota
	PAPER    = iota
	SCISSORS = iota
)

type Score struct {
	Wins   int
	Losses int
	Ties   int
}


func GetPlayerInput() int {
	var selection int

	for {
		fmt.Println("[1] - Rock, [2] - Paper, [3] - Scissors")
		fmt.Scan(&selection)

		if selection > 0 && selection < 4 {
			break
		}
	}
	
	return selection - 1
}

func GetComputerChoice() int {
	return rand.Intn(3) // 0 to 2 
}


func main() {
	test := GetPlayerInput();
	fmt.Printf("The choice was %d\n", test)
	fmt.Printf("The computer choice is %d\n", GetComputerChoice())
}
