package main

import (
	c "mocking/countdown"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	c.Countdown(os.Stdout, &DefaultSleeper{})
}
