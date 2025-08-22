package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Implements the Writer interface, returning total bytes
	// written and an error, if any.
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
