package day1

import (
	"fmt"
	"testing"
)

func TestSumFuelRequirements(t *testing.T) {
	t.Run("check func", func(t *testing.T) {
		total := SumFuelRequirements()
		fmt.Print(total)
	})
}
