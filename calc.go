package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var romanNums = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func toRoman(num int) string {
	if num <= 0 || num > 4000 {
		return ""
	}

	roman := ""

	for num > 0 {
		if num >= 1000 {
			roman += "M"
			num -= 1000
		} else if num >= 900 {
			roman += "CM"
			num -= 900
		} else if num >= 500 {
			roman += "D"
			num -= 500
		} else if num >= 400 {
			roman += "CD"
			num -= 400
		} else if num >= 100 {
			roman += "C"
			num -= 100
		} else if num >= 90 {
			roman += "XC"
			num -= 90
		} else if num >= 50 {
			roman += "L"
			num -= 50
		} else if num >= 40 {
			roman += "XL"
			num -= 40
		} else if num >= 10 {
			roman += "X"
			num -= 10
		} else if num >= 9 {
			roman += "IX"
			num -= 9
		} else if num >= 5 {
			roman += "V"
			num -= 5
		} else if num >= 4 {
			roman += "IV"
			num -= 4
		} else {
			roman += "I"
			num--
		}
	}

	return roman
}

func toArabic(roman string) (int, error) {
	var total int
	prev := 0

	for i := len(roman) - 1; i >= 0; i-- {
		val := romanNums[string(roman[i])]

		if val < prev {
			total -= val
		} else {
			total += val
		}

		prev = val
	}

	return total, nil
}

func main() {
	var input string
	fmt.Println("Введите выражение:")
	fmt.Scanln(&input)

	// Проверяем, содержит ли выражение римские цифры
	match, _ := regexp.MatchString("[IVXLCDM]+", input)
	isRoman := match

	// Разбиваем входную строку на отдельные элементы (числа и операторы)
	elements := strings.Fields(input)

	// Преобразуем каждый элемент в число или оператор
	numbers := []float64{}
	operators := []string{}

	for _, element := range elements {
		if isRoman {
			val, err := toArabic(element)
			if err == nil {
				numbers = append(numbers, float64(val))
			} else {
				operators = append(operators, element)
			}
		} else {
			if val, err := strconv.ParseFloat(element, 64); err == nil {
				numbers = append(numbers, val)
			} else {
				operators = append(operators, element)
			}
		}
	}
	
	// Выполняем операции, начиная с самых приоритетных (умножение и деление)
for i := 0; i < len(operators); i++ {
	if operators[i] == "*" {
		numbers[i] = numbers[i] * numbers[i+1]
		numbers = append(numbers[:i+1], numbers[i+2:]...)
		operators = append(operators[:i], operators[i+1:]...)
		i--
	} else if operators[i] == "/" {
		numbers[i] = numbers[i] / numbers[i+1]
		numbers = append(numbers[:i+1], numbers[i+2:]...)
		operators = append(operators[:i], operators[i+1:]...)
		i--
	}
}

// Выполняем оставшиеся операции (сложение и вычитание)
for i := 0; i < len(operators); i++ {
	if operators[i] == "+" {
		numbers[i] = numbers[i] + numbers[i+1]
		numbers = append(numbers[:i+1], numbers[i+2:]...)
		operators = append(operators[:i], operators[i+1:]...)
		i--
	} else if operators[i] == "-" {
		numbers[i] = numbers[i] - numbers[i+1]
		numbers = append(numbers[:i+1], numbers[i+2:]...)
		operators = append(operators[:i], operators[i+1:]...)
		i--
	}
}

// Выводим результат
if isRoman {
	result := toRoman(int(numbers[0]))
	fmt.Println(result)
} else {
	fmt.Println(numbers[0])
}

