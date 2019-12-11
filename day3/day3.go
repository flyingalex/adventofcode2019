package day3

import (
	"log"
	"strconv"
)

func absInt(number int) int {
	if number > 0 {
		return number
	}
	return -number
}

func minInt(result, newNumber int) int {
	if result == 0 {
		return newNumber
	}
	if result < newNumber {
		return result
	}
	return newNumber
}

func FindIntersectionPoint(wire1 []string, wire2 []string) int {
	wire1Points := make(map[string]bool)
	current := [2]int{0, 0}
	for _, point := range wire1 {
		// collect wire1's position
		position := string([]rune(point)[0])
		steps, err := strconv.Atoi(point[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch position {
		case "R":
			for i := 0; i < steps; i++ {
				current[0]++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = true
			}
		case "L":
			for i := 0; i < steps; i++ {
				current[0]--
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = true
			}
		case "U":
			for i := 0; i < steps; i++ {
				current[1]++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = true
			}
		case "D":
			for i := 0; i < steps; i++ {
				current[1]--
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = true
			}
		}
	}
	current = [2]int{0, 0}
	result := 0
	for _, point := range wire2 {
		// collect wire1's position
		position := string([]rune(point)[0])
		steps, err := strconv.Atoi(point[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch position {
		case "R":
			for i := 0; i < steps; i++ {
				current[0]++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				currentDistance := absInt(current[0]) + absInt(current[1])
				if wire1Points[key] {
					result = minInt(result, currentDistance)
				}
			}
		case "L":
			for i := 0; i < steps; i++ {
				current[0]--
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				currentDistance := absInt(current[0]) + absInt(current[1])
				if wire1Points[key] {
					result = minInt(result, currentDistance)
				}
			}
		case "U":
			for i := 0; i < steps; i++ {
				current[1]++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				currentDistance := absInt(current[0]) + absInt(current[1])
				if wire1Points[key] {
					result = minInt(result, currentDistance)
				}
			}
		case "D":
			for i := 0; i < steps; i++ {
				current[1]--
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				currentDistance := absInt(current[0]) + absInt(current[1])
				if wire1Points[key] {
					result = minInt(result, currentDistance)
				}
			}
		}
	}
	return result
}

func FindMinimumSteps(wire1 []string, wire2 []string) int {
	wire1Points := make(map[string]int)
	current := [2]int{0, 0}
	step := 0
	for _, point := range wire1 {
		// collect wire1's position
		position := string([]rune(point)[0])
		steps, err := strconv.Atoi(point[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch position {
		case "R":
			for i := 0; i < steps; i++ {
				current[0]++
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = step
			}
		case "L":
			for i := 0; i < steps; i++ {
				current[0]--
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = step
			}
		case "U":
			for i := 0; i < steps; i++ {
				current[1]++
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = step
			}
		case "D":
			for i := 0; i < steps; i++ {
				current[1]--
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				wire1Points[key] = step
			}
		}
	}

	current = [2]int{0, 0}
	result := 0
	step = 0
	for _, point := range wire2 {
		// collect wire1's position
		position := string([]rune(point)[0])
		steps, err := strconv.Atoi(point[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch position {
		case "R":
			for i := 0; i < steps; i++ {
				current[0]++
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				if wire1Points[key] > 0 {
					result = minInt(result, step+wire1Points[key])
				}
			}
		case "L":
			for i := 0; i < steps; i++ {
				current[0]--
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				if wire1Points[key] > 0 {
					result = minInt(result, step+wire1Points[key])
				}
			}
		case "U":
			for i := 0; i < steps; i++ {
				current[1]++
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				if wire1Points[key] > 0 {
					result = minInt(result, step+wire1Points[key])
				}
			}
		case "D":
			for i := 0; i < steps; i++ {
				current[1]--
				step++
				key := strconv.Itoa(current[0]) + "|" + strconv.Itoa(current[1])
				if wire1Points[key] > 0 {
					result = minInt(result, step+wire1Points[key])
				}
			}
		}
	}
	return result
}
