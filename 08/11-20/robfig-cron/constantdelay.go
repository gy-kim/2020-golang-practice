package cron

import "time"

// ConstantDelaySchedule represents a simple recurring duty cycle, e.g. "Every 5 minutes".
// It does not support jobs more frequent than once a second.
type ConstantDelaySchedule struct {
	Delay time.Duration
}

// Every returns a contab Schedule that activates once every duration.
func Every(duration time.Duration) ConstantDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next returns the next time this should be run.
func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}
