package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/weakai/evolution"
)

var StateNames = []string{"Facebook", "Class1", "Class2", "Class3"}
var ActionNames = [][]string{
	[]string{"Facebook", "Quit"},
	[]string{"Study"},
	[]string{"Sleep", "Study"},
	[]string{"Pub", "Study"},
}
var Rewards = [][]float64{
	[]float64{-1, 0},
	[]float64{-2},
	[]float64{0, -2},
	[]float64{1, 10},
}
var Transitions = [][]map[int]float64{
	{map[int]float64{0: 1}, map[int]float64{1: 1}},
	{map[int]float64{2: 1}},
	{map[int]float64{}, map[int]float64{3: 1}},
	{map[int]float64{1: 0.2, 2: 0.4, 3: 0.4}, map[int]float64{}},
}

const DiscountFactor = 1

func main() {
	rand.Seed(time.Now().UnixNano())
	entity := PolicyEntity(randomPolicy())
	solver := evolution.Solver{
		StepCount:            200,
		StepSizeInitial:      1,
		StepSizeFinal:        0.00001,
		MaxPopulation:        10,
		MutateProbability:    0.5,
		CrossOverProbability: 0.1,
		SelectionProbability: 0.9,
		DFTradeoff:           evolution.LinearDFTradeoff(1, 1),
	}
	solution := solver.Solve([]evolution.Entity{entity})[0].(PolicyEntity)
	fmt.Println("End policy:")
	values := solution.valueFunction()
	for i, name := range StateNames {
		fmt.Printf(" %s (value=%.3f)\n", name, values[i])
		for j, prob := range solution[i] {
			fmt.Printf("  -> %s (%0.3f)\n", ActionNames[i][j], prob)
		}
	}
}

func transitionMatrix(policy [][]float64) *linalg.Matrix {
	mat := linalg.NewMatrix(len(StateNames), len(StateNames))
	for source := range StateNames {
		for action, actProb := range policy[source] {
			trans := Transitions[source][action]
			for dest, destProb := range trans {
				mat.Set(source, dest, destProb*actProb)
			}
		}
	}
	return mat
}

func immediateRewards(policy [][]float64) linalg.Vector {
	res := make(linalg.Vector, len(StateNames))
	for source := range StateNames {
		for act, actProb := range policy[source] {
			res[source] += actProb * Rewards[source][act]
		}
	}
	return res
}

func randomPolicy() [][]float64 {
	var res [][]float64
	for _, x := range ActionNames {
		var probs linalg.Vector
		for _ = range x {
			probs = append(probs, rand.NormFloat64())
		}
		var sum float64
		for i, y := range probs {
			probs[i] = math.Exp(y)
			sum += probs[i]
		}
		probs.Scale(1 / sum)
		res = append(res, probs)
	}
	return res
}
