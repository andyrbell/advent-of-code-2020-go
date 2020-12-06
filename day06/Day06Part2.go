package day05

import "fmt"

func solvePart2(input []string) int {
	answers := make(map[string]int)
	sumOfCounts := 0
	numberInGroup := 0

	for _, line := range input {
		if line == "" {
			fmt.Println(numberInGroup)
			fmt.Println(answers)
			for _, v := range answers {
				if numberInGroup == v {
					sumOfCounts = sumOfCounts + 1
				}
			}
			answers = make(map[string]int)
			numberInGroup = 0
		} else {
			numberInGroup++
			for _, c := range line {
				question := string(c)
				answers[question] = answers[question] + 1
			}
		}
	}
	return sumOfCounts
}
