package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// Instead of testing the speeds of the websites one after another
	// We should be able to check both at the same time (when Go is great at concurrency)
	aDurationChannel := make(chan time.Duration)
	bDurationChannel := make(chan time.Duration)
	go func() {
		aDurationChannel <- measureResponseTime(a)
	}()
	go func() {
		bDurationChannel <- measureResponseTime(b)
	}()

	if <-aDurationChannel < <-bDurationChannel {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
