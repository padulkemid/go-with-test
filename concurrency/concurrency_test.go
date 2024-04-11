package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "danger://not.sin"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("should check the bool returns", func(t *testing.T) {
		urls := []string{
			"http://google.com",
			"http://blog.peenoise.us",
			"danger://not.sin",
		}

		got := CheckWebsites(mockWebsiteChecker, urls)
		want := map[string]bool{
			"http://google.com":       true,
			"http://blog.peenoise.us": true,
			"danger://not.sin":        false,
		}

		t.Logf("this is the result: %v", got)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
