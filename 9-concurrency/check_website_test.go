package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "what://beautiful.world"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"what://beautiful.world",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"what://beautiful.world":     false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// The benchmark tests CheckWebsites using a slice of one hundred urls and uses
// a new fake implementation of WebsiteChecker - slowStubWebsiteChecker
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	// We use b.ResetTimer() in this test to reset the time of our test before it actually runs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

// slowStubWebsiteChecker is deliberately slow.
// It uses time.Sleep to wait exactly 20 milliseconds and then it returns true.
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
