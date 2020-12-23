package day10

import "sort"

func solvePart1(input []int) int {

	outputJoltage := 0
	input = append(input, outputJoltage)
	sort.Ints(input)
	deviceJoltage := input[len(input)-1] + 3
	input = append(input, deviceJoltage)
	count1 := 0
	count3 := 0

	for i := 0; i < len(input)-1; i++ {
		difference := input[i+1] - input[i]
		switch difference {
		case 1:
			count1++
		case 3:
			count3++
		}
	}

	return count3 * count1
}
