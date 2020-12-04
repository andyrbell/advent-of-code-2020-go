package day03

import "advent-of-code-2020-go/utils"

func solvePart2(input []string) int {
	slopes := []utils.Point{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	result := 1
	grid := utils.Grid{Geology: input}
	for _, slope := range slopes {
		result = result * treeCount(grid, slope)
	}
	return result
}

func treeCount(grid utils.Grid, slope utils.Point) int {
	treeCount := 0
	position := utils.Origin
	for grid.Contains(position) {
		if grid.IsTree(position) {
			treeCount++
		}
		position = position.Add(slope)
	}
	return treeCount
}
