package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Interface who's implementing the Sleep method
type Sleeper interface {
	Sleep()
}

// Struct for "prod" code
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Struct for tests
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func Countdown(w io.Writer, s Sleeper) {
	const finalWord = "GO"
	const initialCountdown = 3
	for i := initialCountdown; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
		s.Sleep()
		if i == 1 {
			fmt.Fprintf(w, "GO")
		}
	}
}

func main() {
	sleeper := &ConfigurableSleeper{
		1 * time.Second,
		time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
