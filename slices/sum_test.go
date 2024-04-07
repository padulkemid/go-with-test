package slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		num := []int{1, 2, 3}

		got := Sum(num)
		want := 6

		if got != want {
			t.Errorf("expected '%d' got '%d' given %v", want, got, num)
		}
	})
}
