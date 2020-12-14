package day08

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	result := solvePart2(input)
	if result != 8 {
		t.Errorf("result: %v\n", result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day08_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}
