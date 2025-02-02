package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

// To improve our tests, we want to keep track of the calls to Sleep and Write in the Countdown function
// We can do this by creating a SpyCountdownOperations struct
// that has a slice of strings to keep track of the calls to Sleep and Write
// We can then implement the Sleep and Write methods on the SpyCountdownOperations struct

// SpyCountdownOperations is a struct that keeps track of the calls to Sleep and Write
type SpyCountdownOperations struct {
	Calls []string
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep appends the sleep operation to the Calls slice
// of the SpyCountdownOperations struct
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write appends the write operation to the Calls slice
// of the SpyCountdownOperations struct
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
