package rolling

import (
	"errors"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

// Split creates rolling train/test windows from ordered candles.
func Split(candles []models.OHLCV, trainSize, testSize, step int) ([]Window, error) {
	if trainSize <= 0 || testSize <= 0 || step <= 0 {
		return nil, errors.New("trainSize, testSize, and step must be positive")
	}
	if len(candles) < trainSize+testSize {
		return nil, errors.New("not enough candles for one window")
	}
	var windows []Window
	for start, index := 0, 0; start+trainSize+testSize <= len(candles); start, index = start+step, index+1 {
		train := append([]models.OHLCV(nil), candles[start:start+trainSize]...)
		test := append([]models.OHLCV(nil), candles[start+trainSize:start+trainSize+testSize]...)
		windows = append(windows, Window{Index: index, Train: train, Test: test})
	}
	return windows, nil
}
