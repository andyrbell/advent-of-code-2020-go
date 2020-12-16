package day09

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	result := solvePart2(input, 127)
	if result != 62 {
		t.Errorf("result: %v\n", result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadInt64s("day09_input.txt")
	result := solvePart2(input, 1309761972)
	fmt.Printf("Part 2: %v\n", result)
}
