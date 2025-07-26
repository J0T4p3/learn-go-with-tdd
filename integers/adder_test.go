package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(1, 1)

	got := sum
	expected := 2
	if got != expected {
		t.Errorf("Got '%d' expected '%d' ", got, expected)
	}
}
