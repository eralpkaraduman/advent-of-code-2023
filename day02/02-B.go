package day02

func SolutionB(input []string) int {
	// input = test

	sumPowers := 0

	for _, line := range input {
		game := parseGame(line)

		minColorSet := NewColorSet()

		for _, set := range game.Sets {
			if set.Red > minColorSet.Red {
				minColorSet.Red = set.Red
			}
			if set.Green > minColorSet.Green {
				minColorSet.Green = set.Green
			}
			if set.Blue > minColorSet.Blue {
				minColorSet.Blue = set.Blue
			}
		}

		sumPowers += minColorSet.Red * minColorSet.Green * minColorSet.Blue
	}

	return sumPowers
}