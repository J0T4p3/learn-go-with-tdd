package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to GO", func(t *testing.T) {
		spyCountdown := SpyCountDownOperations{}
		buf := bytes.Buffer{}
		Countdown(&buf, &spyCountdown)

		got := buf.String()
		want := "3\n2\n1\nGO"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		Countdown(
			spySleepPrinter,
			spySleepPrinter,
		)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v",
				want,
				spySleepPrinter.Calls,
			)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		sleepTime,
		spyTime.SetDurationSlept,
	}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf(
			"should have slept for %v, slept %v",
			sleepTime,
			spyTime.durationSlept,
		)
	}

}
