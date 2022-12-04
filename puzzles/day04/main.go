package main

import (
	"fmt"
	"strings"

	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := files.ReadLines(name)

	pairsFullyContained := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		firstParts := strings.Split(parts[0], "-")
		firstMin := ints.FromString(firstParts[0])
		firstMax := ints.FromString(firstParts[1])
		secondParts := strings.Split(parts[1], "-")
		secondMin := ints.FromString(secondParts[0])
		secondMax := ints.FromString(secondParts[1])

		if (firstMin <= secondMin && firstMax >= secondMax) || (secondMin <= firstMin && secondMax >= firstMax) {
			pairsFullyContained++
		}
	}
	return pairsFullyContained
}

func part2(name string) int {
	lines := files.ReadLines(name)

	pairsThatOverlap := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		firstParts := strings.Split(parts[0], "-")
		firstMin := ints.FromString(firstParts[0])
		firstMax := ints.FromString(firstParts[1])
		secondParts := strings.Split(parts[1], "-")
		secondMin := ints.FromString(secondParts[0])
		secondMax := ints.FromString(secondParts[1])

		if (secondMin >= firstMin && secondMin <= firstMax) || (secondMax >= firstMin && secondMax <= firstMax) ||
			(firstMin >= secondMin && firstMin <= secondMax) || (firstMax >= secondMin && firstMax <= secondMax) {
			pairsThatOverlap++
		}
	}
	return pairsThatOverlap
}
