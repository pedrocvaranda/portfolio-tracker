package scraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

const yahooChartURL = "https://query1.finance.yahoo.com/v8/finance/chart/"

// Yahoo fetches candles from Yahoo Finance chart API.
type Yahoo struct {
	Client  *http.Client
	BaseURL string
}

func NewYahoo() *Yahoo {
	return &Yahoo{Client: &http.Client{Timeout: 10 * time.Second}, BaseURL: yahooChartURL}
}

func (y *Yahoo) Fetch(ctx context.Context, req Request) (Response, error) {
	if strings.TrimSpace(req.Ticker) == "" {
		return Response{}, errors.New("ticker is required")
	}
	client := y.Client
	if client == nil {
		client = http.DefaultClient
	}
	base := y.BaseURL
	if base == "" {
		base = yahooChartURL
	}

	u, err := url.Parse(base + url.PathEscape(strings.ToUpper(req.Ticker)))
	if err != nil {
		return Response{}, err
	}
	q := u.Query()
	if !req.Start.IsZero() || !req.End.IsZero() {
		start := req.Start
		if start.IsZero() {
			start = time.Now().AddDate(-1, 0, 0)
		}
		end := req.End
		if end.IsZero() {
			end = time.Now()
		}
		q.Set("period1", fmt.Sprint(start.Unix()))
		q.Set("period2", fmt.Sprint(end.Unix()))
	} else {
		if req.Range == "" {
			req.Range = "1y"
		}
		q.Set("range", req.Range)
	}
	if req.Interval == "" {
		req.Interval = "1d"
	}
	q.Set("interval", req.Interval)
	q.Set("events", "history")
	u.RawQuery = q.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return Response{}, err
	}
	httpReq.Header.Set("User-Agent", "portfolio-tracker/1.0")

	res, err := client.Do(httpReq)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return Response{}, fmt.Errorf("yahoo returned status %d", res.StatusCode)
	}

	var payload yahooResponse
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return Response{}, err
	}
	candles, err := payload.toCandles(strings.ToUpper(req.Ticker))
	if err != nil {
		return Response{}, err
	}
	return Response{Ticker: strings.ToUpper(req.Ticker), Provider: "yahoo", Candles: candles, Fetched: time.Now().UTC()}, nil
}

type yahooResponse struct {
	Chart struct {
		Result []struct {
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []*float64 `json:"open"`
					High   []*float64 `json:"high"`
					Low    []*float64 `json:"low"`
					Close  []*float64 `json:"close"`
					Volume []*int64   `json:"volume"`
				} `json:"quote"`
				AdjClose []struct {
					AdjClose []*float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error any `json:"error"`
	} `json:"chart"`
}

func (r yahooResponse) toCandles(ticker string) ([]models.OHLCV, error) {
	if r.Chart.Error != nil {
		return nil, fmt.Errorf("yahoo chart error: %v", r.Chart.Error)
	}
	if len(r.Chart.Result) == 0 || len(r.Chart.Result[0].Indicators.Quote) == 0 {
		return nil, errors.New("yahoo response did not include quotes")
	}
	result := r.Chart.Result[0]
	quote := result.Indicators.Quote[0]
	candles := make([]models.OHLCV, 0, len(result.Timestamp))
	for i, ts := range result.Timestamp {
		if !hasFloat(quote.Open, i) || !hasFloat(quote.High, i) || !hasFloat(quote.Low, i) || !hasFloat(quote.Close, i) {
			continue
		}
		bar := models.OHLCV{
			Ticker:    ticker,
			Timestamp: time.Unix(ts, 0).UTC(),
			Open:      *quote.Open[i],
			High:      *quote.High[i],
			Low:       *quote.Low[i],
			Close:     *quote.Close[i],
		}
		if hasInt(quote.Volume, i) {
			bar.Volume = *quote.Volume[i]
		}
		if len(result.Indicators.AdjClose) > 0 && hasFloat(result.Indicators.AdjClose[0].AdjClose, i) {
			bar.AdjClose = *result.Indicators.AdjClose[0].AdjClose[i]
		}
		candles = append(candles, bar)
	}
	return candles, nil
}

func hasFloat(values []*float64, i int) bool {
	return i >= 0 && i < len(values) && values[i] != nil
}

func hasInt(values []*int64, i int) bool {
	return i >= 0 && i < len(values) && values[i] != nil
}
