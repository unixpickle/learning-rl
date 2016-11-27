package main

const LotCapacity = 20

type State struct {
	Lot1 int
	Lot2 int
}

func (s State) ApplyAction(action int) State {
	res := s
	res.Lot2 += action
	res.Lot1 -= action
	if res.Lot1 < 0 {
		res.Lot2 += res.Lot1
		res.Lot1 = 0
	} else if res.Lot1 > LotCapacity {
		res.Lot2 += res.Lot1 - LotCapacity
		res.Lot1 = LotCapacity
	}
	if res.Lot2 < 0 {
		res.Lot1 += res.Lot2
		res.Lot2 = 0
	} else if res.Lot2 > LotCapacity {
		res.Lot1 += res.Lot2 - LotCapacity
		res.Lot2 = LotCapacity
	}
	return res
}

type Transitions struct {
	Lot1 *LotDist
	Lot2 *LotDist
}

func (t *Transitions) Next(s1 State, action int) (next map[State]float64, imm float64) {
	morningState := s1.ApplyAction(action)
	imm = t.Lot1.Rewards[morningState.Lot1] + t.Lot2.Rewards[morningState.Lot2]
	next = map[State]float64{}
	for i := 0; i <= LotCapacity; i++ {
		iProb := t.Lot1.EndCounts[morningState.Lot1][i]
		if iProb == 0 {
			continue
		}
		for j := 0; j <= LotCapacity; j++ {
			jProb := t.Lot2.EndCounts[morningState.Lot2][j]
			if jProb != 0 {
				next[State{Lot1: i, Lot2: j}] = iProb * jProb
			}
		}
	}
	return
}
