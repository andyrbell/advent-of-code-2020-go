package day09

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
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
	result := solvePart1(input, 5)
	if result != 127 {
		t.Errorf("result: %v\n", result)
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadInt64s("day09_input.txt")
	result := solvePart1(input, 25)
	fmt.Printf("Part 1: %v\n", result)
}
