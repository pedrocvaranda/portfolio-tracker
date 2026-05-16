package backtesting

import (
	"errors"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Forecast contains predicted intervals aligned with observed candles.
type Forecast struct {
	Expected float64
	Lower    float64
	Upper    float64
}

// Result stores observations and computed metrics.
type Result struct {
	Ticker      string
	Forecasts   []Forecast
	Actuals     []float64
	Metrics     Metrics
	Comparisons []Comparison
}

// Engine compares forecasts against actual market prices.
type Engine struct{}

func NewEngine() Engine { return Engine{} }

func (Engine) Run(ticker string, forecasts []Forecast, actualCandles []models.OHLCV) (Result, error) {
	if len(forecasts) == 0 || len(actualCandles) == 0 {
		return Result{}, errors.New("forecasts and actual candles are required")
	}
	if len(forecasts) != len(actualCandles) {
		return Result{}, errors.New("forecasts and actual candles must have the same length")
	}
	actuals := make([]float64, len(actualCandles))
	for i, candle := range actualCandles {
		actuals[i] = candle.Price()
	}
	comparisons := Compare(forecasts, actuals)
	return Result{Ticker: ticker, Forecasts: forecasts, Actuals: actuals, Metrics: CalculateMetrics(forecasts, actuals), Comparisons: comparisons}, nil
}
