package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("Welcome, I have though of a number between 0 and 100, can you guess what it is?")
	var answer int = r1.Intn(100)
	fmt.Println("Have a guess")
	game(answer)
}

func game(answer int) {
	var counter int = 1
	for {
		var input int

		fmt.Scan(&input)
		switch {
		case input < answer:
			counter++
			fmt.Println("That number is too low!")
			if answer-input <= 5 {
				fmt.Println("getting warmer")
			}
			if answer-input >= 20 {
				fmt.Println("you're nowhere near!!")
			}

		case input > answer:
			counter++
			fmt.Println("That number is too high!")
			if input-answer <= 5 {
				fmt.Println("getting warmer")
			}
			if input-answer >= 20 {
				fmt.Println("you're nowhere near!!")
			}

		case input == answer:
			{
				fmt.Println("That's it! You read my mind!")
				fmt.Printf("Congratulations, it only took you %d tries", counter)
				fmt.Println("\nWould you like to play again?")
				var yesno string
				fmt.Scanln(&yesno)
				switch yesno {
				case "yes":
					fmt.Println("Okay, starting again")
					main()
				case "no":
					fmt.Println("Exiting")
					os.Exit(0)
				}
			}
		}

	}

}
