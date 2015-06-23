package fuzzies

// https://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(s, t string) int {
	if s == t {
		return 0
	}

	if len(s) == 0 {
		return len(t)
	}

	if len(t) == 0 {
		return len(s)
	}

	v0 := make([]int, len(t)+1)
	v1 := make([]int, len(t)+1)

	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < len(s); i++ {
		v1[0] = i + 1

		for j := 0; j < len(t); j++ {
			cost := 1
			if s[i] == t[j] {
				cost = 0
			}

			v1[j+1] = minimum(v1[j]+1, v0[j+1]+1, v0[j]+cost)
		}

		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}
	}

	return v1[len(t)]
}

func minimum(val ...int) int {
	min := val[0]
	for _, v := range val {
		if v < min {
			min = v
		}
	}
	return min
}
