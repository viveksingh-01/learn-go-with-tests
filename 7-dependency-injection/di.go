package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	// Here, Fprintf is used to write to the writer instead of fmt.Printf which writes to stdout
	// This allows us to test the output by passing in a buffer instead of stdout
	// This is an example of dependency injection

	// Fprintf takes a writer and a format string and writes to the writer
	// Since bytes.Buffer implements the writer interface, it can be passed in as the writer argument to Fprintf
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	// We get "cannot use os.Stdout as *bytes.Buffer value in argument to Greet" error
	// And fmt.Fprintf allows us to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.
	// therefore, if we change our code to use the more general purpose interface we can then
	// use it in both tests and in our application.
	Greet(os.Stdout, "Johnny")
}
