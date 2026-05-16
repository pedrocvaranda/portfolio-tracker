package calibration

import "github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"

// Adjustment nudges parameters after backtesting feedback.
type Adjustment struct {
	DriftLearningRate      float64
	VolatilityLearningRate float64
}

func (a Adjustment) Apply(params montecarlo.Params, feedback Feedback) montecarlo.Params {
	driftRate := a.DriftLearningRate
	if driftRate == 0 {
		driftRate = 0.10
	}
	volRate := a.VolatilityLearningRate
	if volRate == 0 {
		volRate = 0.05
	}
	relative := feedback.RelativeError()
	params.Drift += driftRate * relative
	if relative < 0 {
		relative = -relative
	}
	params.Volatility += volRate * relative
	return params
}
