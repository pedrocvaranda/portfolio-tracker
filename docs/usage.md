# Usage

**All entry points go through `cmd/`. Choose a mode depending on what you need.**

---

## Run Tests

```bash
go test ./...
```

---

## Monte Carlo Demo

Run a local simulation and print results to the terminal:

```bash
go run ./cmd --mode demo --ticker AAPL --price 100 --days 30 --paths 500
```

---

## Web Interface

Start the embedded server:

```bash
go run ./cmd --mode web --addr :8080
```

Then open:

```text
http://localhost:8080
```

---

## Fetch and Calibrate

Pull Yahoo Finance history and calibrate GBM parameters:

```bash
go run ./cmd --mode fetch --ticker AAPL --range 1mo --interval 1d
```

---

## Programmatic Usage

Use the packages directly in your own Go code:

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

---

## Notes

- Yahoo Finance requests depend on network availability and the public endpoint behavior.
- For production usage, wrap the scraper with retries, rate limiting, and persistent caching.
