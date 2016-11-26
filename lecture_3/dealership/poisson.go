package main

import (
	"math"
	"math/rand"
)

// Poisson implements a Poisson process.
type Poisson struct {
	lambda float64
}

// NewPoisson creates a Poisson process.
func NewPoisson(lambda float64) *Poisson {
	return &Poisson{
		lambda: lambda,
	}
}

// Sample samples the events in a Poisson process from
// time 0 to time 1.
func (p *Poisson) Sample() []float64 {
	var t float64
	var res []float64
	for {
		num := -math.Log((1 - rand.Float64())) / p.lambda
		if t+num > 1 {
			break
		}
		t += num
		res = append(res, t)
	}
	return res
}
