package main

import "fmt"

func main() {
	var total_numbers int
	number1 := 0
	number2 := 1
	next_number := 0

	fmt.Print("Enter how many numbers of the Fibonacci Sequence you want to see: ")
	fmt.Scan(&total_numbers)
	fmt.Print("This is the sequence: ")
	for i := 1; i <= total_numbers; i++ {
		if i == 1 {
			fmt.Print(" ", number1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", number2)
			continue
		}
		next_number = number1 + number2
		number1 = number2
		number2 = next_number
		fmt.Print(" ", next_number)
	}
}
