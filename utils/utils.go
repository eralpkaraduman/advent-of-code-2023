package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(path string) ([]string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current directory: %w", err)
	}

	filePath := filepath.Join(currentDir, path, "input.txt")
	fmt.Println("Reading file:", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}