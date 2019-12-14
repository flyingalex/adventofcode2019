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
	return recursionCount(input, []string{"COM"}, 0)
}

func recursionPoint(input map[string][]string, orbits []string, point string, routes *[]string) bool {
	for _, orbit := range orbits {
		if orbit == point {
			return true
		}
		*routes = append(*routes, orbit)
		i := len(*routes) - 1
		found := recursionPoint(input, input[orbit], point, routes)
		if found {
			return true
		} else {
			*routes = (*routes)[:i]
		}
	}
	return false
}

func MinimumTransfers(point1, point2 string, input map[string][]string) int {
	var point1Routes []string
	found1 := recursionPoint(input, []string{"COM"}, point1, &point1Routes)
	var point2Routes []string
	found2 := recursionPoint(input, []string{"COM"}, point2, &point2Routes)
	sameIdx := 0
	if found1 && found2 {
		for i, oribit := range point1Routes {
			if oribit != point2Routes[i] {
				sameIdx = i
				break
			}
		}
	}
	return len(point2Routes) + len(point1Routes) - 2*sameIdx
}
