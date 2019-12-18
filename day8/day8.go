package day8

func DecodeImage(inputs []int) int {
	layerLength := 25 * 6
	i := 0
	zero := 0
	one := 0
	two := 0
	for {
		if i == len(inputs)-layerLength {
			break
		}
		a := 0
		b := 0
		c := 0
		for _, num := range inputs[i : i+layerLength] {
			switch num {
			case 0:
				a++
			case 1:
				b++
			case 2:
				c++
			}
		}
		if zero == 0 || zero > a {
			zero = a
			one = b
			two = c
		}
		i += layerLength
	}
	return one * two
}
