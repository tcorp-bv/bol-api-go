package backoff

import (
	"time"
)

// Simple backoff interface
type Backoff interface {
	// Get the backoff duration
	Get(attempt uint) time.Duration
}

type exponentialBackoff struct {
	// The start duration to wait after the first Next() call
	Start time.Duration
	// The maximum duration to wait
	Max time.Duration
}

func NewExponentialBackoff(start time.Duration, max time.Duration) Backoff {
	return &exponentialBackoff{Start: start, Max: max}
}

func (b *exponentialBackoff) Get(attempt uint) time.Duration {
	d := b.Start
	for i := uint(0); i < attempt; i++ {
		d *= time.Duration(2)
		if d > b.Max {
			return b.Max
		}
	}
	return d
}
