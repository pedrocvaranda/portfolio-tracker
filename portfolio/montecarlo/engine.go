package montecarlo

import (
	"errors"
	"math/rand"
	"sort"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Engine runs Monte Carlo simulations.
type Engine struct{}

func NewEngine() Engine { return Engine{} }

func (Engine) Run(params Params) (models.SimulationResult, error) {
	params = params.WithDefaults()
	if params.StartPrice <= 0 {
		return models.SimulationResult{}, errors.New("start price must be positive")
	}
	if params.Volatility < 0 {
		return models.SimulationResult{}, errors.New("volatility cannot be negative")
	}
	rng := rand.New(rand.NewSource(params.Seed))
	paths := make([]models.SimulationPath, params.Paths)
	ends := make([]float64, params.Paths)
	for i := 0; i < params.Paths; i++ {
		prices := GeneratePath(params, rng)
		paths[i] = models.SimulationPath{Prices: prices}
		ends[i] = prices[len(prices)-1]
	}
	sort.Float64s(ends)
	return models.SimulationResult{
		Ticker:       params.Ticker,
		StartPrice:   params.StartPrice,
		Paths:        paths,
		GeneratedAt:  time.Now().UTC(),
		Confidence:   0.90,
		ExpectedEnd:  mean(ends),
		Percentile05: percentile(ends, 0.05),
		Percentile50: percentile(ends, 0.50),
		Percentile95: percentile(ends, 0.95),
	}, nil
}

func mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var total float64
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

func percentile(sorted []float64, p float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	if p <= 0 {
		return sorted[0]
	}
	if p >= 1 {
		return sorted[len(sorted)-1]
	}
	idx := int(p * float64(len(sorted)-1))
	return sorted[idx]
}
