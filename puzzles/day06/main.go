package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/strings"
)

func main() {
	input := files.Read("input.txt")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(name string) int {
	i, done := findMarkerIndex(name, 4)
	if done {
		return i
	}

	panic("Input did not have any valid marker")
}

func part2(name string) int {
	i, done := findMarkerIndex(name, 14)
	if done {
		return i
	}

	panic("Input did not have any valid marker")
}

func findMarkerIndex(name string, digits int) (int, bool) {
	for i := digits; i < len(name); i++ {
		marker := name[i-digits : i]
		if strings.HasAllUniqueRunes(marker) {
			return i, true
		}
	}
	return 0, false
}
