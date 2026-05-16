package validation

import (
	"fmt"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// BasicChecks validates ordering, required fields, duplicates, and simple price consistency.
type BasicChecks struct {
	MaxGap time.Duration
}

func (b BasicChecks) Validate(candles []models.OHLCV) []Issue {
	var issues []Issue
	seen := map[time.Time]bool{}
	for i, candle := range candles {
		if !candle.Valid() {
			issues = append(issues, Issue{Code: "invalid_candle", Message: "candle has missing or inconsistent OHLC values", Index: i})
		}
		if seen[candle.Timestamp] {
			issues = append(issues, Issue{Code: "duplicate_timestamp", Message: fmt.Sprintf("duplicate timestamp %s", candle.Timestamp.Format(time.RFC3339)), Index: i})
		}
		seen[candle.Timestamp] = true
		if i > 0 {
			if !candle.Timestamp.After(candles[i-1].Timestamp) {
				issues = append(issues, Issue{Code: "not_sorted", Message: "candles must be sorted ascending by timestamp", Index: i})
			}
			maxGap := b.MaxGap
			if maxGap == 0 {
				maxGap = 96 * time.Hour
			}
			if candle.Timestamp.Sub(candles[i-1].Timestamp) > maxGap {
				issues = append(issues, Issue{Code: "gap", Message: "large gap between candles", Index: i})
			}
		}
	}
	return issues
}
