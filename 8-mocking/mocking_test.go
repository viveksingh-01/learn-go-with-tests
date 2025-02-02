package main

import (
	"bytes"
	"testing"
)

const expectedCountdown = `3
2
1
Go!`

// If we can mock time.Sleep we can use dependency injection to use it instead of
// a "real" time.Sleep and then we can spy on the calls to make assertions on them.
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := expectedCountdown
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
