package calibration

import (
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"
)

// Calibrator estimates simulation parameters from market data.
type Calibrator interface {
	Calibrate(candles []models.OHLCV) (montecarlo.Params, error)
}
