package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func solvePart1(input []string) int {
	rules := make(map[string]map[string]int)
	bagCount := 0

	for _, line := range input {
		outerBag := parseOuterBag(line)
		innerBags := parseInnerBags(line)
		rules[outerBag] = innerBags
	}

	for k := range rules {
		if containsGoldBag(k, rules) {
			bagCount++
		}
	}
	return bagCount
}

func solvePart2(input []string) int {
	rules := make(map[string]map[string]int)

	for _, line := range input {
		outerBag := parseOuterBag(line)
		innerBags := parseInnerBags(line)
		rules[outerBag] = innerBags
	}

	return bagCount("shiny gold", rules) - 1
}

func bagCount(bag string, rules map[string]map[string]int) int {
	result := 1
	for innerBag, count := range rules[bag] {
		fmt.Printf("innerBag: %v; count: %v\n", innerBag, count)
		result = result + count*bagCount(innerBag, rules)
	}
	return result
}

func containsGoldBag(bag string, rules map[string]map[string]int) bool {
	innerBags := rules[bag]
	if count, ok := innerBags["shiny gold"]; ok {
		fmt.Printf("Found %v shiny gold bags in %v\n", count, bag)
		return true
	}
	for innerBag := range innerBags {
		if containsGoldBag(innerBag, rules) {
			return true
		}
	}
	return false
}

var outerBagRegex = regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain .*`)
var innerBagRegex = regexp.MustCompile(`(\d+ [a-z]+ [a-z]+) bags?`)

func parseOuterBag(rule string) string {
	return outerBagRegex.FindStringSubmatch(rule)[1]
}

func parseInnerBags(rule string) map[string]int {
	results := make(map[string]int)

	innerBags := innerBagRegex.FindAllStringSubmatch(rule, -1)

	for _, innerBag := range innerBags {
		countNamePair := strings.SplitN(innerBag[1], " ", 2)
		count := countNamePair[0]
		name := countNamePair[1]
		fmt.Printf("innerBag: %v %v\n", count, name)
		results[name], _ = strconv.Atoi(count)
	}
	return results
}
