package main

import (
	"bytes"
	"reflect"
	"testing"
)

const expectedCountdown = `3
2
1
Go!`

func TestCountdown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("print from 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := expectedCountdown
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})
}
