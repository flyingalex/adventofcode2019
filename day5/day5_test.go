package main

import (
	"adventofcode"
	"log"
	"reflect"
	"testing"
)

func TestDiagnostTool(t *testing.T) {
	t.Run("check splitInstruction", func(t *testing.T) {
		got := splitInstruction(76)
		want := []int{0, 0, 7, 6}
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("want: %#v, got: %#v", want, got)
		}
	})
	t.Run("check IntCodeComputer part1", func(t *testing.T) {
		instructionsInt := adventofcode.GetInputsNumInOneline("./input.txt", ",")
		got := IntCodeComputer(instructionsInt, 1)
		want := 6745903
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})

	t.Run("check IntCodeComputer part2", func(t *testing.T) {
		instructionsInt := adventofcode.GetInputsNumInOneline("./input.txt", ",")
		got := IntCodeComputer(instructionsInt, 5)
		want := 9168267
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})
}
