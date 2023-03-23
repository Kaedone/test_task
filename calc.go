package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите выражение:")
	fmt.Scanln(&input)

	// Разбиваем ввод на числа и операторы
	elements := strings.Fields(input)

	// Преобразуем
	numbers := []float64{}
	operators := []string{}

	for _, element := range elements {
		if val, err := strconv.ParseFloat(element, 64); err == nil {
			numbers = append(numbers, val)
		} else {
			operators = append(operators, element)
		}
	}

	// Выполняем вычисления (работает порядок операций)
	for _, operator := range []string{"*", "/", "+", "-"} {
		for i := 0; i < len(operators); i++ {
			if operators[i] == operator {
				result := 0.0

				switch operator {
				case "+":
					result = numbers[i] + numbers[i+1]
				case "-":
					result = numbers[i] - numbers[i+1]
				case "*":
					result = numbers[i] * numbers[i+1]
				case "/":
					result = numbers[i] / numbers[i+1]
				}

				numbers[i] = result
				numbers = append(numbers[:i+1], numbers[i+2:]...)
				operators = append(operators[:i], operators[i+1:]...)
				i--
			}
		}
	}

	// Вывод результата
	fmt.Printf("Результат: %.2f", numbers[0])
}
