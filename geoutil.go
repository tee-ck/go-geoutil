// Package geoutil provides utility functions for working with geographic
package geoutil

import (
	"fmt"
)

// Point represents a point in 2D space.
type Point struct {
	// Lat latitude in degrees.
	Lat float64 `json:"lat" bson:"lat" yaml:"lat" toml:"lat" xml:"lat"`

	// Lng longitude in degrees.
	Lng float64 `json:"lng" bson:"lng" yaml:"lng" toml:"lng" xml:"lng"`
}

// DistanceTo returns the distance between two points.
func (p *Point) DistanceTo(dist *Point) Distance {
	return GetDistance(p, dist)
}

// BoundaryOf returns the boundary of a point. The boundary is a square.
func (p *Point) BoundaryOf(meters float64) (rect *Boundary) {
	return GetBoundary(p, meters)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%.9f, %.9f)", p.Lat, p.Lng)
}

// GeoHashEncode encode into a geohash.
func (p *Point) GeoHashEncode(precision ...int) []byte {
	return GeoHashEncode(p, precision...)
}

// GeoHashDecode decode from a geohash.
func (p *Point) GeoHashDecode(geohash []byte) {
	*p = *GeoHashDecode(geohash)
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
	Meter        = 1
	Feet         = 0.3048 * Meter
	Kilometer    = 1000 * Meter
	Mile         = 1609.344 * Meter
	NauticalMile = 1852 * Meter
)

// Meters returns the distance in meters.
func (d Distance) Meters() float64 {
	return float64(d)
}

// Kilometers returns the distance in kilometers.
func (d Distance) Kilometers() float64 {
	return float64(d) / Kilometer
}

// Miles returns the distance in miles.
func (d Distance) Miles() float64 {
	return float64(d) / Mile
}

// Feet returns the distance in feet.
func (d Distance) Feet() float64 {
	return float64(d) / Feet
}

// NauticalMiles returns the distance in nautical miles.
func (d Distance) NauticalMiles() float64 {
	return float64(d) / NauticalMile
}
