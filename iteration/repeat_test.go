package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B)  {
  for i := 0; i < b.N; i++ {
    Repeat("a", 2)
  }
}

func ExampleRepeat()  {
  got := Repeat("a", 3)
  fmt.Println(got)
  // Output: aaa
}
