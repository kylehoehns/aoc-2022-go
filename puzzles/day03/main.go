package main

import (
	"fmt"

	"github.com/kylehoehns/aoc-2022-go/utils/files"
)

func main() {
	priorities := createPriorities()
	fmt.Println("Part 1: ", part1("input.txt", priorities))
	fmt.Println("Part 2: ", part2("input.txt", priorities))
}

func part1(name string, priorities map[string]int) int {
	lines := files.ReadLines(name)

	totalPriority := 0
	for _, line := range lines {
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

	out:
		for _, left := range compartment1 {
			for _, right := range compartment2 {
				if left == right {
					totalPriority += priorities[string(left)]
					break out
				}
			}
		}
	}
	return totalPriority
}

func part2(name string, priorities map[string]int) int {
	lines := files.ReadLines(name)

	var groups [][]string
	for i := 0; i < len(lines); i += 3 {
		groups = append(groups, lines[i:i+3])
	}

	totalPriority := 0

	for _, group := range groups {
	out:
		for _, one := range group[0] {
			for _, two := range group[1] {
				if one == two {
					for _, three := range group[2] {
						if two == three {
							totalPriority += priorities[string(one)]
							break out
						}
					}
				}
			}
		}
	}

	return totalPriority
}

func createPriorities() map[string]int {
	var priorities = make(map[string]int)
	priority := 1
	for r := 'a'; r <= 'z'; r++ {
		priorities[string(r)] = priority
		priority++
	}
	for r := 'A'; r <= 'Z'; r++ {
		priorities[string(r)] = priority
		priority++
	}
	return priorities
}
