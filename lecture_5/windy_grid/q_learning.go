package main

import "math/rand"

// QLearn performs one episode of Q-learning.
func QLearn(q QFunc, step, epsilon float64) {
	s := NewState()
	for s != nil {
		var a Action
		if rand.Float64() < epsilon {
			a = Action(rand.Intn(4))
		} else {
			a, _ = best(q, s)
		}
		rew, newS := s.Move(a)
		_, futureReward := best(q, newS)
		delta := rew - (q[ActionState{a, *s}] - futureReward)
		q[ActionState{a, *s}] += step * delta
		s = newS
	}
}

func best(q QFunc, s *State) (Action, float64) {
	if s == nil {
		return 0, 0
	}
	var max float64
	var maxA Action
	for a := Action(0); a < 4; a++ {
		v, ok := q[ActionState{a, *s}]
		if ok && (v >= max || a == 0) {
			max = v
			maxA = a
		}
	}
	return maxA, max
}
