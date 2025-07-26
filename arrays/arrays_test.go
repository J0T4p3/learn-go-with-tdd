package arrays

import "testing"

func TestArraySum(t *testing.T) {
	t.Run("Adding any amount of numbers", func(t *testing.T) {
		sum := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

		got := ArraySum(sum)
		want := 30

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, sum)
		}
	})

	t.Run("Adding none numbers", func(t *testing.T) {
		sum := []int{}

		got := ArraySum(sum)
		want := 0

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, sum)
		}
	})
}
