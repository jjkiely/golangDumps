package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter 2 numbers")
	var s1 int
	fmt.Scanln(&s1)
	var s2 int
	fmt.Scanln(&s2)

	fmt.Println(multiply(s1, s2))
}
func multiply(s1 int, s2 int) int {
	return s1 * s2
}
