package fuzzies

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		distance int
	}{
		{
			"kitten",
			"kitten",
			0,
		},
		{
			"kitten",
			"sitting",
			3,
		},
		{
			"",
			"kitten",
			6,
		},
		{
			"kitten",
			"",
			6,
		},
		{
			"test",
			"hest",
			1,
		},
		{
			"airplane",
			"aeroplane",
			2,
		},
	}

	for _, test := range tests {
		d := LevenshteinDistance(test.s, test.t)
		if d != test.distance {
			t.Errorf(
				"Expected distance between '%s' and '%s' to be %d, got %d\n",
				test.s,
				test.t,
				test.distance,
				d,
			)
		}
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LevenshteinDistance("kitten", "sitting")
	}
}
