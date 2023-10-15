package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}

	return roman
}

func Arabic(number string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	total := 0
	prevValue := 0

	for _, numeral := range number {
		value, found := romanNumerals[numeral]
		if !found {
			fmt.Println("Ошибка: введены неверные римские числа.")
		}

		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}

		prevValue = value
	}

	return total

}
func main() {
	var result int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	equation := scanner.Text()
	array := strings.Split(equation, " ")
	if len(array) != 3 {
		log.Fatal("формат математической операции не удовлетворяет заданию ")
	}

	isArabic := false
	isRoman := false

	a, err1 := strconv.Atoi(array[0])
	if err1 != nil {
		a = Arabic(array[0])
		isRoman = true
	} else {
		isArabic = true
	}

	b, err2 := strconv.Atoi(array[2])
	if err2 != nil {
		b = Arabic(array[2])
		isRoman = true
	} else {
		isArabic = true
	}

	if a < 1 || a > 10 && b < 1 || b > 10 {
		log.Fatal("Можно использовать только числа от 1 до 10")
	}

	if isArabic && isRoman {
		log.Fatal("Используются одновременно разные системы счисления")
	}

	switch array[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		log.Fatal("Неверное выражение")
	}

	if isRoman == true {
		fmt.Println(Roman(result))
		if result < 1 {
			log.Fatal("В римской системе нет отрицательных чисел")
		}
	} else {
		fmt.Println(result)
	}

}
