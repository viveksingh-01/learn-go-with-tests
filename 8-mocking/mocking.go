package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown function to count from 3 to 1
// We know we want our Countdown function to write data somewhere and io.Writer is
// the de-facto way of capturing that as an interface in Go.
// In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// In test we will send to bytes.Buffer so our tests can capture what data is being generated.

func Countdown(out io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
