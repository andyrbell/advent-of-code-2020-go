package day10

import "sort"

func solvePart2(input []int) int {

	outputJoltage := 0
	input = append(input, outputJoltage)
	sort.Ints(input)
	deviceJoltage := input[len(input)-1] + 3
	input = append(input, deviceJoltage)
	cache := make(map[int]int)

	return countPaths(input, 0, cache)
}

func countPaths(adapters []int, i int, cache map[int]int) int {
	if i == len(adapters)-1 {
		return 1
	}
	if v, ok := cache[i]; ok {
		return v
	}
	result := 0
	for j := i + 1; j <= i+3 && j < len(adapters); j++ {
		if adapters[j]-adapters[i] <= 3 {
			result += countPaths(adapters, j, cache)
		}
	}
	cache[i] = result
	return result
}
