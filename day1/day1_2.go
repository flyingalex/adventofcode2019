package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func requiredFuel(number int) int {
	required := int(math.Floor(float64(number/3))) - 2
	if required > 0 {
		return required
	}
	return 0
}

func SumFuelRequirements() int {
	total := 0
	file, err := os.Open("./day1_2.txt")
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
			fuel := requiredFuel(number)
			for fuel != 0 {
				total += fuel
				fuel = requiredFuel(fuel)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total=%d", total)
	return total
}
