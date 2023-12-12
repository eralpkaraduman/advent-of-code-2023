package day01

import (
	"advent_of_code/utils"
	"strconv"
)

func Solution() string { 
	var input, _ = utils.ReadInput("day01")
	return "A: " + strconv.Itoa(SolutionA(input)) + ", B: " + strconv.Itoa(SolutionB(input))
}
