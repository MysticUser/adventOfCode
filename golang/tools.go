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
