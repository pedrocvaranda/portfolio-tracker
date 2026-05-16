package backtesting

import "math"

// Metrics summarizes forecast quality.
type Metrics struct {
	RMSE                float64
	MAE                 float64
	Coverage            float64
	Bias                float64
	DirectionalAccuracy float64
}

func CalculateMetrics(forecasts []Forecast, actuals []float64) Metrics {
	if len(forecasts) == 0 || len(forecasts) != len(actuals) {
		return Metrics{}
	}
	var squared, absolute, covered, bias, direction float64
	for i, forecast := range forecasts {
		err := forecast.Expected - actuals[i]
		squared += err * err
		absolute += math.Abs(err)
		bias += err
		if actuals[i] >= forecast.Lower && actuals[i] <= forecast.Upper {
			covered++
		}
		if i > 0 {
			predMove := forecast.Expected - forecasts[i-1].Expected
			actualMove := actuals[i] - actuals[i-1]
			if (predMove == 0 && actualMove == 0) || predMove*actualMove > 0 {
				direction++
			}
		}
	}
	denom := float64(len(forecasts))
	dirDenom := math.Max(1, float64(len(forecasts)-1))
	return Metrics{RMSE: math.Sqrt(squared / denom), MAE: absolute / denom, Coverage: covered / denom, Bias: bias / denom, DirectionalAccuracy: direction / dirDenom}
}
