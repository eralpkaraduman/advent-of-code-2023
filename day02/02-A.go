package day02

import (
	"strconv"
	"strings"
)

var test = []string {
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

type Game struct {
	ID    int
	Sets  []ColorSet
}

type ColorSet struct {
	Red   int
	Green int
	Blue  int
}

func NewColorSet() ColorSet {
	return ColorSet{
		Red: 0,
		Green: 0,
		Blue: 0,
	}
}

var maxColors = ColorSet{
	Red: 12,
	Green: 13,
	Blue: 14,
}


func SolutionA(input []string) int {
	// input = test

	sumPossibleIds := 0

	for _, line := range input {
		game := parseGame(line)
		possible := true
		for _, set := range game.Sets {
			if set.Red > maxColors.Red || set.Green > maxColors.Green || set.Blue > maxColors.Blue {
				possible = false
			} 		
		}		

		if possible {
			sumPossibleIds += game.ID 
		}
	}

	return sumPossibleIds
}

func parseGame(line string) Game {
	game := Game{}
	line = strings.Replace(line, "Game", "", 1)
	parts := strings.Split(line, ":")
	gameIdInt, _ := strconv.Atoi(strings.TrimSpace(parts[0])) 
	game.ID = gameIdInt
	setsStrings := strings.Split(strings.TrimSpace(parts[1]), ";")
	for _, setString := range setsStrings {
		setString = strings.TrimSpace(setString)
		game.Sets = append(game.Sets, parseSet(setString))
	}
	return game
}

func parseSet(setString string) ColorSet {
	set := NewColorSet()
	for _, colorString := range strings.Split(setString, ",") {
		pairs := strings.Split(strings.TrimSpace(colorString), " ")
		countString := strings.TrimSpace(pairs[0])
		color := strings.TrimSpace(pairs[1])
		if (color == "red") {
			set.Red, _ = strconv.Atoi(countString)
		}
		if (color == "blue") {
			set.Blue, _ = strconv.Atoi(countString)
		}
		if (color == "green") {
			set.Green, _ = strconv.Atoi(countString)
		}
	}
	return set
}
