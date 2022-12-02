package ints

import "testing"

func TestSumList(t *testing.T) {

	t.Run("Should total all values in a list of ints", func(t *testing.T) {
		expected := 6
		actual := Sum([]int{1, 2, 3})
		if expected != actual {
			t.Fail()
			t.Logf("Expected %d but actual was %d", expected, actual)
		}
	})

}
