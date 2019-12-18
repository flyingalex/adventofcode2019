package day7

func leftPhase(initial, used []int) []int {
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
	initial := []int{0, 1, 2, 3, 4}
	leftPhases := leftPhase(initial, used)
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

func allPhaseSettingSequence(used []int, settings *[][]int) {
	initial := []int{5, 6, 7, 8, 9}
	leftPhases := leftPhase(initial, used)
	for _, phase := range leftPhases {
		if len(used) == 4 {
			phaseSetting := append(used, phase)
			*settings = append(*settings, phaseSetting)
		} else {
			updatedUsed := append(used, phase)
			allPhaseSettingSequence(updatedUsed, settings)
		}
	}
}

func runIntCodeComputer(instructions, phaseSetting []int) int {
	ea := make(chan int, 1) // must be buffered to receive final result
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	// This channel will receive a value each time an amplifier halts.
	halt := make(chan bool)

	// Start amplifiers in parallel.
	go IntCodeComputer2(instructions, ea, ab, halt)
	go IntCodeComputer2(instructions, ab, bc, halt)
	go IntCodeComputer2(instructions, bc, cd, halt)
	go IntCodeComputer2(instructions, cd, de, halt)
	go IntCodeComputer2(instructions, de, ea, halt)

	// Provide phase settings.
	ea <- phaseSetting[0]
	ab <- phaseSetting[1]
	bc <- phaseSetting[2]
	cd <- phaseSetting[3]
	de <- phaseSetting[4]

	// Send initial input signal.
	ea <- 0

	// Wait for all amplifiers to halt.
	for i := 0; i < 5; i++ {
		<-halt
	}

	// Read the final result.
	return <-ea
}

// idea comes from: https://github.com/GreenLightning/aoc19
func FindLargestOutput2(instructions []int) int {
	var phaseSettings [][]int
	allPhaseSettingSequence([]int{}, &phaseSettings)

	max := 0
	for _, phaseSetting := range phaseSettings {
		result := runIntCodeComputer(instructions, phaseSetting)
		if result > max {
			max = result
		}
	}
	return max
}
