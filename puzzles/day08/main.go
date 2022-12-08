package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := files.ReadLines(name)
	grid := buildGrid(lines)
	numberOfVisibleTrees := 0
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	for rowIndex, row := range grid {
		for colIndex, treeHeight := range row {
			// if on edge of the grid, it is visible
			if rowIndex == 0 || colIndex == 0 || rowIndex == gridHeight-1 || colIndex == gridWidth-1 {
				numberOfVisibleTrees++
			} else {

				// a tree is visible if all the other trees between it and one edge are shorter

				visibleUp := true
				for i := rowIndex - 1; i >= 0; i-- {
					if grid[i][colIndex] >= treeHeight {
						visibleUp = false
						break
					}
				}

				visibleDown := true
				for i := rowIndex + 1; i <= gridHeight-1; i++ {
					if grid[i][colIndex] >= treeHeight {
						visibleDown = false
						break
					}
				}

				visibleLeft := true
				for i := colIndex - 1; i >= 0; i-- {
					if grid[rowIndex][i] >= treeHeight {
						visibleLeft = false
						break
					}
				}

				visibleRight := true
				for i := colIndex + 1; i <= gridWidth-1; i++ {
					if grid[rowIndex][i] >= treeHeight {
						visibleRight = false
						break
					}
				}

				if visibleUp || visibleDown || visibleLeft || visibleRight {
					numberOfVisibleTrees++
				}

			}

		}
	}
	return numberOfVisibleTrees
}

func part2(name string) int {
	lines := files.ReadLines(name)
	grid := buildGrid(lines)

	highestScenicScore := 0
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	for rowIndex, row := range grid {
		for colIndex, treeHeight := range row {
			scoreUp, scoreDown, scoreLeft, scoreRight := 0, 0, 0, 0

			// if we're on an edge, score is always zero
			if rowIndex == 0 || colIndex == 0 || rowIndex == gridHeight-1 || colIndex == gridWidth-1 {
				continue
			}

			// get a scenic score by seeing how many trees are visible from the current tree in each direction
			for i := rowIndex - 1; i >= 0; i-- {
				compareTreeHeight := grid[i][colIndex]
				if compareTreeHeight < treeHeight {
					scoreUp++
				} else if compareTreeHeight >= treeHeight {
					scoreUp++
					break
				} else {
					break
				}
			}

			for i := rowIndex + 1; i <= gridHeight-1; i++ {
				compareTreeHeight := grid[i][colIndex]
				if compareTreeHeight < treeHeight {
					scoreDown++
				} else if compareTreeHeight >= treeHeight {
					scoreDown++
					break
				} else {
					break
				}
			}

			for i := colIndex - 1; i >= 0; i-- {
				compareTreeHeight := grid[rowIndex][i]
				if compareTreeHeight < treeHeight {
					scoreLeft++
				} else if compareTreeHeight >= treeHeight {
					scoreLeft++
					break
				} else {
					break
				}
			}

			for i := colIndex + 1; i <= gridWidth-1; i++ {
				compareTreeHeight := grid[rowIndex][i]
				if compareTreeHeight < treeHeight {
					scoreRight++
				} else if compareTreeHeight >= treeHeight {
					scoreRight++
					break
				} else {
					break
				}
			}
			currentScore := scoreUp * scoreDown * scoreLeft * scoreRight
			if currentScore > highestScenicScore {
				highestScenicScore = currentScore
			}

		}
	}
	return highestScenicScore
}

func buildGrid(lines []string) [][]int {
	columns := len(lines[0])
	rows := len(lines)
	grid := make([][]int, rows)
	for i := range grid {
		row := make([]int, columns)
		for col, val := range lines[i] {
			row[col] = ints.FromString(string(val))
		}
		grid[i] = row
	}
	return grid
}
