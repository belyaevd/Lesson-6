package functions

import "testing"

type testpair struct {
	values []float64
	sum    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestSumSet(t *testing.T) {
	for _, pair := range tests {
		v := sum(pair.values)
		if v != pair.sum {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", v,
			)
		}
	}
}
