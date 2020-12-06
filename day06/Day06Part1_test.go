package day05

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
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
	result := solvePart1(input)
	if result != 11 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadStrings("day06_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
