package main

import (
	"fmt"
)

func main() {
	//this programme will print a star sequence up to the given number
	fmt.Println("Please enter the number you want the sequence to stop at")
	var a int
	fmt.Scanln(&a)
	fmt.Println(fib(a))

}

func fib(a int) string {
	if a == 1 {
		return "*"
	}
	if a == 2 {
		return "**"
	} else {
		return "*" + fib(a-1)
	}

}
