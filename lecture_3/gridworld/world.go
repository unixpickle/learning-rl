package main

type Action int

const (
	North Action = iota
	East
	South
	West
)

type State struct {
	X int
	Y int
}

func (s State) Move(a Action) (next State, reward float64) {
	if s.X == 0 && s.Y == 0 {
		return s, 0
	}
	res := s
	switch a {
	case North:
		res.Y--
		if res.Y < 0 {
			res.Y = 0
		}
	case East:
		res.X++
		if res.X > 3 {
			res.X = 3
		}
	case South:
		res.Y++
		if res.Y > 3 {
			res.Y = 3
		}
	case West:
		res.X--
		if res.X < 0 {
			res.X = 0
		}
	}
	return res, -1
}
