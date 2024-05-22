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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

// NOTE: Benchmark tests can be done by running the command - `go test -bench="."`
