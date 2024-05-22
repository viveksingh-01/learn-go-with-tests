package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character N times", func(t *testing.T) {
		got := Repeat("a", 7)
		expected := "aaaaaaa"
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("repeat character 5 times if the repeatCount is less than or equal to 0", func(t *testing.T) {
		got := Repeat("a", 0)
		expected := "aaaaa"
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}
