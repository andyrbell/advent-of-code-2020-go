package day02

import (
	"advent-of-code-2020-go/utils"
)

func solvePart2(input []string) int {
	var validPasswords []utils.Password
	for _, line := range input {
		password := utils.ParsePassword(line)
		if password.IsValidPosition() {
			validPasswords = append(validPasswords, password)
		}
	}

	return len(validPasswords)
}
