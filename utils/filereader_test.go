package utils

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {

	t.Run("Take a file name relative to the current file and return a string array of the lines", func(t *testing.T) {

		expected := []string{"this", "is", "a", "test", "file"}
		actual := ReadFile("./input.txt")

		if !reflect.DeepEqual(expected, actual) {
			t.Fail()
			t.Logf("Expected %s but actual was %s", expected, actual)
		}

	})

}
