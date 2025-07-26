package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Pedro", "")
		want := "Hello, Pedro!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when nothing is supplyied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("João", "French")
		want := "Bonjour, João!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In a not configured language", func(t *testing.T) {
		got := Hello("João", "Viatnamese")
		want := "Hello, João!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
