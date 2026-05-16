package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/calibration"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/scraper"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/validation"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/web"
)

func main() {
	mode := flag.String("mode", "demo", "demo, fetch, or web")
	ticker := flag.String("ticker", "AAPL", "asset ticker")
	price := flag.Float64("price", 100, "starting price for demo simulation")
	days := flag.Int("days", 30, "number of simulated days")
	paths := flag.Int("paths", 500, "number of Monte Carlo paths")
	rangeValue := flag.String("range", "1mo", "Yahoo Finance range")
	interval := flag.String("interval", "1d", "Yahoo Finance interval")
	addr := flag.String("addr", ":8080", "web server address")
	flag.Parse()

	switch *mode {
	case "demo":
		runDemo(*ticker, *price, *days, *paths)
	case "fetch":
		runFetch(*ticker, *rangeValue, *interval)
	case "web":
		runWeb(*addr)
	default:
		log.Fatalf("unknown mode %q", *mode)
	}
}

func runDemo(ticker string, price float64, days int, paths int) {
	result, err := montecarlo.NewEngine().Run(montecarlo.Params{Ticker: ticker, StartPrice: price, Drift: 0.08, Volatility: 0.22, Days: days, Paths: paths, Seed: 42})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s Monte Carlo (%d paths, %d days)\n", ticker, paths, days)
	fmt.Printf("Expected end: %.2f\n", result.ExpectedEnd)
	fmt.Printf("P05/P50/P95: %.2f / %.2f / %.2f\n", result.Percentile05, result.Percentile50, result.Percentile95)
}

func runFetch(ticker, rangeValue, interval string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	resp, err := scraper.NewYahoo().Fetch(ctx, scraper.Request{Ticker: ticker, Range: rangeValue, Interval: interval})
	if err != nil {
		log.Fatal(err)
	}
	issues := validation.Composite{Validators: []validation.Validator{validation.BasicChecks{}, validation.OutlierDetector{}}}.Validate(resp.Candles)
	params, err := calibration.Historical{}.Calibrate(resp.Candles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fetched %d candles for %s from %s\n", len(resp.Candles), resp.Ticker, resp.Provider)
	fmt.Printf("Validation issues: %d\n", len(issues))
	fmt.Printf("Calibrated drift %.4f volatility %.4f start %.2f\n", params.Drift, params.Volatility, params.StartPrice)
}

func runWeb(addr string) {
	server := web.MustNewServer()
	fmt.Printf("Portfolio Tracker web UI: http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Handler()))
}
