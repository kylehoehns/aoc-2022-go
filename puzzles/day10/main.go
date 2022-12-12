package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
	"github.com/kylehoehns/aoc-2022-go/utils/utils"
	"math"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	part2("input.txt")
}

func part1(name string) int {
	lines := files.ReadLines(name)
	cycle := 1
	cycleValues := make(map[int]int, 0)
	x := 1
	for _, line := range lines {
		cycleValues[cycle] = x
		parts := strings.Split(line, " ")
		command := parts[0]
		cycle++
		cycleValues[cycle] = x
		if command == "addx" {
			cycle++
			cycleValues[cycle] = x
			amount := ints.FromString(parts[1])
			x += amount
		}
	}
	values := []int{20, 60, 100, 140, 180, 220}
	signalStrength := 0
	for _, v := range values {
		signalStrength += cycleValues[v] * v
	}
	return signalStrength
}

func determineDrawValue(spriteLocation int, cycle int, row *int) string {
	current := (cycle % 40) - 2
	row = utils.Ptr(int(math.Floor(float64(cycle) / 40)))
	if current-1 == spriteLocation || current == spriteLocation || current+1 == spriteLocation {
		return "#"
	}
	return " "
}

func part2(name string) {
	lines := files.ReadLines(name)
	cycle := 1
	cycleValues := make(map[int]int, 0)
	drawValues := make(map[int]string, 0)
	spriteLocation := 0
	row := 0
	for _, line := range lines {
		cycleValues[cycle] = spriteLocation
		drawValues[cycle] = determineDrawValue(spriteLocation, cycle, &row)
		cycle++
		cycleValues[cycle] = spriteLocation
		parts := strings.Split(line, " ")
		command := parts[0]
		drawValues[cycle] = determineDrawValue(spriteLocation, cycle, &row)
		if command == "addx" {
			cycle++
			cycleValues[cycle] = spriteLocation
			drawValues[cycle] = determineDrawValue(spriteLocation, cycle, &row)
			amount := ints.FromString(parts[1])
			spriteLocation += amount
		}
	}

	for i := 1; i <= len(drawValues); i++ {
		fmt.Print(drawValues[i])
		if i%40 == 0 {
			fmt.Println("")
		}
	}
}
