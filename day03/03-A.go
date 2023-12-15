package day03

import (
	"strconv"
	"unicode"
)

func SolutionA(input []string) int {
	// input = []string{
	// 	"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	"......755.",
	// 	"...$.*....",
	// 	".664.598..",
	// }

	sum := 0

	for lineIndex, line := range input {
		foundNumberAtIndex := -1
		numberBuffer := ""
		for charIndex, char := range line {
			isDigit := unicode.IsDigit(char)
			isEndOfLine := charIndex == len(line)-1

			if isDigit {
				if foundNumberAtIndex == -1 {
					foundNumberAtIndex = charIndex
				}
				numberBuffer += string(char)
			}

			if !isDigit || isEndOfLine {
				if numberBuffer != "" {
					if checkIfAdjacentToSymbol(input, len(numberBuffer), lineIndex, foundNumberAtIndex) {
						foundNumber, _ := strconv.Atoi(numberBuffer)
						// println("adding", foundNumber)
						sum += foundNumber
					}
				}
				foundNumberAtIndex = -1
				numberBuffer = ""
			}
		}
	}

	return sum
}

func checkIfAdjacentToSymbol(input []string, length int, lineIndex int, indexInLine int) bool {
	valid := false

	lineStartIndex := indexInLine - 1
	lineEndIndex := indexInLine + length + 1
	if lineStartIndex < 0 {
		lineStartIndex = 0
	}
	lineLength := len(input[lineIndex])
	if lineEndIndex >= lineLength {
		lineEndIndex = lineLength - 1
	}

	// check above
	aboveIndex := lineIndex - 1
	if aboveIndex >= 0 && searchLineForSymbols(input[aboveIndex], lineStartIndex, lineEndIndex) {
		valid = true
	}

	// check current line
	if searchLineForSymbols(input[lineIndex], lineStartIndex, lineEndIndex) {
		valid = true
	}

	// check below
	belowIndex := lineIndex + 1
	if (belowIndex < len(input)) && searchLineForSymbols(input[belowIndex], lineStartIndex, lineEndIndex) {
		valid = true
	}

	// println(input[lineIndex][indexInLine:indexInLine+length], valid)

	return valid
}

func searchLineForSymbols(line string, startIndex int, endIndex int) bool {
	hasSymbol := false

	for i := startIndex; i < endIndex; i++ {
		char := []rune(line)[i]
		if string(char) != "." && !unicode.IsDigit(char) {
			hasSymbol = true
		}
	}

	// println(line[startIndex:endIndex], "->", hasSymbol)
	return hasSymbol
}
