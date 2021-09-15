package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5

	x := rand.Intn(max-min+1) + min
	fmt.Println(x)

	var first int
	var second int
	var third int

	fmt.Println("I'm thinking of a number between 1 and 5, Enter your guess")

	fmt.Scanln(&first)

	if first != x {
		fmt.Println("Incorrect! 2 more tries")

		fmt.Scanln(&second)

		if second != x {
			fmt.Println("Incorrect! 1 more try")

			fmt.Scanln(&third)

			if third != x {
				fmt.Println("Incorrect! You lose!")
				time.Sleep(2)
			} else {
				fmt.Println("Correct! You win!")
				time.Sleep(2)
			}

		} else {
			fmt.Println("Correct! You win!")
			time.Sleep(2)
		}

	} else {
		fmt.Println("Correct! You win!")
		time.Sleep(2)
	}
}
