package models

import "time"

// OHLCV stores one market candle.
type OHLCV struct {
	Ticker    string
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	AdjClose  float64
	Volume    int64
}

func (o OHLCV) Price() float64 {
	if o.AdjClose > 0 {
		return o.AdjClose
	}
	return o.Close
}

func (o OHLCV) Valid() bool {
	return o.Ticker != "" && !o.Timestamp.IsZero() && o.Open > 0 && o.High > 0 && o.Low > 0 && o.Close > 0 && o.Low <= o.High
}
