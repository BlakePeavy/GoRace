package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to move the player's car forward
func moveForward(position *int) {
	// Generate a random number between 1 and 10
	distance := rand.Intn(10) + 1

	// Move the player's car forward by the generated distance
	*position += distance

	fmt.Printf("Your car moved forward by %d units\n", distance)
}

// Function to move the opponent's car forward
func moveOpponent(position *int) {
	// Generate a random number between 1 and 10
	distance := rand.Intn(10) + 1

	// Move the opponent's car forward by the generated distance
	*position += distance

	fmt.Printf("The opponent's car moved forward by %d units\n", distance)
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize the player's and opponent's car positions
	playerPosition := 0
	opponentPosition := 0

	// The game loop
	for {
		// Print the current car positions
		fmt.Printf("Your position: %d\n", playerPosition)
		fmt.Printf("Opponent's position: %d\n", opponentPosition)

		// Check if players are tied
		if playerPosition == 100 && opponentPosition == 100 {
			fmt.Println("You tied!")
		}

		// Check if the player's car has reached the finish line
		if playerPosition >= 100 {
			fmt.Println("You won the race!")
			break
		}

		// Check if the opponent's car has reached the finish line
		if opponentPosition >= 100 {
			fmt.Println("You lost the race!")
			break
		}

		// Move the player's car forward
		moveForward(&playerPosition)

		// Move the opponent's car forward
		moveOpponent(&opponentPosition)
	}
}
