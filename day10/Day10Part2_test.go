package day10

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	result := solvePart2(input)
	if result != 8 {
		t.Errorf("result: %v\n", result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadInts("day10_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
