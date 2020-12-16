package day09

import "advent-of-code-2020-go/utils"

func solvePart2(input []int64, invalidNumber int64) int64 {

	for sliceLength := 2; sliceLength < len(input); sliceLength++ {
		for i := 0; i+sliceLength < len(input); i++ {
			slice := input[i : i+sliceLength]
			sumOfSlice := utils.Sum(slice)
			if sumOfSlice == invalidNumber {
				return utils.MinAndMax(slice).Sum()
			}
		}
	}

	return -1
}
