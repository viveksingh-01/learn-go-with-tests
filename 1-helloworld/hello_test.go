package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying to hello people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'Hello, world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in French", func(t *testing.T) {
		got := Hello("Max", "French")
		want := "Bonjour, Max"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in German", func(t *testing.T) {
		got := Hello("Antonia", "German")
		want := "Hallo, Antonia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper. By doing this, when it fails,
	// the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
