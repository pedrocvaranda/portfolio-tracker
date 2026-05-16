package calibration

import (
	"errors"
	"math"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"
)

// Historical estimates annualized drift and volatility from historical log returns.
type Historical struct {
	TradingDays float64
}

func (h Historical) Calibrate(candles []models.OHLCV) (montecarlo.Params, error) {
	if len(candles) < 2 {
		return montecarlo.Params{}, errors.New("at least two candles are required")
	}
	tradingDays := h.TradingDays
	if tradingDays == 0 {
		tradingDays = 252
	}
	returns := make([]float64, 0, len(candles)-1)
	for i := 1; i < len(candles); i++ {
		prev, current := candles[i-1].Price(), candles[i].Price()
		if prev <= 0 || current <= 0 {
			continue
		}
		returns = append(returns, math.Log(current/prev))
	}
	if len(returns) == 0 {
		return montecarlo.Params{}, errors.New("no valid returns available")
	}
	mu, sigma := returnStats(returns)
	last := candles[len(candles)-1]
	return montecarlo.Params{
		Ticker:      last.Ticker,
		StartPrice:  last.Price(),
		Drift:       mu * tradingDays,
		Volatility:  sigma * math.Sqrt(tradingDays),
		TradingDays: tradingDays,
	}, nil
}

func returnStats(values []float64) (float64, float64) {
	var sum float64
	for _, value := range values {
		sum += value
	}
	mean := sum / float64(len(values))
	var variance float64
	for _, value := range values {
		delta := value - mean
		variance += delta * delta
	}
	return mean, math.Sqrt(variance / float64(len(values)))
}
