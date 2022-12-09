package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		actual := calculate("./sample_input.txt", 2)
		expected := 13
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		actual := calculate("./sample_input_2.txt", 10)
		expected := 36
		if actual != expected {
			t.Errorf("Expected %d but actual was %d", expected, actual)
		}
	})

}
