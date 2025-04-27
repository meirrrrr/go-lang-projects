package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Invalid input. Exiting program.")
		os.Exit(1)
	}

	fmt.Print("Enter an operator (+, -, *, /, ^ for exponent, sqrt for square root): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid input. Exiting program.")
		os.Exit(1)
	}

	if operator == "sqrt" {
		result := math.Sqrt(num1)
		fmt.Printf("The square root of %.2f is %.2f\n", num1, result)
		return
	}

	if operator != "sqrt" {
		fmt.Print("Enter the second number: ")
		_, err := fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("Invalid input. Exiting program.")
			os.Exit(1)
		}
	}

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	case "^":
		result := math.Pow(num1, num2)
		fmt.Printf("%.2f ^ %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Invalid operator. Please use one of the following: +, -, *, /, ^, sqrt.")
	}
}
