package day01

import (
	"strconv"
	"unicode"
)

func SolutionA(input []string) int {
	var result int = 0
	for _, line := range input {
		pair := findFirstAndLast(line)
		num, _ := strconv.Atoi(pair)
		result += num
	}
	return result
}

func findFirstAndLast(code string) (string) {
	var first, last string
	for _, char := range code {
		if unicode.IsDigit(char) {
			if first == "" {
				first = string(char)
			}
			last = string(char)
		}
	}
	return first + last
}


