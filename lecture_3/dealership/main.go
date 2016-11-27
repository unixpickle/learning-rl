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
	Epsilon   = 1e-3
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
}
