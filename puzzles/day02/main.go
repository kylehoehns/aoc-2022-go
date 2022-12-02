package main

import (
	"fmt"
	"strings"

	"github.com/kylehoehns/aoc-2022-go/utils/files"
)

func main() {
	rounds := parseRounds("input.txt")
	fmt.Println("Part 1: ", part1(rounds))
	fmt.Println("Part 2: ", part2(rounds))
}

func parseRounds(name string) []round {
	rows := files.ReadLines(name)
	var rounds []round
	for _, row := range rows {
		values := strings.Split(row, " ")
		rounds = append(rounds, round{opponentPlay: values[0], ourPlay: values[1]})
	}
	return rounds
}

type round struct {
	opponentPlay string
	ourPlay      string
}

func (r round) roundScorePart1() int {
	return r.resultScore() + shapeScore(r.ourPlay)
}

func (r round) resultScore() int {
	if r.isDraw() {
		return 3
	} else if r.opponentWins() {
		return 0
	} else {
		// we win
		return 6
	}
}

func (r round) isDraw() bool {
	return (r.opponentPlay == "A" && r.ourPlay == "X") ||
		(r.opponentPlay == "B" && r.ourPlay == "Y") ||
		(r.opponentPlay == "C" && r.ourPlay == "Z")
}

func (r round) opponentWins() bool {
	// Opponent A -> Rock, B -> Paper, C -> Scissors
	// Us X --> Rock, Y --> Paper, Z --> Scissors
	return (r.opponentPlay == "A" && r.ourPlay == "Z") ||
		(r.opponentPlay == "B" && r.ourPlay == "X") ||
		(r.opponentPlay == "C" && r.ourPlay == "Y")
}

func shapeScore(shape string) int {
	// Rock -> X -> 1, Paper -> Y -> 2, Scissors -> Z -> 3
	if shape == "X" {
		return 1
	} else if shape == "Y" {
		return 2
	} else {
		return 3
	}
}

func (r round) predictedScore() int {
	if r.ourPlay == "X" {
		return 0
	} else if r.ourPlay == "Y" {
		return 3
	} else {
		return 6
	}
}

func (r round) drawShape() string {
	if r.opponentPlay == "A" {
		return "X"
	} else if r.opponentPlay == "B" {
		return "Y"
	} else {
		return "Z"
	}
}

func (r round) loseShape() string {
	if r.opponentPlay == "A" {
		return "Z"
	} else if r.opponentPlay == "B" {
		return "X"
	} else {
		return "Y"
	}
}

func (r round) winShape() string {
	if r.opponentPlay == "A" {
		return "Y"
	} else if r.opponentPlay == "B" {
		return "Z"
	} else {
		return "X"
	}
}

func (r round) predictedScoreAndShape() (int, string) {
	predictedScore := r.predictedScore()
	var shape string
	if predictedScore == 0 {
		shape = r.loseShape()
	} else if predictedScore == 3 {
		shape = r.drawShape()
	} else {
		shape = r.winShape()
	}
	return predictedScore, shape
}

func (r round) roundScorePart2() int {
	resultScore, resultShape := r.predictedScoreAndShape()
	return resultScore + shapeScore(resultShape)
}

func part1(rounds []round) int {
	total := 0
	for _, round := range rounds {
		total += round.roundScorePart1()
	}
	return total
}

func part2(rounds []round) int {
	total := 0
	for _, round := range rounds {
		total += round.roundScorePart2()
	}
	return total
}
