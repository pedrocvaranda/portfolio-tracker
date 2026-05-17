# Portfolio Tracker

**Monte Carlo simulation engine for market data collection, price validation, parameter calibration, and forecast backtesting**

---

## What is This?

Portfolio Tracker is a Go application that answers a practical question: **"How reliable are probabilistic price forecasts, and how can they be continuously improved?"**

**Key Features:**

- **Yahoo Finance scraper** — pulls OHLCV history via the public chart endpoint
- **Monte Carlo simulation** — Geometric Brownian Motion with configurable drift and volatility
- **Validation pipeline** — detects missing data, duplicate dates, gaps, and price outliers
- **Historical calibration** — estimates GBM parameters from real data with feedback adjustment
- **Rolling optimization** — walk-forward window splitting to prevent lookahead bias
- **Backtesting metrics** — RMSE, MAE, coverage, bias, and directional accuracy
- **Embedded web UI** — Canvas chart and JSON API, no external server required
- **In-memory storage** — zero-dependency persistence layer

---

## Quick Start

### Requirements

- Go 1.22 or newer

### Installation

```bash
git clone https://github.com/pedrocvaranda/portfolio-tracker.git
cd portfolio-tracker
```

### Run Tests

```bash
go test ./...
```

### Web Interface

```bash
go run ./cmd --mode web --addr :8080
```

Open:

```text
http://localhost:8080
```

### Monte Carlo Demo

```bash
go run ./cmd --mode demo --ticker AAPL --price 100 --days 30 --paths 500
```

### Fetch Market Data

```bash
go run ./cmd --mode fetch --ticker AAPL --range 1mo --interval 1d
```

---

## Build

```bash
go build -o portfolio-tracker ./cmd
```

On Windows:

```powershell
go build -o portfolio-tracker.exe ./cmd
.\portfolio-tracker.exe --mode web --addr :8080
```

---

## Model

### Simulation Engine

The core engine runs Geometric Brownian Motion paths from a given start price:

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

### Backtesting Metrics

| Metric | Description |
| --- | --- |
| RMSE | Root mean squared error between simulated median and actual |
| MAE | Mean absolute error |
| Coverage | Fraction of actuals within the simulated confidence interval |
| Bias | Systematic over- or under-estimation |
| Directional Accuracy | Fraction of correctly predicted price directions |

### Rolling Optimization

Walk-forward validation splits historical data into repeated train/test windows. Each window calibrates parameters independently, preventing lookahead bias and producing out-of-sample performance estimates.

---

## Project Structure

```text
portfolio-tracker/
 cmd/ CLI and web server entry point
 docs/ Architecture and usage notes
 portfolio/ Core packages and embedded web interface
```

---

## API

The web server exposes:

- `GET /` — web interface
- `POST /api/simulate` — Monte Carlo simulation endpoint

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

---

## Documentation

- [Architecture](docs/architecture.md)
- [Rolling Optimization](docs/rolling_optimization.md)
- [Usage](docs/usage.md)

---

## About the Author

**Pedro Coutinho Varanda**

- **#1 Brazil** — National Astronomy Olympiad (OBA 2025, Perfect Score)
- **#2 Brazil** — OBA 2023
- **#3 Brazil** — OBA 2024
- **3x Selected** — International Olympiad on Astronomy and Astrophysics (IOAA)
- **4x Gold** — Canguru Mathematics Competition (2022–2025)

ML/AI enthusiast | Rio de Janeiro, Brazil

[GitHub](https://github.com/pedrocvaranda) • [ORCID](https://orcid.org/0009-0004-5199-1745) • [Email](mailto:pedrocvaranda@gmail.com)

---

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for local setup and guidelines.

---

## Related Projects

- [Cash Allocation Model](https://github.com/pedrocvaranda/modelo_alocacao_caixa) — ML-based capital allocation optimizer with Monte Carlo simulation
- [Varandian Optics Simulator](https://github.com/pedrocvaranda/varandian-optics-simulator) — Light propagation simulator in curved spaces

---

## License

MIT. See [LICENSE](LICENSE).

---

[![Go](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://golang.org) [![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE) [![Status](https://img.shields.io/badge/status-active-success.svg)]()