package day05

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	result := solvePart1(input)
	if result != 820 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadStrings("day05_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}
func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day05_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}

func TestRow(t *testing.T) {
	input := "FBFBBFFRLR"
	row := calculateRow(input)
	if row != 44 {
		fmt.Printf("It no worky %v\n", row)
		t.Fail()
	}
}

func TestColumn(t *testing.T) {
	input := "FBFBBFFRLR"
	column := calculateColumn(input)
	if column != 5 {
		fmt.Printf("It no worky %v\n", column)
		t.Fail()
	}
}
