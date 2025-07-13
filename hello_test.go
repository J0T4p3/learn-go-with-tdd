package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Pedro")
		want := "Hello, Pedro"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when nothing is supplyied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
