package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) string {
	lines := files.ReadLines(name)

	crateLines, instructionLines := parseCratesAndInstructions(lines)
	numberOfStacks := determineNumberOfStacks(crateLines)
	crates := parseCrates(crateLines, numberOfStacks)

	var move int
	var from int
	var to int
	for _, line := range instructionLines {
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		// pull off the ones we're going to move
		cratesToMove := crates[from-1][:move]
		// change the stack to remove those entries
		crates[from-1] = crates[from-1][move:]
		for _, crate := range cratesToMove {
			crates[to-1] = append([]string{crate}, crates[to-1]...)
		}
	}
	return determineOutput(numberOfStacks, crates)
}

func part2(name string) string {
	lines := files.ReadLines(name)

	crateLines, instructionLines := parseCratesAndInstructions(lines)
	numberOfStacks := determineNumberOfStacks(crateLines)
	crates := parseCrates(crateLines, numberOfStacks)

	var move int
	var from int
	var to int
	for _, line := range instructionLines {
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		cratesToMove := crates[from-1][:move]
		crates[from-1] = crates[from-1][move:]
		newArray := make([]string, len(cratesToMove))
		copy(newArray, cratesToMove)
		crates[to-1] = append(newArray, crates[to-1]...)
	}
	return determineOutput(numberOfStacks, crates)
}

func determineOutput(numberOfStacks int, crates map[int][]string) string {
	output := ""
	for i := 0; i <= numberOfStacks-1; i++ {
		output += crates[i][0]
	}
	return output
}

func determineNumberOfStacks(crateLines []string) int {
	lastCrateLine := crateLines[len(crateLines)-1]
	lastSpace := strings.LastIndex(lastCrateLine, "   ")
	numberOfStacks := ints.FromString(lastCrateLine[lastSpace+3:])
	return numberOfStacks
}

func parseCrates(crateLines []string, numberOfStacks int) map[int][]string {
	crates := make(map[int][]string)
	for _, line := range crateLines {
		if strings.ContainsRune(line, '[') {
			for i := 0; i < numberOfStacks; i++ {
				chunk := 1
				if i > 0 {
					chunk = 1 + (i * 4)
				}
				if len(line) >= chunk {
					letter := string(line[chunk])
					if letter != " " {
						crates[i] = append(crates[i], letter)
					}
				}
			}
		}
	}
	return crates
}

func parseCratesAndInstructions(lines []string) ([]string, []string) {
	startingCrates := make([]string, 0)
	instructions := make([]string, 0)
	readingCrates := true
	for _, line := range lines {
		if line == "" {
			readingCrates = false
			continue
		}
		if readingCrates {
			startingCrates = append(startingCrates, line)
		} else {
			instructions = append(instructions, line)
		}
	}
	return startingCrates, instructions
}
