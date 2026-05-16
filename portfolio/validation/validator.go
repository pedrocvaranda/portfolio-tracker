package validation

import "github.com/pedrocvaranda/portfolio-tracker/portfolio/models"

// Issue describes one data quality problem.
type Issue struct {
	Code    string
	Message string
	Index   int
}

// Validator checks candles before calibration or simulation.
type Validator interface {
	Validate(candles []models.OHLCV) []Issue
}

// Composite runs multiple validators.
type Composite struct {
	Validators []Validator
}

func (c Composite) Validate(candles []models.OHLCV) []Issue {
	var issues []Issue
	for _, validator := range c.Validators {
		issues = append(issues, validator.Validate(candles)...)
	}
	return issues
}
