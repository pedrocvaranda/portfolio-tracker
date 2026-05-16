package storage

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

var ErrNotFound = errors.New("not found")

// Memory is a concurrency-safe in-memory storage implementation.
type Memory struct {
	mu         sync.RWMutex
	portfolios map[string]models.Portfolio
	candles    map[string][]models.OHLCV
}

func NewMemory() *Memory {
	return &Memory{portfolios: map[string]models.Portfolio{}, candles: map[string][]models.OHLCV{}}
}

func (m *Memory) SavePortfolio(_ context.Context, p models.Portfolio) error {
	if p.ID == "" {
		return errors.New("portfolio id is required")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.portfolios[p.ID] = p
	return nil
}

func (m *Memory) GetPortfolio(_ context.Context, id string) (models.Portfolio, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.portfolios[id]
	if !ok {
		return models.Portfolio{}, ErrNotFound
	}
	return p, nil
}

func (m *Memory) SaveCandles(_ context.Context, ticker string, candles []models.OHLCV) error {
	ticker = strings.ToUpper(strings.TrimSpace(ticker))
	if ticker == "" {
		return errors.New("ticker is required")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.candles[ticker] = append([]models.OHLCV(nil), candles...)
	return nil
}

func (m *Memory) GetCandles(_ context.Context, ticker string) ([]models.OHLCV, error) {
	ticker = strings.ToUpper(strings.TrimSpace(ticker))
	m.mu.RLock()
	defer m.mu.RUnlock()
	candles, ok := m.candles[ticker]
	if !ok {
		return nil, ErrNotFound
	}
	return append([]models.OHLCV(nil), candles...), nil
}
