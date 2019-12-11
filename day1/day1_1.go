package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func SumFuelRequirements() int {
	total := 0
	file, err := os.Open("./day1_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, error := strconv.Atoi(scanner.Text())
		if error != nil {
			log.Fatal(error)
		} else {
			total += int(math.Floor(float64(number/3))) - 2
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total=%d", total)
	return total
}
