package day02

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	result := solvePart2(input)
	if result != 1 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day02_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
