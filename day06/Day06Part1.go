package day06

func solvePart1(input []string) int {
	answers := make(map[string]bool)
	sumOfCounts := 0

	for _, line := range input {
		if line == "" {
			sumOfCounts = sumOfCounts + len(answers)
			answers = make(map[string]bool)
		} else {
			for _, c := range line {
				answers[string(c)] = true
			}
		}
	}
	return sumOfCounts
}
