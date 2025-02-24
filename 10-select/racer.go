package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
	// time.After is a very handy function when using select.
	// Although it didn't happen in our case, we can potentially write code that blocks forever
	// if the channels you're listening on never return a value. time.After returns a chan (like ping)
	// and will send a signal down it after the amount of time you define.

	// For us this is perfect; if a or b manage to return they win, but if we get to 10 seconds
	// then our time.After will send a signal and we'll return an error.
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	// we start a goroutine which will send a signal into that channel once we have completed http.Get(url)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
