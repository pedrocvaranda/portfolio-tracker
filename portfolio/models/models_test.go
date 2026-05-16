package models

import "testing"

func TestPortfolioMarketValue(t *testing.T) {
	asset := NewAsset("aapl", "Apple", "usd")
	p := NewPortfolio("p1", "Main", "USD")
	p.Cash = 10
	p.Positions = []Position{{Asset: asset, Quantity: 2, AvgCost: 90}}

	got := p.MarketValue(map[string]float64{"AAPL": 100})
	if got != 210 {
		t.Fatalf("MarketValue() = %v, want 210", got)
	}
}

func TestOHLCVValid(t *testing.T) {
	bar := OHLCV{Ticker: "AAPL", Open: 1, High: 2, Low: 1, Close: 1.5}
	if bar.Valid() {
		t.Fatal("bar without timestamp should be invalid")
	}
}
