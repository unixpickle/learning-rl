package main

import "math/rand"

// Observable is an observable state.
type Observable struct {
	CurrentSum    int
	DealerShowing int
}

// State is the state of the game.
// This is simplified from the game in the lecture,
// because it does not include aces.
type State struct {
	Observable
	DealerSum int
}

// NewState creates a random starting state.
func NewState() *State {
	showing := randCard()
	sum := showing
	for dealerPolicy(sum) {
		c := randCard()
		sum += c
	}
	return &State{
		Observable: Observable{
			CurrentSum:    randCard(),
			DealerShowing: showing,
		},
		DealerSum: sum,
	}
}

// Timestep takes an action.
// If next is nil, the episode is over.
func (s *State) Timestep(a Action) (reward int, next *State) {
	if !a {
		if s.Observable.CurrentSum > s.DealerSum || s.DealerSum > 21 {
			return 1, nil
		} else if s.Observable.CurrentSum < s.DealerSum {
			return -1, nil
		}
		return 0, nil
	}
	newState := *s
	newState.Observable.CurrentSum += randCard()
	if newState.CurrentSum > 21 {
		return -1, nil
	}
	return 0, &newState
}

// Action is true for a hit, or false for a stick.
type Action bool

// A Policy decides what to do based on a state.
type Policy interface {
	Action(o Observable) Action
}

func dealerPolicy(sum int) Action {
	return sum < 21
}

func randCard() int {
	return rand.Intn(10) + 1
}
