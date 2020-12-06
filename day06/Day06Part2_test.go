package day05

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
		"",
	}
	result := solvePart2(input)
	if result != 6 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day06_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
