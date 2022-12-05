package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		elfFoodItems := convertFileToElfFood("sample_input.txt")
		actual := part1(elfFoodItems)
		expected := 24000
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		elfFoodItems := convertFileToElfFood("sample_input.txt")
		actual := part2(elfFoodItems)
		expected := 45000
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}
