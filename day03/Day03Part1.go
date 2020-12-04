package day03

import "advent-of-code-2020-go/utils"

func solvePart1(input []string) int {
	slope := utils.Point{X: 3, Y: 1}
	treeCount := 0
	position := utils.Origin
	grid := utils.Grid{Geology: input}
	for grid.Contains(position) {
		if grid.IsTree(position) {
			treeCount++
		}
		position = position.Add(slope)
	}
	return treeCount
}
