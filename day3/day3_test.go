package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFindIntersectionPoint(t *testing.T) {
	t.Run("check FindIntersectionPoint", func(t *testing.T) {
		wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
		wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
		got := FindIntersectionPoint(wire1, wire2)
		expected := 159
		if got != expected {
			log.Fatalf("want: %d, got: %d", expected, got)
		}

		wire1 = []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
		wire2 = []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
		got = FindIntersectionPoint(wire1, wire2)
		expected = 135
		if got != expected {
			log.Fatalf("want: %d, got: %d", expected, got)
		}
	})

	t.Run("get part 1 puzzle result", func(t *testing.T) {
		file, err := os.Open("./day3_1.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var wires [][]string
		for scanner.Scan() {
			text := scanner.Text()
			wires = append(wires, strings.Split(text, ","))
		}
		got := FindIntersectionPoint(wires[0], wires[1])
		fmt.Println(got)
	})

}

func TestFindMinimumSteps(t *testing.T) {
	t.Run("check GetMinimumSteps", func(t *testing.T) {
		wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
		wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
		got := FindMinimumSteps(wire1, wire2)
		expected := 610
		if got != expected {
			log.Fatalf("want: %d, got: %d", expected, got)
		}

		wire1 = []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
		wire2 = []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
		got = FindMinimumSteps(wire1, wire2)
		expected = 410
		if got != expected {
			log.Fatalf("want: %d, got: %d", expected, got)
		}

	})

	t.Run("get part 2 puzzle result", func(t *testing.T) {
		file, err := os.Open("./day3_2.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var wires [][]string
		for scanner.Scan() {
			text := scanner.Text()
			wires = append(wires, strings.Split(text, ","))
		}
		got := FindMinimumSteps(wires[0], wires[1])
		fmt.Println(got)
	})
}
