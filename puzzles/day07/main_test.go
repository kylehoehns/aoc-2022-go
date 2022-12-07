package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		actual := part1("./sample_input.txt")
		expected := 95437
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		actual := part2("./sample_input.txt")
		expected := 24933642
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}
