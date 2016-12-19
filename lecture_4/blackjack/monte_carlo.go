package main

func MonteCarlo(n int, p Policy) map[Observable]float64 {
	valFunc := map[Observable]float64{}
	m := map[Observable]float64{}
	for i := 0; i < n; i++ {
		recursiveMC(NewState(), p, valFunc, m)
	}
	return valFunc
}

func recursiveMC(s *State, p Policy, vf, n map[Observable]float64) float64 {
	if s == nil {
		return 0
	}
	a := p.Action(s.Observable)
	reward, next := s.Timestep(a)
	totalReward := recursiveMC(next, p, vf, n) + float64(reward)
	n[s.Observable]++
	rate := 1.0 / n[s.Observable]
	vf[s.Observable] += rate * (totalReward - vf[s.Observable])
	return totalReward
}
