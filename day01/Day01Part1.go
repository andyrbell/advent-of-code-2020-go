package day01

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"sort"
)

func solvePart1(input []int) int {
	sort.Ints(input)
	for _, value := range input {
		target := 2020 - value
		if utils.Contains(input, target) {
			result := target * value
			fmt.Printf("%v * %v = %v\n", target, value, result)
			return result
		}
	}
	return -1
}
