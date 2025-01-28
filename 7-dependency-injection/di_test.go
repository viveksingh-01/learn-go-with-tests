package main

import (
	"bytes"
	"testing"
)

// We want to write a function that greets someone, just like we did
// in the hello-world chapter but this time we are going to be testing the actual printing.

// But how can we test this? Calling fmt.Printf prints to stdout, which is pretty hard for us
// to capture using the testing framework.

// What we need to do is to be able to inject (which is just a fancy word for pass in)
// the dependency of printing.

// Our function doesn't need to care where or how the printing happens, so we should accept
// an interface rather than a concrete type.

// If we do that, we can then change the implementation to print to something we control so
// that we can test it. In "real life" you would inject in something that writes to stdout.

func TestGreet(t *testing.T) {
	// The Buffer type from the bytes package implements the Writer interface,
	// because it has the method Write(p []byte) (n int, err error).
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
