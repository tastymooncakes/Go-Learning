package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// check if calculator is running with 3 arguements
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		return
	}

	// set arguments to values
	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	// convert string to number
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Please insert valid number")
		return
	}

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Please insert valid number")
		return
	}

	// perform operation
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /")
		return
	}

	fmt.Println("Result:", result)
}