package main

import (
	"math"
	"math/rand"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/num-analysis/linalg/ludecomp"
	"github.com/unixpickle/weakai/evolution"
)

type PolicyEntity [][]float64

func (p PolicyEntity) Fitness() float64 {
	values := p.valueFunction()
	var sum float64
	for _, x := range values {
		sum += x
	}
	return sum
}

func (p PolicyEntity) Similarity(e []evolution.Entity) float64 {
	// TODO: this.
	return 1
}

func (p PolicyEntity) Mutate(stepSize float64) evolution.Entity {
	result := PolicyEntity{}
	for _, x := range p {
		mutated := make(linalg.Vector, len(x))
		copy(mutated, x)
		var sum float64
		for i, k := range mutated {
			mutated[i] = k * ((math.Abs(rand.NormFloat64()) * stepSize) + 1)
			sum += mutated[i]
		}
		mutated.Scale(1 / sum)
		result = append(result, mutated)
	}
	return result
}

func (p PolicyEntity) CrossOver(e evolution.Entity) evolution.Entity {
	result := PolicyEntity{}
	p1 := e.(PolicyEntity)
	for i := range p {
		if rand.Intn(2) == 0 {
			result = append(result, p[i])
		} else {
			result = append(result, p1[i])
		}
	}
	return result
}

func (p PolicyEntity) valueFunction() linalg.Vector {
	imm := immediateRewards(p)
	trans := transitionMatrix(p)
	trans = linalg.NewMatrixIdentity(trans.Rows).Add(trans.Scale(-DiscountFactor))
	return ludecomp.Decompose(trans).Solve(imm)
}
