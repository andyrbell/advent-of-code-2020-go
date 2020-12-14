package day08

import (
	"strconv"
	"strings"
)

func solvePart1(input []string) int {
	accumulator := 0
	pointer := 0
	visited := make(map[int]bool)

	for visited[pointer] == false {
		command, argument := parseCommand(input[pointer])
		visited[pointer] = true
		switch command {
		case "acc":
			accumulator = accumulator + argument
			pointer++
		case "nop":
			pointer++
		case "jmp":
			pointer = pointer + argument
		}
	}
	return accumulator
}

func parseCommand(input string) (string, int) {
	parts := strings.Split(input, " ")
	argument, _ := strconv.Atoi(parts[1])
	return parts[0], argument
}
