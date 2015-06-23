package fuzzies

// https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient
func DiceCoefficient(s, t string) float64 {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}

	if s == t {
		return 1
	}

	sBigram := make(map[string]bool)
	tBigram := make(map[string]bool)

	for i := 0; i < len(s)-1; i++ {
		ssub := s[i : i+2]
		sBigram[ssub] = true
	}

	for i := 0; i < len(t)-1; i++ {
		tsub := t[i : i+2]
		tBigram[tsub] = true
	}

	intersection := 0

	for sb := range sBigram {
		if tBigram[sb] {
			intersection++
		}
	}

	// calculate dice coefficient
	total := len(sBigram) + len(tBigram)
	return float64(intersection*2) / float64(total)
}
