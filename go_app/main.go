// Created by: Janet
// Created on: Sep 2020
//
// The random number guessing game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	randomNumber := rand.Intn(6) + 1

	// Ask the user to enter a number between 1 and 6
	fmt.Println("Guess a number between 1 and 6:")
	var userNumber int
	fmt.Scanln(&userNumber)

	// Check if the user's guess is correct
	if userNumber == randomNumber {
		fmt.Println("Congratulations! You guessed correctly.")
	} else {
		fmt.Println("Sorry, you guessed wrong. The correct number was", randomNumber)
	}
}
