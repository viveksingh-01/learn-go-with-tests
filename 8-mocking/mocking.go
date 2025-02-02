package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

// If we can mock 'time.Sleep' we can use dependency injection to use it instead of
// a "real" time.Sleep and then we can spy on the calls to make assertions on them.

// Here, we are creating an interface for Sleeper and then implementing it with
// DefaultSleeper and SpySleeper. This way, we can use the SpySleeper to spy on
// the calls to Sleep() and make assertions on them.
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type SpySleeper struct {
	Calls int
}

// DefaultSleeper implements the Sleeper interface and sleeps for 1 second.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
