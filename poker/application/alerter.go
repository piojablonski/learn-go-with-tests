package application

import "time"

type ScheduledAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

type BlindAlerter interface {
	ScheduleAlertAt(at time.Duration, amount int)
}

type BlindAlerterFunc func(at time.Duration, amount int)

func (b BlindAlerterFunc) ScheduleAlertAt(at time.Duration, amount int) {
	b(at, amount)
}
