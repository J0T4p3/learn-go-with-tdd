package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Adding any amount of numbers", func(t *testing.T) {
		sum := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

		got := Sum(sum)
		want := 30

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, sum)
		}
	})

	t.Run("Adding none numbers", func(t *testing.T) {
		sum := []int{}

		got := Sum(sum)
		want := 0

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, sum)
		}
	})

	t.Run("Adding negative numbers", func(t *testing.T) {
		sum := []int{-1, -4}

		got := Sum(sum)
		want := -5

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, sum)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("New slice with the sum of all passed slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 3})
		want := []int{3, 5}

		if !slices.Equal(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}

	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		if !slices.Equal(got, want) {
			t.Helper()
			t.Errorf("Got %v want %v", got, want)
		}
	}

	t.Run("Sum all items of a slice except the first one", func(t *testing.T) {
		got := SumAllTails([]int{1, 5, 6, 2}, []int{2, 3}, []int{2, 3})
		want := []int{13, 3, 3}
		checkSums(t, got, want)
	})

	t.Run("Sum tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{})
		want := []int{0, 0, 0}
		checkSums(t, got, want)
	})

	t.Run("Sum tails of unary slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{3}, []int{100})
		want := []int{0, 0, 0}
		checkSums(t, got, want)
	})

}
