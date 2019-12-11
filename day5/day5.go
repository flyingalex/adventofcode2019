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

func DiagnostTool(instructions []int, input int) int {
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
		instructNumbers = splitInstruction(instructions[i])
		firstIdx = instructions[i+1]
		secondIdx = instructions[i+2]
		outputIdx = instructions[i+3]
		switch instructNumbers[3] {
		case 1:
			fallthrough
		case 2:
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
			if instructNumbers[3] == plusOpcode {
				instructions[outputIdx] = fistNumber + secondNumber
			} else {
				instructions[outputIdx] = fistNumber * secondNumber
			}
			i += 4
		case 3:
			if instructNumbers[1] == positionMode {
				instructions[firstIdx] = input
			} else {
				instructions[i+1] = input
			}
			i += 2
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
		}
	}
	return result
}
