package day01

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	result := solvePart1(input)
	if result != 514579 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadInts("day01_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
