package main

import "strconv"

const (
	Width  = 10
	Height = 7
)

// An Action represents a decision to move in a
// given direction.
type Action int

const (
	Up Action = iota
	Down
	Left
	Right
)

func (a Action) String() string {
	switch a {
	case Up:
		return "\u2191"
	case Down:
		return "\u2193"
	case Left:
		return "\u2190"
	case Right:
		return "\u2192"
	default:
		panic("unknown action: " + strconv.Itoa(int(a)))
	}
}

// A State stores the player's current position.
type State struct {
	Row int
	Col int
}

// NewState creates a start state.
func NewState() *State {
	return &State{Row: 3, Col: 0}
}

// Move applies an action to the state.
// The resulting target state will be nil if the goal
// has been reached.
func (s *State) Move(a Action) (reward float64, target *State) {
	blowAmount := []int{0, 0, 0, 1, 1, 1, 2, 2, 1, 0}[s.Col]
	res := *s
	res.Row -= blowAmount
	switch a {
	case Up:
		res.Row -= 1
	case Down:
		res.Row += 1
	case Left:
		res.Col -= 1
	case Right:
		res.Col += 1
	}
	res.Row = limitValue(res.Row, Height)
	res.Col = limitValue(res.Col, Width)
	if res.Row == 3 && res.Col == 7 {
		return 0, nil
	}
	return -1, &res
}

func limitValue(x, max int) int {
	if x < 0 {
		x = 0
	} else if x >= max {
		x = max-1
	}
	return x
}
