package slices

import (
	"testing"

	"golang.org/x/exp/slices"
)

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

func TestSumAll(t *testing.T) {
	t.Run("size should be summed", func(t *testing.T) {
		n1 := []int{1, 2}
		n2 := []int{0, 9}

		got := SumAll(n1, n2)
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("expected '%v' got '%v'", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum the tails only", func(t *testing.T) {
		n1 := []int{1, 2}
		n2 := []int{0, 9}

		got := SumAllTails(n1, n2)
		want := []int{2, 9}

		if !slices.Equal(got, want) {
			t.Errorf("expected '%v' got '%v'", want, got)
		}
	})

	t.Run("should sum more than once", func(t *testing.T) {
		n1 := []int{1, 2, 4, 6}
		n2 := []int{3, 5}
		n3 := []int{1}

		got := SumAllTails(n1, n2, n3)
		want := []int{12, 5, 1}

		if !slices.Equal(got, want) {
			t.Errorf("expected '%v' got '%v'", want, got)
		}
	})
}
