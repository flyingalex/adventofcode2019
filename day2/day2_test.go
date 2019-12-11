package day2

import (
	"log"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestIntcodeProgram(t *testing.T) {
	t.Run("check IntcodeProgram", func(t *testing.T) {
		arr1 := []int{1, 0, 0, 0, 99}
		got := IntcodeProgram(arr1)
		expected := []int{2, 0, 0, 0, 99}
		if !Equal(got, expected) {
			log.Fatalf("want: %#v, got: %#v", expected, got)
		}

		arr1 = []int{2, 3, 0, 3, 99}
		got = IntcodeProgram(arr1)
		expected = []int{2, 3, 0, 6, 99}
		if !Equal(got, expected) {
			log.Fatalf("want: %#v, got: %#v", expected, got)
		}

		arr1 = []int{2, 4, 4, 5, 99, 0}
		got = IntcodeProgram(arr1)
		expected = []int{2, 4, 4, 5, 99, 9801}
		if !Equal(got, expected) {
			log.Fatalf("want: %#v, got: %#v", expected, got)
		}

		arr1 = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		got = IntcodeProgram(arr1)
		expected = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
		if !Equal(got, expected) {
			log.Fatalf("want: %#v, got: %#v", expected, got)
		}

	})
}

func TestIntcodeProgram2(t *testing.T) {
	t.Run("check IntcodeProgram2", func(t *testing.T) {
		IntcodeProgram2()
	})
}
