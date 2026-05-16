package montecarlo

import "math/rand"

func GeneratePath(params Params, rng *rand.Rand) []float64 {
	params = params.WithDefaults()
	prices := make([]float64, params.Days+1)
	prices[0] = params.StartPrice
	dt := 1.0 / params.TradingDays
	for i := 1; i <= params.Days; i++ {
		prices[i] = GBM(prices[i-1], params.Drift, params.Volatility, dt, rng)
	}
	return prices
}
