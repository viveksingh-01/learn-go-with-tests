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

// ConfigurableSleeper is a struct that contains a duration and a sleep function
// that takes a duration as an argument.
// This is useful for testing because we can pass a SpyTime struct that records
// the duration it was called with.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

// Sleep calls the sleep function of the ConfigurableSleeper with the duration
// of the ConfigurableSleeper.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// This is an implementation of the Sleeper interface.
// It has a Sleep method that takes a duration and assigns it to the durationSlept field.
// This is useful for testing because we can check the duration that Sleep was called with.
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

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
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
