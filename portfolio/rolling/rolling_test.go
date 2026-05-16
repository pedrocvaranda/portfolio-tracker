package rolling

import (
	"testing"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

func TestSplit(t *testing.T) {
	candles := make([]models.OHLCV, 10)
	for i := range candles {
		candles[i] = models.OHLCV{Ticker: "AAPL", Timestamp: time.Now().AddDate(0, 0, i), Open: 1, High: 1, Low: 1, Close: 1}
	}
	windows, err := Split(candles, 5, 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(windows) != 4 {
		t.Fatalf("len(windows) = %d, want 4", len(windows))
	}
}

func TestOptimize(t *testing.T) {
	windows := []Window{{Index: 0}, {Index: 1}}
	candidates := []Candidate[int]{{Name: "a", Value: 3}, {Name: "b", Value: 1}}
	best, err := Optimize(windows, candidates, func(_ Window, value int) (float64, error) {
		return float64(value), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if best.Name != "b" || best.Score != 1 {
		t.Fatalf("unexpected best: %+v", best)
	}
}
