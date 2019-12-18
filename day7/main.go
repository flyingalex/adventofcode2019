package day7

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readFile("day7/input.txt")

	var program []int
	for _, value := range strings.Split(input, ",") {
		program = append(program, toInt(value))
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findBestSignal(program, []int{0, 1, 2, 3, 4}))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(findBestSignal(program, []int{5, 6, 7, 8, 9}))
	}
}

func findBestSignal(program []int, phaseValues []int) int {
	bestSignal := 0
	for _, phaseSettings := range allPermutations(phaseValues) {
		signal := emulateAmplifiers(program, phaseSettings)
		bestSignal = max(bestSignal, signal)
	}
	return bestSignal
}

func emulateAmplifiers(program []int, phaseSettings []int) int {
	// Set up the channels connecting the amplifiers.
	ea := make(chan int, 1) // must be buffered to receive final result
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	// This channel will receive a value each time an amplifier halts.
	halt := make(chan bool)

	// Start amplifiers in parallel.
	go emulate(program, ea, ab, halt)
	go emulate(program, ab, bc, halt)
	go emulate(program, bc, cd, halt)
	go emulate(program, cd, de, halt)
	go emulate(program, de, ea, halt)

	// Provide phase settings.
	ea <- phaseSettings[0]
	ab <- phaseSettings[1]
	bc <- phaseSettings[2]
	cd <- phaseSettings[3]
	de <- phaseSettings[4]

	// Send initial input signal.
	ea <- 0

	// Wait for all amplifiers to halt.
	for i := 0; i < 5; i++ {
		<-halt
	}

	// Read the final result.
	return <-ea
}

func emulate(program []int, input <-chan int, output chan<- int, halt chan<- bool) {
	// Copy the program into memory, so that we do not modify the original.
	memory := make([]int, len(program))
	copy(memory, program)

	ip := 0
	for {
		instruction := memory[ip]
		opcode := instruction % 100

		getParameter := func(offset int) int {
			parameter := memory[ip+offset]
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0: // position mode
				return memory[parameter]
			case 1: // immediate mode
				return parameter
			default:
				panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", ip, instruction, offset, mode))
			}
		}

		switch opcode {

		case 1: // ADD
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			memory[c] = a + b
			ip += 4

		case 2: // MULTIPLY
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			memory[c] = a * b
			ip += 4

		case 3: // INPUT
			a := memory[ip+1]
			memory[a] = <-input
			ip += 2

		case 4: // OUTPUT
			a := getParameter(1)
			output <- a
			ip += 2

		case 5: // JUMP IF TRUE
			a, b := getParameter(1), getParameter(2)
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParameter(1), getParameter(2)
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}

		case 7: // LESS THAN
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			if a < b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case 8: // EQUAL
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			if a == b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case 99: // HALT
			halt <- true
			return

		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", ip, instruction, opcode))
		}
	}
}

// Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
// Source: https://groups.google.com/d/msg/golang-nuts/PnLnr4bc9Wo/z9ZGv2DYxXoJ
func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func allPermutations(values []int) (result [][]int) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}
	for i, current := range values {
		others := make([]int, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, route := range allPermutations(others) {
			result = append(result, append(route, current))
		}
	}
	return
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
