package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	got := dictionary.Search("test")
	want := "this is just a test"
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
