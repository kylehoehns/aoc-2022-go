package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		rounds := parseRounds("./sample_input.txt")
		actual := part1(rounds)
		expected := 15
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		rounds := parseRounds("./sample_input.txt")
		actual := part2(rounds)
		expected := 12
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}
