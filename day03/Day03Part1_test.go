package day03

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
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
	result := solvePart1(input)
	if result != 7 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadStrings("day03_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
