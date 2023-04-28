// Created by: Janet
// Created on: Sep 2020
//
// The random number guessing game

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Generate a random number between 1 and 6
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}

	// Ask the user to enter a number between 1 and 6
	fmt.Println("Guess a number between 1 and 6:")
	var userNumber int
	fmt.Scanln(&userNumber)

	// Check if the user's guess is correct
	if userNumber == int(randomNumber.Int64())+1 {
		fmt.Println("Congratulations! You guessed correctly.")
	} else {
		fmt.Println("Sorry, you guessed wrong. The correct number was", randomNumber.Int64()+1)
	}
}
