# Rolling Optimization

Rolling optimization, also called walk-forward optimization, evaluates a strategy over repeated train/test slices.

1. Split ordered historical candles into windows.
2. Calibrate or configure candidates using the training slice.
3. Score each candidate against the test slice.
4. Keep the candidate with the lowest objective score.

The `rolling.Split` function creates windows with configurable train size, test size, and step. The generic `rolling.Optimize` function accepts any candidate type, so it can optimize Monte Carlo parameters, validation thresholds, or strategy parameters.

Example:

```go
windows, _ := rolling.Split(candles, 252, 21, 21)
best, _ := rolling.Optimize(windows, candidates, func(w rolling.Window, c MyCandidate) (float64, error) {
    return scoreCandidate(w, c), nil
})
```
