package clockface

import (
	"math"
	"time"
)

// Requirements:
// Hour hand -> 50 unit long
// Minute hand -> 80 unit long
// Second hand -> 90 unit long
const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

// A Point represents a two-dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// Translate seconds into radians to get the angle (alpha) of the point.
func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

// Get the second hand point based on angle of the radians.
func secondHandPoint(t time.Time) Point {
	return angleHandPoint(secondsInRadians(t))
}

// SecondHand is the unit vector of the second hand of an analogue clock at time
// `t` represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	p = Point{
		p.X * secondHandLength,
		p.Y * secondHandLength,
	} // scale the point to the length of hand based on Requirements.
	p = Point{p.X, -p.Y} // flip the point
	p = Point{
		p.X + clockCentreX,
		p.Y + clockCentreY,
	} // translate to the scale of 1 unit in the circle

	return p
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	return angleHandPoint(minutesInRadians(t))
}

func angleHandPoint(a float64) Point {
	x := math.Sin(a)
	y := math.Cos(a)

	return Point{x, y}
}

func MinuteHand(t time.Time) Point {
	p := minuteHandPoint(t)

	p = Point{
		p.X * minuteHandLength,
		p.Y * minuteHandLength,
	} // scale the point to the length of hand based on Requirements.
	p = Point{p.X, -p.Y} // flip the point
	p = Point{
		p.X + clockCentreX,
		p.Y + clockCentreY,
	} // translate to the scale of 1 unit in the circle

	return p
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

func hourHandPoint(t time.Time) Point {
	return angleHandPoint(hoursInRadians(t))
}

func HourHand(t time.Time) Point {
	p := hourHandPoint(t)

	p = Point{
		p.X * hourHandLength,
		p.Y * hourHandLength,
	} // scale the point to the length of hand based on Requirements.
	p = Point{p.X, -p.Y} // flip the point
	p = Point{
		p.X + clockCentreX,
		p.Y + clockCentreY,
	} // translate to the scale of 1 unit in the circle

	return p
}
