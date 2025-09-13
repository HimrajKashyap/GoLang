package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate a random numberbetween 1 to 100
	target := random.Intn(100) + 1

	//Welcome Message
	fmt.Println("Welcome to the Guessing game!!")
	fmt.Println("I have chosen a number from 1-100")
	fmt.Println("Can you guess it?")

	var v int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&v)

		if v == target {
			fmt.Println("Correct!@!")
			break
		} else if v < target {
			fmt.Println("Too Low")
		} else {
			fmt.Println("Too High")
		}

	}

}
