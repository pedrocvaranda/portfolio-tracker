package validation

import (
	"testing"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

func TestBasicChecksDuplicate(t *testing.T) {
	now := time.Now()
	candles := []models.OHLCV{
		{Ticker: "AAPL", Timestamp: now, Open: 1, High: 2, Low: 1, Close: 1},
		{Ticker: "AAPL", Timestamp: now, Open: 1, High: 2, Low: 1, Close: 1},
	}
	issues := BasicChecks{}.Validate(candles)
	if len(issues) == 0 {
		t.Fatal("expected validation issues")
	}
}

func TestComposite(t *testing.T) {
	issues := Composite{Validators: []Validator{BasicChecks{}, OutlierDetector{}}}.Validate(nil)
	if len(issues) != 0 {
		t.Fatalf("got %d issues, want none", len(issues))
	}
}
