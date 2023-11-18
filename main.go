package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	min, max := 1, 100
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretNumber := r.Intn(max - min)

	fmt.Println("Guess a number between 1 and 100 ", secretNumber)

	attempts := 0
	for {

		fmt.Print("Please input your guess: ")
		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		attempts++

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts")
			break
		}
	}
}

// to do
