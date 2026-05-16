package models

import "time"

// SimulationPath contains one simulated future trajectory.
type SimulationPath struct {
	Prices []float64
}

// SimulationResult summarizes a Monte Carlo run.
type SimulationResult struct {
	Ticker       string
	StartPrice   float64
	Paths        []SimulationPath
	GeneratedAt  time.Time
	Confidence   float64
	ExpectedEnd  float64
	Percentile05 float64
	Percentile50 float64
	Percentile95 float64
}

func (r SimulationResult) Count() int {
	return len(r.Paths)
}
