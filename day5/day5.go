package main

import (
	"strconv"
	"strings"
)

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func splitInstruction(number int) []int {
	splited := []int{0, 0, 0, 0}
	str := reverse(strings.Split(strconv.Itoa(number), ""))
	for i, num := range str {
		splited[3-i], _ = strconv.Atoi(num)
	}
	return splited
}

func IntCodeComputer(instructions []int, input int) int {
	var result int
	i := 0
	positionMode := 0
	plusOpcode := 1
	var instructNumbers []int
	var firstIdx int
	var secondIdx int
	var outputIdx int
loop:
	for {
		// ex: 2 -> []int{0,0,0,1}
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
		switch instructNumbers[3] {
		// adds
		// multiplies
		case 1:
			fallthrough
		case 2:
			fistNumber, secondNumber := getParams()
			if instructNumbers[3] == plusOpcode {
				instructions[outputIdx] = fistNumber + secondNumber
			} else {
				instructions[outputIdx] = fistNumber * secondNumber
			}
			i += 4
			// inputs
		case 3:
			if instructNumbers[1] == positionMode {
				instructions[firstIdx] = input
			} else {
				instructions[i+1] = input
			}
			i += 2
			// outputs
		case 4:
			if instructNumbers[1] == positionMode {
				result = instructions[firstIdx]
			} else {
				result = firstIdx
			}
			if instructions[i+2] == 99 {
				break loop
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
		}
	}
	return result
}
