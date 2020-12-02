package day01

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"sort"
)

func solvePart2(input []int) int {
	sort.Ints(input)
	for _, a := range input {
		target := 2020 - a

		b, c := solve(input, target)

		if b != -1 && c != -1 {
			result := a * b * c
			fmt.Printf("%v * %v * %v = %v\n", a, b, c, result)
			return result
		}
	}
	return -1
}

func solve(input []int, targetSum int) (int, int) {
	for _, value := range input {
		target := targetSum - value
		if utils.Contains(input, target) {
			return target, value
		}
	}
	return -1, -1
}
