package models

import "strings"

// Asset represents a tradable instrument tracked by the portfolio.
type Asset struct {
	Ticker   string
	Name     string
	Currency string
	Exchange string
}

func NewAsset(ticker, name, currency string) Asset {
	return Asset{
		Ticker:   strings.ToUpper(strings.TrimSpace(ticker)),
		Name:     strings.TrimSpace(name),
		Currency: strings.ToUpper(strings.TrimSpace(currency)),
	}
}

func (a Asset) Valid() bool {
	return a.Ticker != "" && a.Currency != ""
}
