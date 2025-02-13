package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Ideally, we don't want to be relying on external services to test our code because
// they can be slow, flaky and can't test edge cases
// In the standard library, there is a package called net/http/httptest which enables
// users to easily create a mock HTTP server.

func TestRacer(t *testing.T) {

	// httptest.NewServer takes an http.HandlerFunc which we are sending in via an anonymous function.
	// http.HandlerFunc is a type that looks like this: type HandlerFunc func(ResponseWriter, *Request)

	// This is also how we would write a real HTTP server in Go.
	// The only difference here is we are wrapping it in an httptest.NewServer which makes it easier
	// to use with testing, as it finds an open port to listen on and then we can close it
	// when we're done with our test.

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
