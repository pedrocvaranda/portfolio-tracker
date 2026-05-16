package validation

import (
	"math"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// OutlierDetector flags unusually large log returns.
type OutlierDetector struct {
	ZScore float64
}

func (o OutlierDetector) Validate(candles []models.OHLCV) []Issue {
	if len(candles) < 3 {
		return nil
	}
	threshold := o.ZScore
	if threshold == 0 {
		threshold = 4
	}
	returns := make([]float64, 0, len(candles)-1)
	indexes := make([]int, 0, len(candles)-1)
	for i := 1; i < len(candles); i++ {
		prev, current := candles[i-1].Price(), candles[i].Price()
		if prev > 0 && current > 0 {
			returns = append(returns, math.Log(current/prev))
			indexes = append(indexes, i)
		}
	}
	mu, sigma := meanStd(returns)
	if sigma == 0 {
		return nil
	}
	var issues []Issue
	for i, ret := range returns {
		z := math.Abs((ret - mu) / sigma)
		if z > threshold {
			issues = append(issues, Issue{Code: "outlier", Message: "return exceeds configured z-score threshold", Index: indexes[i]})
		}
	}
	return issues
}

func meanStd(values []float64) (float64, float64) {
	if len(values) == 0 {
		return 0, 0
	}
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
