package structs

import (
	"reflect"
	"testing"
)

const (
	intErrorMessage   = "got '%d' want '%d'"
	floatErrorMessage = "got '%g' want '%g'"
)

func assertNum[T float64 | int](t testing.TB, got, want T) {
	t.Helper()

	em := intErrorMessage
	if reflect.ValueOf(got).Kind() == reflect.Float64 {
		em = floatErrorMessage
	}

	if got != want {
		t.Errorf(em)
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()

	if got != want {
		t.Errorf(floatErrorMessage, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := Perimeter(rect)
	want := 40.0

	assertNum(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("should pass area tests", func(t *testing.T) {
		tests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range tests {
			checkArea(t, tt.shape, tt.want)
		}
	})
}
