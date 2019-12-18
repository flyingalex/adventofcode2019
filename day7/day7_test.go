package day7

import (
	"adventofcode"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestFindLargestOutput(t *testing.T) {
	t.Run("test leftPhase", func(t *testing.T) {
		got := leftPhase([]int{0, 1}, []int{0, 1, 2, 3, 4})
		want := []int{2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("want: %#v, got: %#v", want, got)
		}

		got = leftPhase([]int{0, 1, 3}, []int{0, 1, 2, 3, 4})
		want = []int{2, 4}
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("want: %#v, got: %#v", want, got)
		}

		got = leftPhase([]int{0, 3}, []int{0, 1, 2, 3, 4})
		want = []int{1, 2, 4}
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("want: %#v, got: %#v", want, got)
		}
	})
	t.Run("test FindLargestOutput1", func(t *testing.T) {
		instructions := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
		got := FindLargestOutput1(instructions)
		want := 43210
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})
	t.Run("test FindLargestOutput1", func(t *testing.T) {
		instructions := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
		got := FindLargestOutput1(instructions)
		want := 54321
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})
	t.Run("test FindLargestOutput1", func(t *testing.T) {
		instructions := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
		got := FindLargestOutput1(instructions)
		want := 65210
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})

	t.Run("check FindLargestOutput1 result", func(t *testing.T) {
		instructions := adventofcode.GetInputsNumInOneline("./input.txt", ",")
		got := FindLargestOutput1(instructions)
		want := 92663 // result for refactor
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})

	t.Run("test allPhaseSettingSequence", func(t *testing.T) {
		var settings [][]int
		allPhaseSettingSequence([]int{}, &settings)
		fmt.Printf("%#v", settings)
	})

	t.Run("test FindLargestOutput2", func(t *testing.T) {
		instructions := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
		got := FindLargestOutput2(instructions)
		want := 139629729
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})

	t.Run("test FindLargestOutput2", func(t *testing.T) {
		instructions := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
		got := FindLargestOutput2(instructions)
		want := 139629729
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})
}
