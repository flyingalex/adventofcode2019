package day7

func IntCodeComputer2(program []int, input <-chan int, output chan<- int, halt chan<- bool) {
	instructions := make([]int, len(program))
	copy(instructions, program)
	i := 0
	positionMode := 0
	plusOpcode := 1
	var instructNumbers []int
	var firstIdx int
	var secondIdx int
	var outputIdx int
	for {
		// ex: 2 -> []int{0,0,0,1}
		opcode := instructions[i] % 100
		instructNumbers = splitInstruction(instructions[i])
		firstIdx = instructions[i+1]
		secondIdx = instructions[i+2]
		outputIdx = instructions[i+3]
		getParams := func() (int, int) {
			var fistNumber int
			var secondNumber int
			if instructNumbers[1] == positionMode {
				fistNumber = instructions[firstIdx]
			} else {
				fistNumber = firstIdx
			}
			if instructNumbers[0] == positionMode {
				secondNumber = instructions[secondIdx]
			} else {
				secondNumber = secondIdx
			}
			return fistNumber, secondNumber
		}
		switch opcode {
		// adds
		// multiplies
		case 1:
			fallthrough
		case 2:
			fistNumber, secondNumber := getParams()
			if opcode == plusOpcode {
				instructions[outputIdx] = fistNumber + secondNumber
			} else {
				instructions[outputIdx] = fistNumber * secondNumber
			}
			i += 4
			// inputs
		case 3:
			instructions[instructions[i+1]] = <-input
			i += 2
			// outputs
		case 4:
			if instructNumbers[1] == positionMode {
				output <- instructions[firstIdx]
			} else {
				output <- firstIdx
			}
			i += 2
		case 5:
			fistNumber, secondNumber := getParams()
			if fistNumber != 0 {
				i = secondNumber
			} else {
				i += 3
			}
			// jump-if-false
		case 6:
			fistNumber, secondNumber := getParams()
			if fistNumber == 0 {
				i = secondNumber
			} else {
				i += 3
			}
			// less than
		case 7:
			fistNumber, secondNumber := getParams()
			if fistNumber < secondNumber {
				instructions[outputIdx] = 1
			} else {
				instructions[outputIdx] = 0
			}
			i += 4
			// equals
		case 8:
			fistNumber, secondNumber := getParams()
			if fistNumber == secondNumber {
				instructions[outputIdx] = 1
			} else {
				instructions[outputIdx] = 0
			}
			i += 4
		case 99:
			halt <- true
			return
		}
	}
}
