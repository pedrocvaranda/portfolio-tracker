package scraper

import (
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Request describes historical data desired from a market data provider.
type Request struct {
	Ticker   string
	Range    string
	Interval string
	Start    time.Time
	End      time.Time
}

// Response contains normalized candles and provider metadata.
type Response struct {
	Ticker   string
	Provider string
	Candles  []models.OHLCV
	Fetched  time.Time
}
