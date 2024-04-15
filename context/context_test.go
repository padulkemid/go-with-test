package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true

	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error)  {
	s.written = true

	return 0, errors.New("not implemented?")
}

func (s *SpyResponseWriter) WriteHeader(sc int)  {
	s.written = true
}

type SpyStore struct {
	res string
	t   *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		log.Printf("LOG -- the response: %v \n", s.res)

		for _, c := range s.res {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	data := "hello, world!"

	t.Run("run server", func(t *testing.T) {
		store := &SpyStore{res: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, res.Body.String(), data)
		}
	})

	t.Run("cancel store if cancelled", func(t *testing.T) {
		store := &SpyStore{res: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		req = req.WithContext(cancelCtx)
		res := &SpyResponseWriter{}

		svr.ServeHTTP(res, req)

		if res.written {
			t.Error("should not write a response!")
		}
	})
}
