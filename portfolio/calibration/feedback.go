package calibration

// Feedback captures the difference between predicted and observed outcomes.
type Feedback struct {
	Predicted float64
	Actual    float64
}

func (f Feedback) Error() float64 {
	return f.Actual - f.Predicted
}

func (f Feedback) RelativeError() float64 {
	if f.Actual == 0 {
		return 0
	}
	return f.Error() / f.Actual
}
