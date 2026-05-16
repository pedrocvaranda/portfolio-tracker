package backtesting

// Comparison stores one forecast-vs-actual row.
type Comparison struct {
	Index    int
	Expected float64
	Actual   float64
	Error    float64
	Covered  bool
}

func Compare(forecasts []Forecast, actuals []float64) []Comparison {
	limit := len(forecasts)
	if len(actuals) < limit {
		limit = len(actuals)
	}
	rows := make([]Comparison, limit)
	for i := 0; i < limit; i++ {
		rows[i] = Comparison{Index: i, Expected: forecasts[i].Expected, Actual: actuals[i], Error: forecasts[i].Expected - actuals[i], Covered: actuals[i] >= forecasts[i].Lower && actuals[i] <= forecasts[i].Upper}
	}
	return rows
}
