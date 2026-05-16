# Usage

Run all tests:

```bash
go test ./...
```

Run a local Monte Carlo demo:

```bash
go run ./cmd --mode demo --ticker AAPL --price 100 --days 30 --paths 500
```

Run the web interface:

```bash
go run ./cmd --mode web --addr :8080
```

Then open:

```text
http://localhost:8080
```

Fetch Yahoo Finance history and calibrate parameters:

```bash
go run ./cmd --mode fetch --ticker AAPL --range 1mo --interval 1d
```

Use the packages directly:

```go
engine := montecarlo.NewEngine()
result, err := engine.Run(montecarlo.Params{
    Ticker: "AAPL",
    StartPrice: 100,
    Drift: 0.08,
    Volatility: 0.22,
    Days: 30,
    Paths: 1000,
})
```

Yahoo Finance requests depend on network availability and the public endpoint behavior. For production usage, wrap the scraper with retries, rate limiting, and persistent caching.
