package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func GetInputs(path string) map[string][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputs := make(map[string][]string)
	for scanner.Scan() {
		text := scanner.Text()
		input := strings.Split(text, ")")
		inputs[input[0]] = append(inputs[input[0]], input[1])
	}

	return inputs
}

func TestCaculateOrbits(t *testing.T) {
	t.Run("test CaculateOrbits", func(t *testing.T) {
		inputs := GetInputs("./test.txt")
		got := CaculateOrbits(inputs)
		want := 42
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})
	t.Run("check CaculateOrbits", func(t *testing.T) {
		inputs := GetInputs("./input.txt")
		got := CaculateOrbits(inputs)
		fmt.Println(got)
	})
}
