package fuzzies

import (
	"testing"
)

func TestDiceCoefficient(t *testing.T) {
	tests := []struct {
		s           string
		t           string
		coefficient float64
	}{
		{
			"night",
			"",
			0,
		},
		{
			"night",
			"night",
			1,
		},
		{
			"night",
			"nacht",
			0.25,
		},
		{
			"kitten",
			"sittin",
			0.4,
		},
	}

	for _, test := range tests {
		c := DiceCoefficient(test.s, test.t)
		if c != test.coefficient {
			t.Errorf(
				"Expected coefficient for '%s' and '%s' to be %f, got %f\n",
				test.s,
				test.t,
				test.coefficient,
				c,
			)
		}
	}
}

func BenchmarkDiceCoefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiceCoefficient("night", "nacht")
	}
}
