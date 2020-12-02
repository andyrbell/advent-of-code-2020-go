package day01

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	result := solvePart2(input)
	if result != 241861950 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadInts("day01_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
