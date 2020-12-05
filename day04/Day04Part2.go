package day04

import (
	"advent-of-code-2020-go/utils"
	"strings"
)

func solvePart2(input []string) int {
	passportData := make(map[string]string)
	var validPassports []utils.Passport

	for _, line := range input {
		if line == "" {
			passport := utils.Passport{PassportData: passportData}
			if passport.HasValidData() {
				validPassports = append(validPassports, passport)
			}
			passportData = make(map[string]string)
		} else {
			fields := strings.Fields(line)
			for _, field := range fields {
				tokens := strings.Split(field, ":")
				name := tokens[0]
				value := tokens[1]
				passportData[name] = value
			}
		}
	}
	return len(validPassports)
}
