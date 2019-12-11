package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDiagnostTool(t *testing.T) {
	t.Run("check splitInstruction", func(t *testing.T) {
		got := splitInstruction(76)
		//want := []int{0, 0, 1, 1}
		fmt.Printf("%#v", got)

	})
	t.Run("check DiagnostTool", func(t *testing.T) {
		file, err := os.Open("./input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var instructions []string
		for scanner.Scan() {
			text := scanner.Text()
			instructions = strings.Split(text, ",")
			break
		}

		var instructionsInt []int
		for _, num := range instructions {
			number, _ := strconv.Atoi(num)
			instructionsInt = append(instructionsInt, number)
		}

		fmt.Println(DiagnostTool(instructionsInt, 1))
	})
}
