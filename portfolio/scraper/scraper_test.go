package scraper

import (
	"encoding/json"
	"testing"
)

func TestYahooResponseToCandles(t *testing.T) {
	payloadJSON := `{
		"chart": {
			"result": [{
				"timestamp": [1000],
				"indicators": {
					"quote": [{
						"open": [10],
						"high": [11],
						"low": [9.5],
						"close": [10.5],
						"volume": [100]
					}],
					"adjclose": [{"adjclose": [10.4]}]
				}
			}],
			"error": null
		}
	}`
	var payload yahooResponse
	if err := json.Unmarshal([]byte(payloadJSON), &payload); err != nil {
		t.Fatal(err)
	}

	candles, err := payload.toCandles("AAPL")
	if err != nil {
		t.Fatal(err)
	}
	if len(candles) != 1 || candles[0].Price() != 10.4 {
		t.Fatalf("unexpected candles: %+v", candles)
	}
}
