package day09

import (
	"advent-of-code-2020-go/utils"
)

func solvePart1(input []int64, preambleSize int) int64 {

	for i := 0; i+preambleSize < len(input); i++ {
		numberToValidate := input[i+preambleSize]
		preamble := input[i : i+preambleSize]
		combinations := utils.Combinations(preamble)
		sumOfPairs := make(map[int64]bool)
		for _, pair := range combinations {
			sumOfPairs[pair.Sum()] = true
		}
		if sumOfPairs[numberToValidate] == false {
			return numberToValidate
		}
	}

	return -1
}
