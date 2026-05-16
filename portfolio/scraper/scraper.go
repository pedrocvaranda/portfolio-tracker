package scraper

import (
	"context"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Scraper fetches historical market candles from a provider.
type Scraper interface {
	Fetch(ctx context.Context, req Request) (Response, error)
}

// HistoryProvider is an alias kept for callers that prefer domain wording.
type HistoryProvider = Scraper

func candlesFromResponse(resp Response) []models.OHLCV {
	return resp.Candles
}
