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

func FindLargestOutput(instructions []int) int {
	phaseSetting := []int{0, 1, 2, 3, 4}
	var results []int
	for _, phase1 := range phaseSetting {
		output1 := IntCodeComputer(instructions, phase1, 0)
		leftPhases := leftPhase([]int{phase1})
		for _, phase2 := range leftPhases {
			output2 := IntCodeComputer(instructions, phase2, output1)
			leftPhases := leftPhase([]int{phase1, phase2})
			for _, phase3 := range leftPhases {
				output3 := IntCodeComputer(instructions, phase3, output2)
				leftPhases := leftPhase([]int{phase1, phase2, phase3})
				for _, phase4 := range leftPhases {
					output4 := IntCodeComputer(instructions, phase4, output3)
					leftPhases := leftPhase([]int{phase1, phase2, phase3, phase4})
					for _, phase5 := range leftPhases {
						output := IntCodeComputer(instructions, phase5, output4)
						results = append(results, output)
					}
				}
			}
		}
	}
	max := 0
	for _, result := range results {
		if result > max {
			max = result
		}
	}
	return max
}
