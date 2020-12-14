package day08

import "fmt"

func solvePart2(input []string) int {
	c := make(chan Result)

	for i, instruction := range input {
		command, argument := parseCommand(instruction)
		switch command {
		case "nop":
			tweakedInstructions := tweak(i, input, fmt.Sprintf("jmp %v\n", argument))
			go runBootCode(tweakedInstructions, c)
		case "jmp":
			tweakedInstructions := tweak(i, input, fmt.Sprintf("nop %v\n", argument))
			go runBootCode(tweakedInstructions, c)
		}
	}
	success := <-c
	return success.Accumulator
}

func tweak(index int, instructions []string, newInstruction string) []string {
	var tweakedInstructions = make([]string, len(instructions))
	copy(tweakedInstructions, instructions)
	tweakedInstructions[index] = newInstruction
	return tweakedInstructions
}

type Result struct {
	Accumulator int
	Success     bool
}

func runBootCode(instructions []string, c chan Result) {
	accumulator := 0
	pointer := 0
	visited := make(map[int]bool)

	for visited[pointer] == false && pointer < len(instructions) {
		command, argument := parseCommand(instructions[pointer])
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
	if pointer >= len(instructions) {
		c <- Result{Accumulator: accumulator, Success: true}
	}
}
