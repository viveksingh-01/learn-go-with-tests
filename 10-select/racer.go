package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// We use 'Now()' method on 'time' to record just before we try and get the URL
	startA := time.Now()
	// Then we use http.Get to try and perform an HTTP GET request against the URL.
	// This function returns an http.Response and an error
	http.Get(a)
	// time.Since takes the start time and returns a time.Duration of the difference
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}
