package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func mockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		},
	))
}

func TestRacer(t *testing.T) {
	t.Run("returning the fastest server", func(t *testing.T) {
		slowServer := mockServer(20 * time.Millisecond)
		fastServer := mockServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		got, err := Racer(slowServer.URL, fastServer.URL)
		want := fastServer.URL

    if err != nil {
      t.Fatalf("did not expect error %v", err)
    }

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

  t.Run("should return an error if server timeout", func(t *testing.T) {
    s := mockServer(25 * time.Millisecond)
    mt := 20 * time.Millisecond

    defer s.Close()

    _, err := ConfigurableRacer(s.URL, s.URL, mt)

    if err == nil {
      t.Error("should throw an error")
    }
  })
}
