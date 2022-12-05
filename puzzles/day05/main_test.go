package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		actual := part1("./sample_input.txt")
		expected := "CMZ"
		if actual != expected {
			t.Errorf("Expected %s but actual was %s", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		actual := part2("./sample_input.txt")
		expected := "MCD"
		if actual != expected {
			t.Errorf("Expected %s but actual was %s", expected, actual)
		}
	})

}
