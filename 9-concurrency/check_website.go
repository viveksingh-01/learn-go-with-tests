package concurrency

type WebsiteChecker func(string) bool

// On running the benchmark test with our newly implemented Goroutine, we get error -
// fatal error: concurrent map writes.
// Sometimes, when we run our tests, two of the goroutines write to the results map
// at exactly the same time. Maps in Go don't like it when more than one thing tries
// to write to them at once, and so fatal error.

// This is a race condition, a bug that occurs when the output of our software is dependent
// on the timing and sequence of events that we have no control over.
// Because we cannot control exactly when each goroutine writes to the results map, we are
// vulnerable to two goroutines writing to it at the same time.
// We can solve this data race by coordinating our goroutines using CHANNELS.

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}
	return results
}

// CHANNELS in Go:
// Channels are a Go data structure that can both receive and send values.
// These operations, along with their details, allow communication between different processes.

// In this case we want to think about the communication between the parent process and each of
// the goroutines that it makes to do the work of running the WebsiteChecker function with the url.
