package adventofcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInputsInOneline(path, sep string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var inputs []string
	for scanner.Scan() {
		text := scanner.Text()
		inputs = strings.Split(text, sep)
		break
	}

	return inputs
}

func GetInputsNumInOneline(path, sep string) []int {
	inputs := GetInputsInOneline(path, sep)
	var inputNumbers []int
	for _, num := range inputs {
		number, _ := strconv.Atoi(num)
		inputNumbers = append(inputNumbers, number)
	}
	return inputNumbers
}
