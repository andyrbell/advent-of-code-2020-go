package day07

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"regexp"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	result := solvePart1(input)
	if result != 4 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadStrings("day07_input.txt")
	result := solvePart1(input)
	fmt.Printf("Part 1: %v\n", result)
}

func TestRegex(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	outerBagRegex := regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain .*`)
	outerBag := outerBagRegex.FindStringSubmatch(input)[1]
	fmt.Printf("outerBag: %v\n", outerBag)
	innerBagRegex := regexp.MustCompile(`(\d+ [a-z]+ [a-z]+) bags?`)
	innerBags := innerBagRegex.FindAllStringSubmatch(input, -1)

	for _, innerBag := range innerBags {
		fmt.Printf("innerBag: %v\n", innerBag[1])
	}
}

func TestOuterBagRegex(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	outerBag := parseOuterBag(input)
	if outerBag != "light red" {
		t.Fail()
	}
}

func TestInnerBagRegex(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	innerBags := parseInnerBags(input)
	expected := make(map[string]int)
	expected["bright white"] = 1
	expected["muted yellow"] = 2
	fmt.Printf("expected: %v\n", expected)
	fmt.Printf("actual: %v\n", innerBags)
}
