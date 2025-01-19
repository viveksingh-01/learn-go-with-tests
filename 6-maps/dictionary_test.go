package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// Maps are unordered collections of key-value pairs.
	// They are declared using the map keyword, the key type, and the value type.
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
