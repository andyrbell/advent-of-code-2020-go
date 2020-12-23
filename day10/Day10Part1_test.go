package day10

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
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
	result := solvePart1(input)
	if result != 35 {
		t.Errorf("result: %v\n", result)
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadInts("day10_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
