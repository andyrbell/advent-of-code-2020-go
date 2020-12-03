package day02

import (
	"advent-of-code-2020-go/utils"
)

func solvePart1(input []string) int {
	var validPasswords []utils.Password
	for _, line := range input {
		password := utils.ParsePassword(line)
		if password.IsValid() {
			validPasswords = append(validPasswords, password)
		}
	}

	return len(validPasswords)
}
