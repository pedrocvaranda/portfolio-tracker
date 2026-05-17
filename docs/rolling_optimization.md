# Rolling Optimization

**Walk-forward optimization evaluates a strategy over repeated train/test slices, producing honest out-of-sample performance estimates.**

---

## How It Works

1. Split ordered historical candles into overlapping windows
2. Calibrate or configure candidates using the **training slice**
3. Score each candidate against the **test slice**
4. Keep the candidate with the lowest objective score

This prevents lookahead bias: each window's test period is always strictly in the future relative to its training period.

---

## API

### Splitting Windows

```go
windows, _ := rolling.Split(candles, 252, 21, 21)
// args: candles, trainSize, testSize, step
```

### Optimizing Candidates

```go
best, _ := rolling.Optimize(windows, candidates, func(w rolling.Window, c MyCandidate) (float64, error) {
 return scoreCandidate(w, c), nil
})
```

The generic `rolling.Optimize` function accepts any candidate type, so it can optimize Monte Carlo parameters, validation thresholds, or custom strategy parameters.

---

## Parameters

| Parameter | Description |
| --- | --- |
| `trainSize` | Number of candles in the training slice (e.g. 252 trading days 1 year) |
| `testSize` | Number of candles in the test slice (e.g. 21 1 month) |
| `step` | How far to advance between windows |
