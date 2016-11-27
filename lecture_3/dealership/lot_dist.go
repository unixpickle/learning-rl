package main

import "math/rand"

const (
	RentReward     = 10
	LotSampleCount = 10000
)

// LotDist stores the distribution of things that could
// happen to a lot given its starting inventory.
type LotDist struct {
	Rewards   []float64
	EndCounts [][]float64
}

// NewLotDist approximates a LotDist, given the Poisson
// intensities for returns and rentals.
func NewLotDist(returns, rentals float64, capacity int) *LotDist {
	res := &LotDist{
		Rewards:   make([]float64, capacity+1),
		EndCounts: make([][]float64, capacity+1),
	}
	for i := range res.EndCounts {
		res.EndCounts[i] = make([]float64, capacity+1)
	}

	startCounts := make([]float64, capacity+1)

	pReturns := NewPoisson(returns)
	pRentals := NewPoisson(rentals)
	for i := 0; i < LotSampleCount; i++ {
		startCount := rand.Intn(capacity + 1)
		var reward float64
		rent := pRentals.Sample()
		ret := pReturns.Sample()
		count := startCount
		for len(rent) > 0 && len(ret) > 0 {
			if rent[0] < ret[0] {
				rent = rent[1:]
				if count > 0 {
					reward += RentReward
					count--
				}
			} else {
				ret = ret[1:]
				count++
			}
		}
		for len(rent) > 0 {
			if count > 0 {
				reward += RentReward
				count--
			}
			rent = rent[1:]
		}
		for len(ret) > 0 {
			count++
			ret = ret[1:]
		}
		res.Rewards[startCount] += reward
		res.EndCounts[startCount][count]++
		startCounts[startCount]++
	}
	for i := range res.Rewards {
		res.Rewards[i] /= startCounts[i]
		for j := range res.EndCounts[i] {
			res.EndCounts[i][j] /= startCounts[i]
		}
	}
	return res
}
