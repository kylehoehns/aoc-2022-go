package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", calculate("input.txt", 2))
	fmt.Println("Part 2: ", calculate("input.txt", 10))
}

func calculate(name string, ropeSize int) int {
	lines := files.ReadLines(name)
	visited := make(map[Point]bool, 0)
	rope := make([]Point, ropeSize)
	for i := 0; i < ropeSize; i++ {
		rope[i] = NewPoint()
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount := ints.FromString(parts[1])
		for i := 0; i < amount; i++ {
			// move head
			rope[0] = rope[0].move(direction)

			// catch up the rest of the rope pieces
			for j := 1; j < len(rope); j++ {
				rope[j] = rope[j].moveToward(rope[j-1])
			}
			visited[rope[len(rope)-1]] = true
		}
	}
	return len(visited)
}

type Point struct {
	x int
	y int
}

func NewPoint() Point {
	return Point{
		x: 0,
		y: 0,
	}
}

func (p Point) moveToward(other Point) Point {

	changeX := other.x - p.x
	changeY := other.y - p.y

	if ints.Abs(changeX) > 1 || ints.Abs(changeY) > 1 {
		if changeX < 0 {
			p.x--
		} else if changeX > 0 {
			p.x++
		}
		if changeY < 0 {
			p.y--
		} else if changeY > 0 {
			p.y++
		}
	}

	return p
}

func (p Point) move(direction string) Point {
	if direction == "U" {
		p.y++
	} else if direction == "R" {
		p.x++
	} else if direction == "L" {
		p.x--
	} else if direction == "D" {
		p.y--
	} else {
		panic("unknown direction " + direction)
	}
	return p
}
