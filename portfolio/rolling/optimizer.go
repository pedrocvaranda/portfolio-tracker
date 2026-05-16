package rolling

import "math"

// Candidate is one possible configuration for a rolling optimization run.
type Candidate[T any] struct {
	Name   string
	Value  T
	Score  float64
	Window int
}

// Objective scores a candidate for a specific window. Lower is better.
type Objective[T any] func(Window, T) (float64, error)

func Optimize[T any](windows []Window, candidates []Candidate[T], objective Objective[T]) (Candidate[T], error) {
	best := Candidate[T]{Score: math.Inf(1)}
	for _, window := range windows {
		for _, candidate := range candidates {
			score, err := objective(window, candidate.Value)
			if err != nil {
				return best, err
			}
			if score < best.Score {
				best = candidate
				best.Score = score
				best.Window = window.Index
			}
		}
	}
	return best, nil
}
