package day7

func leftPhase(used []int) []int {
	initial := []int{0, 1, 2, 3, 4}
	var left []int
	for _, value := range initial {
		found := false
		for _, usedKey := range used {
			if value == usedKey {
				found = true
			}
		}
		if !found {
			left = append(left, value)
			found = false
		}
	}
	return left
}

func addOutput(instructions []int, used []int, input int, results *[]int) {
	leftPhases := leftPhase(used)
	for _, phase := range leftPhases {
		output := IntCodeComputer(instructions, phase, input)
		if len(used) == 4 {
			*results = append(*results, output)
		} else {
			updatedUsed := append(used, phase)
			addOutput(instructions, updatedUsed, output, results)
		}
	}
}

func FindLargestOutput1(instructions []int) int {
	var results []int
	addOutput(instructions, []int{}, 0, &results)
	max := 0
	for _, result := range results {
		if result > max {
			max = result
		}
	}
	return max
}
