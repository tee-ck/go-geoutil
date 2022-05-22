// Package geoutil provides utility functions for working with geographic
package geoutil

import (
	"fmt"
)

// Point represents a point in 2D space.
type Point struct {
	// Lat latitude in degrees.
	Lat float64 `json:"lat" bson:"lat" yaml:"lat" toml:"lat" xml:"lat,attr"`

	// Lng longitude in degrees.
	Lng float64 `json:"lng" bson:"lng" yaml:"lng" toml:"lng" xml:"lng,attr"`
}

// DistanceTo returns the distance between two points.
func (p *Point) DistanceTo(dist *Point) Distance {
	return GetDistance(p, dist)
}

// BoundaryOf returns the boundary of a point. The boundary is a square.
func (p *Point) BoundaryOf(distance Distance) (rect *Boundary) {
	return GetBoundary(p, distance)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%.9f, %.9f)", p.Lat, p.Lng)
}

// GeoHashEncode encode into a geohash.
func (p *Point) GeoHashEncode(precision int) []byte {
	return GeoHashEncode(p, precision)
}

// GeoHashDecode decode from a geohash.
func (p *Point) GeoHashDecode(geohash []byte) {
	*p = *GeoHashDecode(geohash)
}

// NewPoint returns a new Point with the given latitude and longitude.
func NewPoint(lat float64, lng float64) *Point {
	return &Point{Lat: lat, Lng: lng}
}

// Boundary represents a rectangular area in 2D space.
type Boundary struct {
	Min Point `json:"min" bson:"min" yaml:"min" toml:"min" xml:"min"`
	Max Point `json:"max" bson:"max" yaml:"max" toml:"max" xml:"max"`
}

// GetDistance returns the distance between two points.
func (r *Boundary) String() string {
	return fmt.Sprintf("(%s, %s)", r.Min.String(), r.Max.String())
}

// Distance represents a distance in meters.
type Distance float64

const (
	Meter        Distance = 1
	Inch         Distance = 0.0254
	Feets        Distance = 0.3048
	Yard         Distance = 0.9144
	Kilometer    Distance = 1000
	Mile         Distance = 1609.344
	NauticalMile Distance = 1852
)

// Meters returns the distance in meters.
func (d Distance) Meters() float64 {
	return float64(d)
}

// Inches returns the distance in inches.
func (d Distance) Inches() float64 {
	return float64(d) / float64(Inch)
}

// Yards returns the distance in yards.
func (d Distance) Yards() float64 {
	return float64(d) / float64(Yard)
}

// Kilometers returns the distance in kilometers.
func (d Distance) Kilometers() float64 {
	return float64(d) / float64(Kilometer)
}

// Miles returns the distance in miles.
func (d Distance) Miles() float64 {
	return float64(d) / float64(Mile)
}

// Feets returns the distance in feet.
func (d Distance) Feets() float64 {
	return float64(d) / float64(Feets)
}

// NauticalMiles returns the distance in nautical miles.
func (d Distance) NauticalMiles() float64 {
	return float64(d) / float64(NauticalMile)
}

// String returns the distance in meters as a string.
func (d Distance) String() string {
	return fmt.Sprintf("%.2f meters", d.Meters())
}
