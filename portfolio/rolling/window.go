package rolling

import "github.com/pedrocvaranda/portfolio-tracker/portfolio/models"

// Window contains training and testing slices for walk-forward evaluation.
type Window struct {
	Index int
	Train []models.OHLCV
	Test  []models.OHLCV
}

func (w Window) Empty() bool {
	return len(w.Train) == 0 || len(w.Test) == 0
}
