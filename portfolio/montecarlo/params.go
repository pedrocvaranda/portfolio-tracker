package montecarlo

import "time"

// Params configures a Monte Carlo run using annualized drift and volatility.
type Params struct {
	Ticker      string
	StartPrice  float64
	Drift       float64
	Volatility  float64
	Days        int
	Paths       int
	TradingDays float64
	Seed        int64
}

func (p Params) WithDefaults() Params {
	if p.Days <= 0 {
		p.Days = 252
	}
	if p.Paths <= 0 {
		p.Paths = 1000
	}
	if p.TradingDays <= 0 {
		p.TradingDays = 252
	}
	if p.Seed == 0 {
		p.Seed = time.Now().UnixNano()
	}
	return p
}
