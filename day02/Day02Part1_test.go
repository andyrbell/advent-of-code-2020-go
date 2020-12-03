package day02

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	result := solvePart1(input)
	if result != 2 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadStrings("day02_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
