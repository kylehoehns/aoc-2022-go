package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		actual := 2
		expected := 2
		if actual != expected {
			t.Fail()
			t.Logf("Expected %d but actual was %d", expected, actual)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		actual := 4
		expected := 4
		if actual != expected {
			t.Fail()
			t.Logf("Expected %d but actual was %d", expected, actual)
		}
	})

}
