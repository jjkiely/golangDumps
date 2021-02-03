//This application will allow the user to create and read from a simple to do list, for the moment it is not persistent, however I may fork it and make it persistent in the future
//I do not currently know how to scan in multi word strings so for now, tasks must be entered without spaces
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome, please select an option")
	fmt.Println("Press A to add an event to your calendar")
	fmt.Println("Press V to view your current events")
	fmt.Println("Press Q to exit the application")
	fmt.Println("To see these options again at any time, press H")

	var list string
	for {
		var instruction string
		fmt.Scanln(&instruction)
		switch instruction {
		case "A":
			fmt.Println("Please enter the task you wish to add, without spaces")
			var temp string
			fmt.Scanln(&temp)
			list += temp
			fmt.Println("Thanks, your task has been added, returning to main menu")
			fmt.Println("PLease Enter a Command")

		case "V":
			fmt.Println(list)
			fmt.Println("returning to main menu")

		case "Q":
			os.Exit(0)
		default:
			fmt.Println("sorry that's not valid")

		}
	}
}
