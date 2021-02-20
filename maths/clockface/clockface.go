package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

//A Point represents a two dimensional Cartesion coordinate
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const minuteHandLength = 80
const hourHandLength = 50
const clockCenterX = 150
const clockCenterY = 150

//SecondHand is the unit vector of the second hand of an analogue clock at time 't'
//represented as a Point.
func SecondHand(t time.Time) Point {
	return MakeHand(SecondHandPoint(t), secondHandLength)
}

//MinuteHand is the unit vector of the minute hand of an analogue clock at time 't'
//represented as a Point.
func MinuteHand(t time.Time) Point {
	return MakeHand(MinuteHandPoint(t), minuteHandLength)
}

//HourHand is the unit vector of the hour hand of an analogue clock at time 't'
//represented as a Point.
func HourHand(t time.Time) Point {
	return MakeHand(HourHandPoint(t), hourHandLength)
}

//MakeHand is a helper method used to return the point of the hand
func MakeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}                //scale
	p = Point{p.X, -p.Y}                                 //flip
	return Point{p.X + clockCenterX, p.Y + clockCenterY} //translate
}

//SecondsInRadians return the second hand radians based on the in time
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

//MinutesInRadians return the minute hand radians based on the in time
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

//HoursInRadians return the hour hand radians based on the in time
func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

//SecondHandPoint return the point of the Second Hand based on the in time
func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	return angleToPoint(angle)
}

//MinuteHandPoint return the point of the Minute Hand based on the in time
func MinuteHandPoint(t time.Time) Point {
	angle := MinutesInRadians(t)
	return angleToPoint(angle)
}

//HourHandPoint return the point of the Hour Hand based on the in time
func HourHandPoint(t time.Time) Point {
	angle := HoursInRadians(t)
	return angleToPoint(angle)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

//SVGWriter writes an SVG representation of an analogue clock, showing the time t to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHandForSVG(w, t)
	MinuteHandForSVG(w, t)
	HourHandForSVG(w, t)
	io.WriteString(w, svgEnd)
}

//SecondHandForSVG writes an second hand SVG representation of an analogue clock, showing the time t to the writer w
func SecondHandForSVG(w io.Writer, t time.Time) {
	p := SecondHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

//MinuteHandForSVG writes an minute hand SVG representation of an analogue clock, showing the time t to the writer w
func MinuteHandForSVG(w io.Writer, t time.Time) {
	p := MinuteHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

//HourHandForSVG writes an hour hand SVG representation of an analogue clock, showing the time t to the writer w
func HourHandForSVG(w io.Writer, t time.Time) {
	p := HourHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
