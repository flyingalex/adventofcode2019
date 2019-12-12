package day6

func recursionCount(input map[string][]string, orbits []string, weight int) int {
	count := 0
	for _, orbit := range orbits {
		count += weight
		count += recursionCount(input, input[orbit], weight+1)
	}
	return count
}

func CaculateOrbits(input map[string][]string) int {
	orbits := input["COM"]
	count := 0
	for _, orbit := range orbits {
		count++
		count += recursionCount(input, input[orbit], 2)
	}
	return count
}
