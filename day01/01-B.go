package day01

import (
	"strings"
	"unicode"
)

var numberWords = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func SolutionB(input []string) int {
	sum := 0
	
	for _, line := range input {
		forwardBuffer := ""
		reverseBuffer := ""
		
		first := -1
		last :=	-1
		
		for i, char := range line {
			forwardBuffer += string(char)
			if forwardFound := findNumber(forwardBuffer); forwardFound != -1 && first == -1 {
				first =	forwardFound
			}
			
			reverseBuffer = string(line[len(line) - 1 - i]) + reverseBuffer
			if reverseFound := findNumber(reverseBuffer); reverseFound != -1 && last == -1 {
				last = reverseFound
			}
		}
		
		num := ( first * 10 ) + last 
		sum += num
	}
	return sum
}

func findNumber(buffer string) int {
	for _, char := range buffer {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	
	for index, numberWord := range numberWords {
		if strings.Contains(buffer, numberWord) {
			return index + 1
		}
	}
	
	return -1
}