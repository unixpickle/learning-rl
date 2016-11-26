package main

import (
	"testing"

	"github.com/unixpickle/approb"
)

func TestPoisson(t *testing.T) {
	dist1 := func() float64 {
		return approb.Poisson(2.5)
	}
	sampler := NewPoisson(2.5)
	dist2 := func() float64 {
		return float64(len(sampler.Sample()))
	}
	corr := approb.Correlation(10000, 0.1, dist1, dist2)
	if corr < 0.99 {
		t.Errorf("expected correlation 1 but got %f", corr)
	}
}
