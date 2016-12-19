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
	for i := 0; i < 100000; i++ {
		QLearn(val, 0.001, 0.01)
	}
	val.Policy().Print()
	fmt.Println("policy works:", policyWorks(val.Policy()))
}

func policyWorks(p Policy) bool {
	s := NewState()
	seen := map[State]bool{}
	for s != nil && !seen[*s] {
		seen[*s] = true
		_, s = s.Move(p[*s])
	}
	return s == nil
}
