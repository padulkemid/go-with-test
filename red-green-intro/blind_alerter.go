package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(dur time.Duration, amt int)
}

type BlindAlerterFunc func(dur time.Duration, amt int)

func (b BlindAlerterFunc) ScheduleAlertAt(dur time.Duration, amt int) {
	b(dur, amt)
}

func StdOutAlerter(dur time.Duration, amt int) {
	time.AfterFunc(dur, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amt)
	})
}
