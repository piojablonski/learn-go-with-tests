package application

import (
	"io"
	"time"
)

type ScheduledAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

type BlindAlerter interface {
	ScheduleAlertAt(at time.Duration, amount int, w io.Writer)
}

type BlindAlerterFunc func(at time.Duration, amount int, w io.Writer)

func (b BlindAlerterFunc) ScheduleAlertAt(at time.Duration, amount int, w io.Writer) {
	b(at, amount, w)
}
