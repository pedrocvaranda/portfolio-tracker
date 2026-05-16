package storage

import (
	"context"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Storage persists portfolio data and market candles.
type Storage interface {
	SavePortfolio(ctx context.Context, p models.Portfolio) error
	GetPortfolio(ctx context.Context, id string) (models.Portfolio, error)
	SaveCandles(ctx context.Context, ticker string, candles []models.OHLCV) error
	GetCandles(ctx context.Context, ticker string) ([]models.OHLCV, error)
}
