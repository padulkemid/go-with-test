package clockface_test

import (
	"bytes"
	"encoding/xml"
	clockface "hello/maths"
	"testing"
	"time"
)

// This is an SVG struct that display clockface fully, imported from
// https://xml-to-go.github.io/
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func assertClockface(t testing.TB, got, want clockface.Point) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertClockLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}

	return false
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(312, time.October, 28, h, m, s, 0, time.UTC)
}

func TestSecondHand(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
		want := clockface.Point{X: 150, Y: 150 - 90}
		got := clockface.SecondHand(tm)

		assertClockface(t, got, want)
	})

	t.Run("at 30 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
		want := clockface.Point{X: 150, Y: 150 + 90}
		got := clockface.SecondHand(tm)

		assertClockface(t, got, want)
	})
}

func TestSVGWriterAtMidnight(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			b := bytes.Buffer{}
			svg := SVG{}

			clockface.SVGWriter(&b, c.time)
			err := xml.Unmarshal(b.Bytes(), &svg)
			if err != nil {
				t.Error("should not return error!")
			}

			if !assertClockLine(c.line, svg.Line) {
				t.Errorf(
					"expect to find the secondhand line %+v, in the SVG output %v",
					c.line,
					svg.Line,
				)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(c.time.Format("15:05:04"), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !assertClockLine(c.line, svg.Line) {
				t.Errorf(
					"Expected to find the minute hand line %+v, in the SVG lines %+v",
					c.line,
					svg.Line,
				)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(c.time.Format("15:05:04"), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !assertClockLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}
