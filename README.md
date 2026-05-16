# Portfolio Tracker

Portfolio Tracker is a Go application for collecting market data, validating price history, calibrating Monte Carlo parameters, running simulations, and backtesting forecast quality.

It includes both a CLI and an embedded web interface.

## Features

- Yahoo Finance scraper using the public chart endpoint
- OHLCV, asset, portfolio, and simulation models
- In-memory storage implementation
- Rolling window splitting and optimization helpers
- Validation checks for missing data, duplicate dates, gaps, and outliers
- Historical calibration and feedback-based parameter adjustment
- Monte Carlo simulation with Geometric Brownian Motion
- Backtesting metrics: RMSE, MAE, coverage, bias, and directional accuracy
- Embedded web UI with a Canvas chart and JSON API

## Requirements

- Go 1.22 or newer

## Quick Start

Clone the repository:

```bash
git clone https://github.com/pedrocvaranda/portfolio-tracker.git
cd portfolio-tracker
```

Run the tests:

```bash
go test ./...
```

Start the web interface:

```bash
go run ./cmd --mode web --addr :8080
```

Open:

```text
http://localhost:8080
```

Run a Monte Carlo simulation in the terminal:

```bash
go run ./cmd --mode demo --ticker AAPL --price 100 --days 30 --paths 500
```

Fetch recent Yahoo Finance data:

```bash
go run ./cmd --mode fetch --ticker AAPL --range 1mo --interval 1d
```

## Build

```bash
go build -o portfolio-tracker ./cmd
```

On Windows:

```powershell
go build -o portfolio-tracker.exe ./cmd
.\portfolio-tracker.exe --mode web --addr :8080
```

## Project Layout

```text
portfolio-tracker/
|-- cmd/                  CLI and web server entry point
|-- docs/                 Architecture and usage notes
`-- portfolio/            Core packages and embedded web interface
```

## API

The web server exposes:

- `GET /`: web interface
- `POST /api/simulate`: Monte Carlo simulation endpoint

Example payload:

```json
{
  "ticker": "AAPL",
  "startPrice": 100,
  "drift": 0.08,
  "volatility": 0.22,
  "days": 90,
  "paths": 1000,
  "seed": 42
}
```

## Documentation

- [Architecture](docs/architecture.md)
- [Rolling Optimization](docs/rolling_optimization.md)
- [Usage](docs/usage.md)

## License

MIT. See [LICENSE](LICENSE).
