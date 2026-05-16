package models

import "time"

// Position represents the quantity held for an asset.
type Position struct {
	Asset    Asset
	Quantity float64
	AvgCost  float64
}

func (p Position) MarketValue(price float64) float64 {
	return p.Quantity * price
}

func (p Position) CostBasis() float64 {
	return p.Quantity * p.AvgCost
}

// Portfolio groups assets and cash for a given owner or strategy.
type Portfolio struct {
	ID        string
	Name      string
	Currency  string
	Cash      float64
	Positions []Position
	CreatedAt time.Time
}

func NewPortfolio(id, name, currency string) Portfolio {
	return Portfolio{ID: id, Name: name, Currency: currency, CreatedAt: time.Now().UTC()}
}

func (p Portfolio) TotalCostBasis() float64 {
	total := p.Cash
	for _, position := range p.Positions {
		total += position.CostBasis()
	}
	return total
}

func (p Portfolio) MarketValue(prices map[string]float64) float64 {
	total := p.Cash
	for _, position := range p.Positions {
		price := prices[position.Asset.Ticker]
		total += position.MarketValue(price)
	}
	return total
}
