package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Policy map[State]Action

func (p Policy) Print() {
	for row := 0; row < Height; row++ {
		for col := 0; col < Width; col++ {
			if entry, ok := p[State{row, col}]; ok {
				fmt.Print(entry.String() + " ")
			} else {
				fmt.Print("? ")
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	val := QFunc{}
	for i := 0; i < 10000; i++ {
		Sarsa(val, 0.5, 0.01, 0.01)
	}
	fmt.Println("Sarsa(0.5):")
	val.Policy().Print()
	fmt.Println("win steps:", policyWinSteps(val.Policy()))

	fmt.Println()

	val = QFunc{}
	for i := 0; i < 100000; i++ {
		QLearn(val, 0.001, 0.01)
	}
	fmt.Println("Q-learning:")
	val.Policy().Print()
	fmt.Println("win steps:", policyWinSteps(val.Policy()))
}

func policyWinSteps(p Policy) int {
	s := NewState()
	seen := map[State]bool{}
	var n int
	for s != nil && !seen[*s] {
		seen[*s] = true
		_, s = s.Move(p[*s])
		n++
	}
	if s != nil {
		return -1
	}
	return n
}
