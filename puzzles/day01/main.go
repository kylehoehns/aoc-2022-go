package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/kylehoehns/aoc-2022-go/utils"
)

func main() {
	elfFoodItems := convertFileToElfFood("input.txt")
	fmt.Println("Part 1: ", part1(elfFoodItems))
	fmt.Println("Part 2: ", part2(elfFoodItems))
}

func convertFileToElfFood(name string) []ElfWithFood {
	lines := utils.ReadFile(name)

	var elfFoodItems []ElfWithFood
	elfWithFood := ElfWithFood{}
	for i, line := range lines {
		if calories, err := strconv.Atoi(line); err == nil {
			elfWithFood.items = append(elfWithFood.items, calories)
		} else {
			elfFoodItems = append(elfFoodItems, elfWithFood)
			elfWithFood = ElfWithFood{}
		}

		if i == len(lines)-1 {
			elfFoodItems = append(elfFoodItems, elfWithFood)
		}
	}
	return elfFoodItems
}

type ElfWithFood struct {
	items []int
}

func (e *ElfWithFood) totalCalories() int {
	total := 0
	for _, i := range e.items {
		total += i
	}
	return total
}

func part1(elfWithFoodItems []ElfWithFood) int {
	var mostCalories int = 0
	for _, elfWithFood := range elfWithFoodItems {
		if mostCalories < elfWithFood.totalCalories() {
			mostCalories = elfWithFood.totalCalories()
		}
	}
	return mostCalories
}

func part2(elfWithFoodItems []ElfWithFood) int {
	var allCalories []int
	for _, elf := range elfWithFoodItems {
		allCalories = append(allCalories, elf.totalCalories())
	}
	// reverse sort and grab the top 3
	sort.Sort(sort.Reverse(sort.IntSlice(allCalories)))
	return allCalories[0] + allCalories[1] + allCalories[2]
}
