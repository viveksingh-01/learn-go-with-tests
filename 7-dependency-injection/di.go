package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	// Here, Fprintf is used to write to the writer instead of fmt.Printf which writes to stdout
	// This allows us to test the output by passing in a buffer instead of stdout
	// This is an example of dependency injection

	// Fprintf takes a writer and a format string and writes to the writer
	// Since bytes.Buffer implements the writer interface, it can be passed in as the writer argument to Fprintf
	fmt.Fprintf(writer, "Hello, %s", name)
}
