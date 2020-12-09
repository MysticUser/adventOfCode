package golang

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func containsInt(haystack []int, needle int) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func find(haystack []string, needle string) int {
	var key = -1
	for _, a := range haystack {
		key++
		if a == needle {
			return key
		}
	}
	return -1
}

func highest(haystack []int) int {
	var highest int
	for _, num := range haystack {
		if num > highest {
			highest = num
		}
	}
	return highest
}

func lowest(haystack []int) int {
	var lowest int
	for _, num := range haystack {
		if lowest == 0 || num < lowest {
			lowest = num
		}
	}
	return lowest
}
