package day03

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	result := solvePart2(input)
	if result != 336 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day03_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
