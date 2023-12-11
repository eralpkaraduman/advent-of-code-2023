package day01

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

func Solution() int {
   currentDir, _ := os.Getwd()
   filePath := filepath.Join(currentDir, "01", "input.txt")
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

   var result int = 0
   for scanner.Scan() {
		line := scanner.Text()
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