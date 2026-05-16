package montecarlo

import "testing"

func TestEngineRun(t *testing.T) {
	result, err := NewEngine().Run(Params{Ticker: "AAPL", StartPrice: 100, Drift: 0.05, Volatility: 0.2, Days: 10, Paths: 20, Seed: 1})
	if err != nil {
		t.Fatal(err)
	}
	if result.Count() != 20 {
		t.Fatalf("Count() = %d, want 20", result.Count())
	}
	if result.Percentile05 <= 0 || result.Percentile95 <= 0 {
		t.Fatal("percentiles should be positive")
	}
}

func TestGeneratePathLength(t *testing.T) {
	result, err := NewEngine().Run(Params{StartPrice: 50, Days: 3, Paths: 1, Seed: 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Paths[0].Prices) != 4 {
		t.Fatalf("path length = %d, want 4", len(result.Paths[0].Prices))
	}
}
