package backtesting

import (
	"testing"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

func TestCalculateMetrics(t *testing.T) {
	forecasts := []Forecast{{Expected: 10, Lower: 9, Upper: 11}, {Expected: 12, Lower: 11, Upper: 13}}
	actuals := []float64{11, 13}
	metrics := CalculateMetrics(forecasts, actuals)
	if metrics.MAE != 1 || metrics.Coverage != 1 {
		t.Fatalf("unexpected metrics: %+v", metrics)
	}
}

func TestBacktestRun(t *testing.T) {
	candles := []models.OHLCV{
		{Ticker: "AAPL", Timestamp: time.Now(), Open: 10, High: 11, Low: 9, Close: 10},
		{Ticker: "AAPL", Timestamp: time.Now().AddDate(0, 0, 1), Open: 12, High: 13, Low: 11, Close: 12},
	}
	result, err := NewEngine().Run("AAPL", []Forecast{{Expected: 10, Lower: 9, Upper: 11}, {Expected: 12, Lower: 11, Upper: 13}}, candles)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Comparisons) != 2 {
		t.Fatalf("comparisons = %d, want 2", len(result.Comparisons))
	}
}
