package day03

import (
	"advent_of_code/utils"
	"strconv"
)

func Solution() string { 
	var input, _ = utils.ReadInput("day03")
	return "A: " + strconv.Itoa(SolutionA(input)) + ", B: " + strconv.Itoa(SolutionB(input))
}
