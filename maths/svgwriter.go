package clockface

import (
	"fmt"
	"io"
	"log"
	"time"
)

const (
	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
	<svg xmlns="http://www.w3.org/2000/svg"
	width="100%"
	height="100%"
	viewBox="0 0 300 300"
	version="2.0">`

	bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

	svgEnd = `</svg>`
)

func makeHandTag(p Point, l float64) Point  {
	p = Point{p.X * l, p.Y * l} // scale the point to the length of hand based on Requirements.
	p = Point{p.X, -p.Y}  // flip the point

	return Point{
		p.X + clockCentreX,
		p.Y + clockCentreY,
	} // translate to the scale of 1 unit in the circle
}

func secondHandTag(w io.Writer, t time.Time) {
	p := makeHandTag(secondHandPoint(t), secondHandLength)

	_, err := fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
	if err != nil {
		log.Fatalf("error in fprintf %v", err)
	}
}

func minuteHandTag(w io.Writer, t time.Time) {
	p := makeHandTag(minuteHandPoint(t), minuteHandLength)

	_, err := fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
	if err != nil {
		log.Fatalf("error in fprintf %v", err)
	}
}

func hourHandTag(w io.Writer, t time.Time) {
	p := makeHandTag(hourHandPoint(t), hourHandLength)

	_, err := fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
	if err != nil {
		log.Fatalf("error in fprintf %v", err)
	}
}

func SVGWriter(w io.Writer, t time.Time) {
	_, err := io.WriteString(w, svgStart)
	if err != nil {
		log.Fatalf("err 1 %v", err)
	}

	_, err = io.WriteString(w, bezel)
	if err != nil {
		log.Fatalf("err 2 %v", err)
	}

	secondHandTag(w, t)
	minuteHandTag(w, t)
	hourHandTag(w, t)

	_, err = io.WriteString(w, svgEnd)
	if err != nil {
		log.Fatalf("err 4 %v", err)
	}
}
