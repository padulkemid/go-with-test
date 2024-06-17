package poker

import (
	"fmt"
	"io"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(dur time.Duration, amt int, to io.Writer)
}

type BlindAlerterFunc func(dur time.Duration, amt int, to io.Writer)

func (b BlindAlerterFunc) ScheduleAlertAt(dur time.Duration, amt int, to io.Writer) {
	b(dur, amt, to)
}

func Alerter(dur time.Duration, amt int,to io.Writer) {
	time.AfterFunc(dur, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amt)
	})
}
