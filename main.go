package main

import (
	"fmt"
	"strconv"
	"strings"
)

var NUMBERS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
}

var HUNDREDS = map[string]int{
	"hundred":  100,
	"thousand": 1000,
	"million":  1000000,
}

func getNumber(word string) int {
	number, ok := NUMBERS[strings.ToLower(word)]
	if !ok {
		return 0
	}
	return number
}

func getHundredAmount(word string) int {
	amount, ok := HUNDREDS[strings.ToLower(word)]
	if !ok {
		return 0
	}
	return amount
}

// setPointsToNumber принимает число и возвращает строку с запятыми для разделения тысяч
func setPointsToNumber(num int) string {
	numStr := strconv.Itoa(num)

	var result strings.Builder
	length := len(numStr)

	for i, digit := range numStr {
		result.WriteRune(digit)

		if (length-i-1)%3 == 0 && i != length-1 {
			result.WriteString(",")
		}
	}

	return result.String()
}

func getNumberFromString(text string, usePoint bool) string {
	words := strings.Fields(text)
	var total int

	for _, word := range words {
		number := getNumber(word)
		if number > 0 {
			total += number
		} else {
			amount := getHundredAmount(word)
			if amount > 0 {
				total *= amount
			}
		}
	}

	if usePoint {
		return setPointsToNumber(total)
	} else {
		return strconv.Itoa(total)
	}
}

func main() {
	NEEDED_STRING := "Five hundred million"
	result := getNumberFromString(NEEDED_STRING, true)

	fmt.Println(result) // 500,000,000
}
