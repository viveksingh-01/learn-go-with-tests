package main

import (
	"fmt"
	"io"
	"os"
)

const countdownStart = 3
const finalWord = "Go!"

// Countdown function to count from 3 to 1
// We know we want our Countdown function to write data somewhere and io.Writer is
// the de-facto way of capturing that as an interface in Go.
// In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// In test we will send to bytes.Buffer so our tests can capture what data is being generated.
func Countdown(out io.Writer) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
