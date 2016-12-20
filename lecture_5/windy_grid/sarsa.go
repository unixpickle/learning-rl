package main

import "math/rand"

// Sarsa performs one episode of Sarsa(Lambda).
func Sarsa(q QFunc, lambda, step, epsilon float64) {
	trace := map[ActionState]float64{}
	s := NewState()
	a, _ := best(q, s)
	if rand.Float64() < epsilon {
		a = Action(rand.Intn(4))
	}
	for s != nil {
		for k := range trace {
			trace[k] *= lambda
		}
		trace[ActionState{a, *s}] += 1

		expected := q[ActionState{a, *s}]

		var rew float64
		rew, s = s.Move(a)
		if rand.Float64() < epsilon {
			a = Action(rand.Intn(4))
		} else {
			a, _ = best(q, s)
		}

		delta := rew - expected
		if s != nil {
			delta += q[ActionState{a, *s}]
		}
		for k, e := range trace {
			q[k] += step * e * delta
		}
	}
}
