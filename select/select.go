package selects

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// func measureResponseTime(url string) time.Duration  {
//   start := time.Now()
//   _, err := http.Get(url)
//
//   if err != nil {
//     log.Fatalf("server got trouble")
//   }
//
//   return time.Since(start)
// }

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		_, err := http.Get(url)
		if err != nil {
			log.Fatal("server down")
		}

		close(ch)
	}()

	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func Racer(a, b string) (winner string, err error)  {
  return ConfigurableRacer(a, b, tenSecondTimeout)
}
