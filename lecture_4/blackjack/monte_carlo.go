package main

func MonteCarlo(n int, startRate, endRate float64, p Policy) map[Observable]float64 {
	valFunc := map[Observable]float64{}
	for i := 0; i < n; i++ {
		rate := startRate + (endRate-startRate)*float64(i)/float64(n)
		recursiveMC(NewState(), rate, p, valFunc)
	}
	return valFunc
}

func recursiveMC(s *State, rate float64, p Policy, vf map[Observable]float64) float64 {
	if s == nil {
		return 0
	}
	a := p.Action(s.Observable)
	reward, next := s.Timestep(a)
	totalReward := recursiveMC(next, rate, p, vf) + float64(reward)
	vf[s.Observable] += rate * (totalReward - vf[s.Observable])
	return totalReward
}
