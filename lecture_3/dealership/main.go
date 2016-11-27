package main

import (
	"fmt"
	"log"
	"math"
)

const (
	MaxMove   = 5
	MaxIters  = 100
	Discount  = 0.9
	Epsilon   = 1e-4
	InitValue = 500
)

func main() {
	log.Println("Building distributions...")
	trans := &Transitions{
		Lot1: NewLotDist(3, 3, LotCapacity),
		Lot2: NewLotDist(2, 4, LotCapacity),
	}
	log.Println("Initializing values...")
	values := map[State]float64{}
	for i := 0; i <= LotCapacity; i++ {
		for j := 0; j <= LotCapacity; j++ {
			values[State{i, j}] = InitValue
		}
	}
	log.Println("Value iteration...")
	for k := 0; k < MaxIters; k++ {
		var maxChange float64
		newVals := map[State]float64{}
		for state, oldValue := range values {
			var maxReward float64
			for action := -MaxMove; action <= MaxMove; action++ {
				next, value := trans.Next(state, action)
				for n, p := range next {
					value += Discount * p * values[n]
				}
				if value > maxReward {
					maxReward = value
				}
			}
			if change := math.Abs(maxReward - oldValue); change > maxChange {
				maxChange = change
			}
			newVals[state] = maxReward
		}
		log.Println("Iteration", k, "max change:", maxChange)
		values = newVals
		if maxChange < Epsilon {
			break
		}
	}
	fmt.Println("value of 0,0", values[State{Lot1: 0, Lot2: 0}])
	fmt.Println("value of 20,20", values[State{Lot1: 20, Lot2: 20}])

	fmt.Println("y=Lot1, x=Lot2")
	policy := optimalPolicy(trans, values)
	for i := 0; i <= LotCapacity; i++ {
		for j := 0; j <= LotCapacity; j++ {
			move := policy[State{Lot2: j, Lot1: LotCapacity - i}]
			moveStr := fmt.Sprintf("%d", move)
			if len(moveStr) == 1 {
				moveStr = " " + moveStr
			}
			fmt.Printf("%s ", moveStr)
		}
		fmt.Println()
	}
}

func optimalPolicy(trans *Transitions, values map[State]float64) map[State]int {
	res := map[State]int{}
	for state := range values {
		var bestValue float64
		var bestAction int
		for i := 0; i <= MaxMove; i++ {
			for s := -1; s <= 1; s += 2 {
				action := i * s
				next, reward := trans.Next(state, action)
				for ns, np := range next {
					reward += np * values[ns]
				}
				if reward > bestValue {
					bestValue = reward
					bestAction = action
				}
			}
		}
		res[state] = bestAction
	}
	return res
}
