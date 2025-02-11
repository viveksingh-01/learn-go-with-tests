package concurrency

type WebsiteChecker func(string) bool

// The new type, result has been made to associate the return value of the
// WebsiteChecker with the url being checked - it's a struct of string and bool.
// As we don't need either value to be named, each of them is anonymous within the struct;
// this can be useful in when it's hard to know what to name a value.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// We create our 'resultChannel' channel, using make in the same way as map.
	// chan result is the type of the channel - a channel of result.
	resultChannel := make(chan result)
	// Now when we iterate over the urls, instead of writing to the map directly we're sending a
	// result struct for each call to wc to the resultChannel with a send statement.
	// This uses the <- operator, taking a channel on the left and a value on the right:
	for _, url := range urls {
		go func() {
			// By sending the results into a channel, we can control the timing of each write into the
			// results map, ensuring that it happens one at a time.
			resultChannel <- result{url, wc(url)}
			// Although each of the calls of wc, and each send to the result channel, is happening
			// concurrently inside its own process, each of the results is being dealt with one at a
			// time as we take values out of the result channel with the receive expression.
		}()
	}
	// The next for loop iterates once for each of the urls.
	// Inside we're using a receive expression, which assigns a value received from a channel to a variable.
	// This also uses the <- operator, but with the two operands now reversed:
	// the channel is now on the right and the variable that we're assigning to is on the left:
	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
