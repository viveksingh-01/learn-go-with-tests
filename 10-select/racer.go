package racer

import (
	"net/http"
)

// We will further refactor and synchronize our code
// Since, we don't really care about the exact response times of the requests,
// we just want to know which one comes back first.

// To do this, we're going to introduce a new construct called 'select'
// which helps us synchronise processes really easily and clearly.
func Racer(a, b string) (winner string) {
	// select allows us to wait on multiple channels.
	// The first one to send a value "wins" and the code underneath the case is executed.
	select {
	// We use ping in our select to set up two channels, one for each of our URLs.
	// Whichever one writes to its channel first will have its code executed in the select,
	// which results in its URL being returned (and being the winner).
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// We have defined a function ping which creates a chan struct{} and returns it.
func ping(url string) chan struct{} {
	// Why struct{} and not another type like a bool?
	// a chan struct{} is the smallest data type available from a memory perspective
	// so we get no allocation versus a bool.
	// Since we are closing and not sending anything on the chan so we don't need to allocate anything
	ch := make(chan struct{})
	// we start a goroutine which will send a signal into that channel once we have completed http.Get(url)
	go func() {
		http.Get(url)
		// In our case, we don't care what type is sent to the channel,
		// we just want to signal we are done and closing the channel works perfectly!
		close(ch)
	}()
	return ch
}

// NOTE:
// Always make channels
// Notice how we have to use make when creating a channel; rather than say var ch chan struct{}.
// When you use var the variable will be initialised with the "zero" value of the type.
// So for string it is "", int it is 0, etc.
// For channels the zero value is nil and if you try and send to it with <- it will block forever
// because you cannot send to nil channels
