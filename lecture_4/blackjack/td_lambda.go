package main

import "math"

func TDLambda(n int, lambda, startRate, endRate float64, p Policy) map[Observable]float64 {
	valFunc := map[Observable]float64{}
	sr := math.Log(startRate)
	er := math.Log(endRate)
	for i := 0; i < n; i++ {
		r := sr + (er-sr)*float64(i)/float64(n)
		rate := math.Exp(r)
		trace := map[Observable]float64{}
		td(rate, p, valFunc, trace, lambda)
	}
	return valFunc
}

func td(rate float64, p Policy, vf, trace map[Observable]float64, lambda float64) {
	s := NewState()
	for s != nil {
		for state := range trace {
			trace[state] *= lambda
		}
		trace[s.Observable] += 1
		a := p.Action(s.Observable)
		reward, next := s.Timestep(a)
		diff := float64(reward) - vf[s.Observable]
		if next != nil {
			diff += vf[next.Observable]
		}
		for state, w := range trace {
			vf[state] += w * rate * diff
		}
		s = next
	}
}
