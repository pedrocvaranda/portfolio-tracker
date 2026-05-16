package montecarlo

import (
	"math"
	"math/rand"
)

// GBM advances one price step under Geometric Brownian Motion.
func GBM(price, drift, volatility, dt float64, rng *rand.Rand) float64 {
	if price <= 0 {
		return 0
	}
	shock := rng.NormFloat64()
	return price * math.Exp((drift-0.5*volatility*volatility)*dt+volatility*math.Sqrt(dt)*shock)
}
