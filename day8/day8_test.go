package day8

import (
	"adventofcode"
	"fmt"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	t.Run("test DecodeImage", func(t *testing.T) {
		input := adventofcode.GetInputsNumInOneline("./input.txt", "")
		got := DecodeImage(input)
		fmt.Println(got)
	})
}
