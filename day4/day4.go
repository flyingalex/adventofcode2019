package day4

import (
	"strconv"
	"strings"
)

func hasAdjacentNumber(number int) bool {
	numStr := strconv.Itoa(number)
	numStrs := strings.Split(numStr, "")
	for i, num := range numStrs[0 : len(numStrs)-1] {
		if num == numStrs[i+1] {
			return true
		}
	}
	return false
}

func notLargerGroupMatching(number int) bool {
	numStr := strconv.Itoa(number)
	numStrs := strings.Split(numStr, "")
	var numbers []int
	for _, num := range numStrs {
		convetedNumber, _ := strconv.Atoi(num)
		numbers = append(numbers, convetedNumber)
	}
	groupCount := make(map[int]int)
	groupNumbers := []int{}
	for _, num := range numbers {
		if groupCount[num] > 0 {
			groupCount[num]++
		} else {
			groupNumbers = append(groupNumbers, num)
			groupCount[num] = 1
		}
	}

	for _, num := range groupCount {
		if num == 2 {
			return true
		}
	}
	return false
}

func FindAllPasswords(min, max string) int {
	minNumber, _ := strconv.Atoi(min[0:1])
	maxNumber, _ := strconv.Atoi(max[0:1])
	maxist, _ := strconv.Atoi(max)
	minist, _ := strconv.Atoi(min)
	passwordsCount := 0
	validNumber := func(num int) bool {
		return num <= maxist
	}
	number1 := 0
	number2 := 0
	number3 := 0
	number4 := 0
	number5 := 0
	number6 := 0
	for i1 := minNumber; i1 <= maxNumber; i1++ {
		// break this loop if smaller than minNumber or bigger than maxNumber
		number1 = i1 * 100000
		if validNumber(number1) {
			for i2 := 0; i2 <= 9; i2++ {
				number2 = number1 + i2*10000
				if i2 >= i1 && validNumber(number2) {
					for i3 := 0; i3 <= 9; i3++ {
						number3 = number2 + i3*1000
						if i3 >= i2 && validNumber(number3) {
							for i4 := 0; i4 <= 9; i4++ {
								number4 = number3 + i4*100
								if i4 >= i3 && validNumber(number4) {
									for i5 := 0; i5 <= 9; i5++ {
										number5 = number4 + i5*10
										if i5 >= i4 && validNumber(number5) {
											for i6 := 0; i6 <= 9; i6++ {
												number6 = number5 + i6
												if i6 >= i5 &&
													validNumber(number6) &&
													number6 >= minist &&
													hasAdjacentNumber(number6) {
													passwordsCount++
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return passwordsCount
}

func FindMoreStrictAllPasswords(min, max string) int {
	minNumber, _ := strconv.Atoi(min[0:1])
	maxNumber, _ := strconv.Atoi(max[0:1])
	maxist, _ := strconv.Atoi(max)
	minist, _ := strconv.Atoi(min)
	passwordsCount := 0
	validNumber := func(num int) bool {
		return num <= maxist
	}
	number1 := 0
	number2 := 0
	number3 := 0
	number4 := 0
	number5 := 0
	number6 := 0
	for i1 := minNumber; i1 <= maxNumber; i1++ {
		// break this loop if smaller than minNumber or bigger than maxNumber
		number1 = i1 * 100000
		if validNumber(number1) {
			for i2 := 0; i2 <= 9; i2++ {
				number2 = number1 + i2*10000
				if i2 >= i1 && validNumber(number2) {
					for i3 := 0; i3 <= 9; i3++ {
						number3 = number2 + i3*1000
						if i3 >= i2 && validNumber(number3) {
							for i4 := 0; i4 <= 9; i4++ {
								number4 = number3 + i4*100
								if i4 >= i3 && validNumber(number4) {
									for i5 := 0; i5 <= 9; i5++ {
										number5 = number4 + i5*10
										if i5 >= i4 && validNumber(number5) {
											for i6 := 0; i6 <= 9; i6++ {
												number6 = number5 + i6
												if i6 >= i5 &&
													validNumber(number6) &&
													number6 >= minist &&
													hasAdjacentNumber(number6) &&
													notLargerGroupMatching(number6) {
													passwordsCount++
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return passwordsCount
}
