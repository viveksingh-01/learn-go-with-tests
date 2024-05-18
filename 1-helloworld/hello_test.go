package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying to hello people", func(t *testing.T) {
		got := Hello("Chris");
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}	
	})
	t.Run("empty string defaults to 'Hello, world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}