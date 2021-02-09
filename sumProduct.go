//This simple application will continue looping until the user exits, computing products or sums of number sequences based on user input
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Welcome, please enter a number, enter 0 to quit")
		var instruction int
		fmt.Scan(&instruction)
		if instruction == 0 {
			fmt.Println("Exiting")
			os.Exit(0)
		}
		fmt.Println("Would you like to compute the sum or product of this number sequence?")
		var sumorproduct string
		fmt.Scanln(&sumorproduct)
		switch sumorproduct {
		case "sum":
			time.Sleep(time.Second)
			fmt.Printf("The sum of all numbers between 1 and %d is %d", instruction, sum(instruction))

		case "product":
			if instruction > 20 {
				time.Sleep(time.Second)
				fmt.Println("Sorry that number is too big to handle, returning to beginning")
				//fmt.Scan(&instruction)
				main()
			}
			fmt.Printf("The product of all numbers between 1 and %d is %d", instruction, product(instruction))

		}
		fmt.Println("\nWould you like to exit?")
		var exitcode string
		fmt.Scanln(&exitcode)
		switch exitcode {
		case "yes":
			fmt.Println("Exiting")
			os.Exit(0)
		case "no":
			fmt.Println("Returning to beginning of program")
			main()
		}

	}

}

func sum(instruction int) int {
	var result int
	for i := 1; i <= instruction; i++ {
		result += i
	}
	return result
}

func product(instruction int) int64 {

	/*	var result2 int64
		result2 = int64(1)
		for i := 1; i <= instruction; i++ {
			result2 += result2 * int64(i)
		}
		return result2*/
	if instruction == 1 {
		return 1

	}
	if instruction > 1 {
		return int64(instruction) * product(instruction-1)
	}
	return 0

}
