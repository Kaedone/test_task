package main

import (
	"fmt"
)

func main() {
	var operator string
	var num1, num2 float64

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Некорректный оператор")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
}
