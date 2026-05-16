package calibration

import (
	"testing"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"
)

func TestHistoricalCalibrate(t *testing.T) {
	now := time.Now()
	candles := []models.OHLCV{
		{Ticker: "AAPL", Timestamp: now, Open: 100, High: 101, Low: 99, Close: 100},
		{Ticker: "AAPL", Timestamp: now.AddDate(0, 0, 1), Open: 102, High: 103, Low: 101, Close: 102},
		{Ticker: "AAPL", Timestamp: now.AddDate(0, 0, 2), Open: 101, High: 102, Low: 100, Close: 101},
	}
	params, err := Historical{}.Calibrate(candles)
	if err != nil {
		t.Fatal(err)
	}
	if params.StartPrice != 101 || params.Ticker != "AAPL" {
		t.Fatalf("unexpected params: %+v", params)
	}
}

func TestAdjustmentApply(t *testing.T) {
	params := Adjustment{}.Apply(montecarlo.Params{Drift: 0.01, Volatility: 0.20}, Feedback{Predicted: 90, Actual: 100})
	if params.Drift <= 0.01 || params.Volatility <= 0.20 {
		t.Fatalf("expected increased params, got %+v", params)
	}
}
