package loop

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	t.Run("Dynamic character printing in lower", func(t *testing.T) {
		got := Loop("a", 3, "lower")
		want := "aaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Dynamic character printing in upper", func(t *testing.T) {
		got := Loop("a", 3, "upper")
		want := "AAA"
		assertCorrectMessage(t, got, want)
	})

	t.Run("None character printing", func(t *testing.T) {
		got := Loop("a", 0, "")
		want := ""
		assertCorrectMessage(t, got, want)
	})

	t.Run("Same characters in, same out without case identifier", func(t *testing.T) {
		got := Loop("b", 3, "")
		want := "bbb"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleLoop() {
	upperString := Loop("a", 5, "upper")
	fmt.Println(upperString)
	// Output: AAAAA
}

func BenchmarkLoop(b *testing.B) {
	for b.Loop() {
		Loop("a", 10, "lower")
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
