package ratelimit

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

// GlobalLimiter allows 100 events per second, with a burst capacity of 200.
// Adjust these to your desired "decillion" scale requirements.
var Limiter = rate.NewLimiter(rate.Every(time.Second/100), 200)

func Wait() {
	Limiter.Wait(context.Background())
}
