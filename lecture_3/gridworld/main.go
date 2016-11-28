package main

import (
	"fmt"
	"math"
)

const (
	DiscountFactor = 1
	NumIters       = 50
)

var InitialValues = map[State]float64{
	State{0, 0}: 6,
	State{0, 1}: 0,
	State{0, 2}: 7,
	State{0, 3}: -6,
	State{1, 0}: 8,
	State{1, 1}: 5,
	State{1, 2}: 12,
	State{1, 3}: 3,
	State{2, 0}: 5,
	State{2, 1}: 7,
	State{2, 2}: 9,
	State{2, 3}: -1,
	State{3, 0}: -3,
	State{3, 1}: 15,
	State{3, 2}: 2,
	State{3, 3}: 3,
}

func main() {
	values := InitialValues
	for i := 0; true; i++ {
		fmt.Println("Step", i)
		printValues(values)
		newValues := map[State]float64{}
		var maxDiff float64
		for state := range values {
			bestReward := math.Inf(-1)
			for action := North; action <= West; action++ {
				next, rew := state.Move(action)
				rew += values[next]
				if rew > bestReward {
					bestReward = rew
				}
			}
			if diff := math.Abs(values[state] - bestReward); diff > maxDiff {
				maxDiff = diff
			}
			newValues[state] = bestReward
		}
		values = newValues
		if maxDiff < 1e-5 {
			break
		}
	}
}

func printValues(v map[State]float64) {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			valStr := fmt.Sprintf("%0.2f", v[State{x, y}])
			if len(valStr) > 5 {
				valStr = valStr[:5]
			}
			for len(valStr) < 5 {
				valStr = " " + valStr + " "
			}
			if len(valStr) == 5 {
				valStr = " " + valStr
			}
			fmt.Printf(" %s ", valStr)
		}
		fmt.Println()
	}
	fmt.Println()
}
