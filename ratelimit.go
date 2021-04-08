package ratelimit

import "time"

type Ratelimiter struct {
	t *time.Ticker
}

func NewLimiter(requests uint32, seconds time.Duration) *Ratelimiter {
	return &Ratelimiter{t: time.NewTicker(seconds)}
}

func (r *Ratelimiter) tick(req uint32, sec time.Duration) {
	select {
	case <-r.t.C:

	}
}

func (r *Ratelimiter) Limit() {
	<-r.t.C
}
