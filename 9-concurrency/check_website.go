package concurrency

type WebsiteChecker func(string) bool

// We observed that our fake website-checker implementation takes 20ms for each request.
// The next URL isn't processed until the current one gets completed.
// Hence we can say it's a blocking operation.
func CheckWebsitesWithoutGoroutine(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

// Therefore, to make 'CheckWebsites' faster, instead of waiting for a website to respond
// before sending a request to the next website, we will tell our computer to make the
// next request while it is waiting.

// INTRODUCTION TO GOROUTINES
// Normally in Go when we call a function doSomething() we wait for it to
// return (even if it has no value to return, we still wait for it to finish).
// We say that this operation is blocking - it makes us wait for it to finish.
// An operation that does not block in Go will run in a separate process called a 'goroutine'.

// Think of a process as reading down the page of Go code from top to bottom,
// going 'inside' each function when it gets called to read what it does.
// When a separate process starts, it's like another reader begins reading inside the function,
// leaving the original reader to carry on going down the page.

// To tell Go to start a new goroutine we turn a function call into a go statement
// by putting the keyword go in front of it: 'go doSomething()'.

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		// The 'go' keyword tells Go to start a new goroutine to run the function
		// wc(url) concurrently with the rest of the function.
		go func() {
			results[url] = wc(url)
		}()
		// here, we use anonymous functions to call wc(url) and assign the result to results[url].
	}
	return results
}

// Anonymous functions in Go:
// Anonymous functions have a number of features which make them useful, two of which we're using above.
// Firstly, they can be executed at the same time that they're declared.
// Secondly they maintain access to the lexical scope in which they are defined - all the
// variables that are available at the point when you declare the anonymous function are
// also available in the body of the function.
