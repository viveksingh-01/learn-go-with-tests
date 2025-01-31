package main

import (
	"bytes"
	"testing"
)

const expectedCountdown = `3
2
1
Go!`

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	Countdown(&buffer)
	got := buffer.String()
	want := expectedCountdown
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
